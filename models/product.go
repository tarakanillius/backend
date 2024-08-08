package models

// Nutriments struct to parse the nutriments data
type Nutriments struct {
	EnergyKcal           float64 `bson:"energy-kcal" json:"energy_kcal"`
	EnergyKcal100g       float64 `bson:"energy-kcal_100g" json:"energy_kcal_100g"`
	EnergyKcalValue      float64 `bson:"energy-kcal_value" json:"energy_kcal_value"`
	EnergyKcalUnit       string  `bson:"energy-kcal_unit" json:"energy_kcal_unit"`
	EnergyKj             float64 `bson:"energy-kj" json:"energy_kj"`
	EnergyKj100g         float64 `bson:"energy-kj_100g" json:"energy_kj_100g"`
	EnergyKjValue        float64 `bson:"energy-kj_value" json:"energy_kj_value"`
	EnergyKjUnit         string  `bson:"energy-kj_unit" json:"energy_kj_unit"`
	Fat                  float64 `bson:"fat" json:"fat"`
	Fat100g              float64 `bson:"fat_100g" json:"fat_100g"`
	FatUnit              string  `bson:"fat_unit" json:"fat_unit"`
	Carbohydrates        float64 `bson:"carbohydrates" json:"carbohydrates"`
	Carbohydrates100g    float64 `bson:"carbohydrates_100g" json:"carbohydrates_100g"`
	CarbohydratesUnit    string  `bson:"carbohydrates_unit" json:"carbohydrates_unit"`
	Fiber                float64 `bson:"fiber" json:"fiber"`
	Fiber100g            float64 `bson:"fiber_100g" json:"fiber_100g"`
	FiberUnit            string  `bson:"fiber_unit" json:"fiber_unit"`
	Sugars               float64 `bson:"sugars" json:"sugars"`
	Sugars100g           float64 `bson:"sugars_100g" json:"sugars_100g"`
	SugarsUnit           string  `bson:"sugars_unit" json:"sugars_unit"`
	SaturatedFat         float64 `bson:"saturated_fat" json:"saturated_fat"`
	SaturatedFat100g     float64 `bson:"saturated_fat_100g" json:"saturated_fat_100g"`
	SaturatedFatUnit     string  `bson:"saturated_fat_unit" json:"saturated_fat_unit"`
	Proteins             float64 `bson:"proteins" json:"proteins"`
	Proteins100g         float64 `bson:"proteins_100g" json:"proteins_100g"`
	ProteinsUnit         string  `bson:"proteins_unit" json:"proteins_unit"`
	Salt                 float64 `bson:"salt" json:"salt"`
	Salt100g             float64 `bson:"salt_100g" json:"salt_100g"`
	SaltUnit             string  `bson:"salt_unit" json:"salt_unit"`
	NovaGroup            int     `bson:"nova-group" json:"nova_group"`
	NutritionScoreFR     int     `bson:"nutrition-score-fr" json:"nutrition_score_fr"`
	NutritionScoreFR100g int     `bson:"nutrition-score-fr_100g" json:"nutrition_score_fr_100g"`
}

// Product struct to parse the product data
type Product struct {
	ID             string            `bson:"_id" json:"id"`
	ProductName    string            `bson:"product_name" json:"product_name"`
	Labels         string            `bson:"labels" json:"labels"`
	NutriscoreTags []string          `bson:"nutriscore_tags" json:"nutriscore_tags"`
	NutrientLevels map[string]string `bson:"nutrient_levels" json:"nutrient_levels"`
	AdditivesTags  []string          `bson:"additives_tags" json:"additives_tags"`
	Nutriments     Nutriments        `bson:"nutriments" json:"nutriments"`
	MaxImgID       string            `bson:"max_imgid,omitempty" json:"max_imgid,omitempty"`
	ImageURL       string            `bson:"imageUrl,omitempty" json:"imageUrl,omitempty"`
  Keywords       []string           `bson:"_keywords" json:"Keywords"`
}
