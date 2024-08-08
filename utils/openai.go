package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	// Load the environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// sendOpenAIRequest sends a request to the OpenAI API and returns the response body.
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

func GenerateReceipt(products []string) (string, error) {
	prompt := fmt.Sprintf("Provide a recipe using the following ingredients: %v", products)

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4-turbo",
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": prompt,
					},
				},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to generate receipt: %v", err)
	}

	// Extract recipe from response
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

// AnalyzeImage sends an image to OpenAI API and returns key terms.
func AnalyzeImage(imageData []byte) (string, error) {
	imageBase64 := base64.StdEncoding.EncodeToString(imageData)

	prompt := "Describe one product in the image and choose the most suitable keywords for this product, for further searching in the openfoodfacts database. Keywords should be in lower case and without spaces, e.g., migros, protein, drink, oh. Use the language that is written on the product. IN RESPONSE BODY SHOULD BE ONLY 4 KEYWORDS"

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4-turbo",
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": prompt,
					},
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url": "data:image/png;base64," + imageBase64,
						},
					},
				},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to analyze image: %v", err)
	}

	// Extract keywords from response
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

// GenerateRecommendations provides recommendations for additional ingredients.
func GenerateRecommendations(products []string) (string, error) {
	prompt := fmt.Sprintf("Given the following products: %v, suggest additional ingredients that could be added to make a meal.", products)

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-4-turbo",
		"messages": []map[string]interface{}{
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "text",
						"text": prompt,
					},
				},
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request body: %v", err)
	}

	result, err := sendOpenAIRequest(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to generate recommendations: %v", err)
	}

	// Extract recommendations from response
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
