// product.go
package models

type Nutriments struct {
	EnergyKcal        float64 `bson:"energy-kcal" json:"energy_kcal"`
	EnergyKcal100g    float64 `bson:"energy-kcal_100g" json:"energy_kcal_100g"`
	EnergyKcalValue   float64 `bson:"energy-kcal_value" json:"energy_kcal_value"`
	EnergyKcalUnit    string  `bson:"energy-kcal_unit" json:"energy_kcal_unit"`
	EnergyKj          float64 `bson:"energy-kj" json:"energy_kj"`
	EnergyKj100g      float64 `bson:"energy-kj_100g" json:"energy_kj_100g"`
	EnergyKjValue     float64 `bson:"energy-kj_value" json:"energy_kj_value"`
	EnergyKjUnit      string  `bson:"energy-kj_unit" json:"energy_kj_unit"`
	Fat               float64 `bson:"fat" json:"fat"`
	Fat100g           float64 `bson:"fat_100g" json:"fat_100g"`
	FatUnit           string  `bson:"fat_unit" json:"fat_unit"`
	Carbohydrates     float64 `bson:"carbohydrates" json:"carbohydrates"`
	Carbohydrates100g float64 `bson:"carbohydrates_100g" json:"carbohydrates_100g"`
	CarbohydratesUnit string  `bson:"carbohydrates_unit" json:"carbohydrates_unit"`
	Fiber             float64 `bson:"fiber" json:"fiber"`
	Fiber100g         float64 `bson:"fiber_100g" json:"fiber_100g"`
	FiberUnit         string  `bson:"fiber_unit" json:"fiber_unit"`
	Sugars            float64 `bson:"sugars" json:"sugars"`
	Sugars100g        float64 `bson:"sugars_100g" json:"sugars_100g"`
	SugarsUnit        string  `bson:"sugars_unit" json:"sugars_unit"`
	SugarsValue       float64 `bson:"sugars_value" json:"sugars_value"`
	SaturatedFat      float64 `bson:"saturated_fat" json:"saturated_fat"`
	SaturatedFat100g  float64 `bson:"saturated_fat_100g" json:"saturated_fat_100g"`
	SaturatedFatUnit  string  `bson:"saturated_fat_unit" json:"saturated_fat_unit"`
	SaturatedFatValue float64 `bson:"saturated_fat_value" json:"saturated_fat_value"`
	Sodium            float64 `bson:"sodium" json:"sodium"`
	Sodium100g        float64 `bson:"sodium_100g" json:"sodium_100g"`
	SodiumUnit        string  `bson:"sodium_unit" json:"sodium_unit"`
	SodiumValue       float64 `bson:"sodium_value" json:"sodium_value"`
	TransFat          float64 `bson:"trans_fat" json:"trans_fat"`
	TransFat100g      float64 `bson:"trans_fat_100g" json:"trans_fat_100g"`
	TransFatServing   float64 `bson:"trans_fat_serving" json:"trans_fat_serving"`
	TransFatUnit      string  `bson:"trans_fat_unit" json:"trans_fat_unit"`
	TransFatValue     float64 `bson:"trans_fat_value" json:"trans_fat_value"`
	VitaminA          float64 `bson:"vitamin_a" json:"vitamin_a"`
	VitaminA100g      float64 `bson:"vitamin_a_100g" json:"vitamin_a_100g"`
	VitaminAServing   float64 `bson:"vitamin_a_serving" json:"vitamin_a_serving"`
	VitaminAUnit      string  `bson:"vitamin_a_unit" json:"vitamin_a_unit"`
	VitaminAValue     float64 `bson:"vitamin_a_value" json:"vitamin_a_value"`
	VitaminC          float64 `bson:"vitamin_c" json:"vitamin_c"`
	VitaminC100g      float64 `bson:"vitamin_c_100g" json:"vitamin_c_100g"`
	VitaminCServing   float64 `bson:"vitamin_c_serving" json:"vitamin_c_serving"`
	VitaminCUnit      string  `bson:"vitamin_c_unit" json:"vitamin_c_unit"`
	VitaminCValue     float64 `bson:"vitamin_c_value" json:"vitamin_c_value"`
	Proteins          float64 `bson:"proteins" json:"proteins"`
	Proteins100g      float64 `bson:"proteins_100g" json:"proteins_100g"`
	ProteinsUnit      string  `bson:"proteins_unit" json:"proteins_unit"`
	ProteinsValue     float64 `bson:"proteins_value" json:"proteins_value"`
	Salt              float64 `bson:"salt" json:"salt"`
	Salt100g          float64 `bson:"salt_100g" json:"salt_100g"`
	SaltUnit          string  `bson:"salt_unit" json:"salt_unit"`
	SaltValue         float64 `bson:"salt_value" json:"salt_value"`
	NovaGroup         int     `bson:"nova-group" json:"nova_group"`
}

type Product struct {
	ID                                string                    `bson:"_id" json:"id"`
	ProductName                       string                    `bson:"product_name" json:"product_name"`
	Labels                            string                    `bson:"labels" json:"labels"`
	NutritionScore                    int                       `bson:"nutriscore_score" json:"nutriscore_score"`
	NutritionGrade                    string                    `bson:"nutriscore_grade" json:"nutriscore_grade"`
	Nutriscore                        map[string]NutriscoreData `bson:"nutriscore" json:"nutriscore"`
	NutrientLevels                    map[string]string         `bson:"nutrient_levels" json:"nutrient_levels"`
	AdditivesTags                     []string                  `bson:"additives_tags" json:"additives_tags"`
	Nutriments                        Nutriments                `bson:"nutriments" json:"nutriments"`
	MaxImgID                          string                    `bson:"max_imgid,omitempty" json:"max_imgid,omitempty"`
	ImageURL                          string                    `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
	Keywords                          []string                  `bson:"_keywords" json:"Keywords"`
	IngredientsNonNutritiveSweeteners int                       `bson:"ingredients_non_nutritive_sweeteners_n" json:"ingredients_non_nutritive_sweeteners_n"`
}

type NutriscoreData struct {
	Data map[string]interface{} `json:"data" bson:"data"`
}
