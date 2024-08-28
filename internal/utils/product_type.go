package utils

import "fmt"

func DetermineProductType(nutriscore2021, nutriscore2023 map[string]interface{}) string {
    getInt32Value := func(data map[string]interface{}, key string) (int32, bool) {
        if value, ok := data[key]; ok {
            switch v := value.(type) {
            case int32:
                return v, true
            case float64:
                return int32(v), true
            default:
                fmt.Printf("Unexpected type for key %s: %T\n", key, v)
            }
        }
        return 0, false
    }

    if value, ok := getInt32Value(nutriscore2023, "is_beverage"); ok && value == 1 {
        return "beverage"
    }

    if value, ok := getInt32Value(nutriscore2021, "is_beverage"); ok && value == 1 {
        return "beverage"
    }

    if value, ok := getInt32Value(nutriscore2023, "is_cheese"); ok && value == 1 {
        return "cheese"
    }

    if value, ok := getInt32Value(nutriscore2023, "is_red_meat_product"); ok && value == 1 {
        return "red_meat_product"
    }

    if value, ok := getInt32Value(nutriscore2023, "is_fat_oil_nuts_seeds"); ok && value == 1 {
        return "fat_oil_nuts_seeds"
    }

    return "general_food"
}
