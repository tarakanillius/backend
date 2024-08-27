package utils

// DetermineProductType determines the type of product based on its Nutri-Score data.
func DetermineProductType(nutriscoreData map[string]interface{}) string {
	if isBeverage, ok := nutriscoreData["is_beverage"].(float64); ok && isBeverage == 1 {
		return "beverage"
	}
	if isCheese, ok := nutriscoreData["is_cheese"].(float64); ok && isCheese == 1 {
		return "cheese"
	}
	if isRedMeat, ok := nutriscoreData["is_red_meat_product"].(float64); ok && isRedMeat == 1 {
		return "red_meat_product"
	}
	if isFatOilNutsSeeds, ok := nutriscoreData["is_fat_oil_nuts_seeds"].(float64); ok && isFatOilNutsSeeds == 1 {
		return "fat_oil_nuts_seeds"
	}
	return "general_food"
}
