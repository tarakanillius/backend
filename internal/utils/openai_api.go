package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io"
	"net/http"
	"os"
	"sync"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}
}

func sendOpenAIRequest(requestBody []byte) (map[string]interface{}, error) {
	openaiAPIKey := os.Getenv("OPENAI_API_KEY")
	if openaiAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is not set")
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return result, nil
}

func extractContent(result map[string]interface{}) (string, error) {
	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("unexpected response format")
	}

	firstChoice := choices[0].(map[string]interface{})
	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	return content, nil
}

func generateRequestBody(model, prompt string, content []map[string]interface{}) ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"model":    model,
		"messages": []map[string]interface{}{{"role": "user", "content": content}},
	})
}

func GenerateReceipt(products []string) (string, error) {
	prompt := fmt.Sprintf("Provide a recipe using the following ingredients: %v", products)

	requestBody, err := generateRequestBody("gpt-4o-mini", prompt, []map[string]interface{}{{"type": "text", "text": prompt}})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to generate receipt: %v", err)
	}

	return extractContent(result)
}

func AnalyzeImage(imageData []byte) (string, error) {
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)
	prompt := "Describe one product in the image and choose the most suitable keywords for this product, for further searching in the openfoodfacts database. Keywords should be in lower case and without spaces, e.g., migros, protein, drink, oh. Use the language that is written on the product. IN RESPONSE BODY SHOULD BE ONLY 4 KEYWORDS"

	requestBody, err := generateRequestBody("gpt-4o-mini", prompt, []map[string]interface{}{
		{"type": "text", "text": prompt},
		{"type": "image_url", "image_url": map[string]string{"url": "data:image/png;base64," + imageBase64}},
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to analyze image: %v", err)
	}

	return extractContent(result)
}

func GenerateRecommendations(products []string) (string, error) {
	prompt := fmt.Sprintf("Given the following products: %v, suggest additional ingredients that could be added to make a meal.", products)

	requestBody, err := generateRequestBody("gpt-4o-mini", prompt, []map[string]interface{}{{"type": "text", "text": prompt}})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to generate recommendations: %v", err)
	}

	return extractContent(result)
}

// Concurrency for generating multiple outputs
func GenerateAll(products []string, imageData []byte) (map[string]string, error) {
	var wg sync.WaitGroup
	results := make(map[string]string)
	errs := make(chan error, 3)
	mu := &sync.Mutex{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		receipt, err := GenerateReceipt(products)
		if err != nil {
			errs <- err
			return
		}
		mu.Lock()
		results["receipt"] = receipt
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		analysis, err := AnalyzeImage(imageData)
		if err != nil {
			errs <- err
			return
		}
		mu.Lock()
		results["analysis"] = analysis
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		recommendations, err := GenerateRecommendations(products)
		if err != nil {
			errs <- err
			return
		}
		mu.Lock()
		results["recommendations"] = recommendations
		mu.Unlock()
	}()

	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
