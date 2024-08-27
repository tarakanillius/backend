package utils

// Calculate the Nutri-Score for beverages
func CalculateBeverageScore(data map[string]float64) int {
	energy := data["energy"]
	sugars := data["sugars"]
	saturatedFat := data["saturated_fat"]
	sodium := data["sodium"]
	fruitsVegetablesNuts := data["fruits_vegetables_nuts_colza_walnut_olive_oils"]
	fiber := data["fiber"]

	energyPoints := calculateEnergyPoints(energy)
	sugarPoints := calculateSugarPoints(sugars)
	saturatedFatPoints := calculateSaturatedFatPoints(saturatedFat)
	sodiumPoints := calculateSodiumPoints(sodium)
	fvlPoints := calculateFruitsVegetablesLegumesPoints(fruitsVegetablesNuts)
	fiberPoints := calculateFiberPoints(fiber)
	proteinPoints := calculateProteinPoints(data["proteins"])

	totalPointsA := energyPoints + sugarPoints + saturatedFatPoints + sodiumPoints
	totalPointsC := fvlPoints + fiberPoints + proteinPoints

	return calculateScore(totalPointsA, totalPointsC)
}

// Calculate the Nutri-Score for general foods
func CalculateGeneralFoodScore(data map[string]float64) int {
	energy := data["energy"]
	sugars := data["sugars"]
	saturatedFat := data["saturated_fat"]
	sodium := data["sodium"]
	fvl := data["fruits_vegetables_nuts_colza_walnut_olive_oils"]
	fiber := data["fiber"]
	protein := data["proteins"]

	energyPoints := calculateEnergyPoints(energy)
	sugarPoints := calculateSugarPoints(sugars)
	saturatedFatPoints := calculateSaturatedFatPoints(saturatedFat)
	sodiumPoints := calculateSodiumPoints(sodium)
	fvlPoints := calculateFruitsVegetablesLegumesPoints(fvl)
	fiberPoints := calculateFiberPoints(fiber)
	proteinPoints := calculateProteinPoints(protein)

	totalPointsA := energyPoints + sugarPoints + saturatedFatPoints + sodiumPoints
	totalPointsC := fvlPoints + fiberPoints + proteinPoints

	return calculateScore(totalPointsA, totalPointsC)
}

func calculateEnergyPoints(energy float64) int {
	switch {
	case energy <= 335:
		return 0
	case energy <= 670:
		return 1
	case energy <= 1005:
		return 2
	case energy <= 1340:
		return 3
	case energy <= 1675:
		return 4
	case energy <= 2010:
		return 5
	case energy <= 2345:
		return 6
	case energy <= 2680:
		return 7
	case energy <= 3015:
		return 8
	case energy <= 3350:
		return 9
	default:
		return 10
	}
}

func calculateSugarPoints(sugar float64) int {
	switch {
	case sugar <= 4.5:
		return 0
	case sugar <= 9:
		return 1
	case sugar <= 13.5:
		return 2
	case sugar <= 18:
		return 3
	case sugar <= 22.5:
		return 4
	case sugar <= 27:
		return 5
	case sugar <= 31:
		return 6
	case sugar <= 36:
		return 7
	case sugar <= 40:
		return 8
	case sugar <= 45:
		return 9
	default:
		return 10
	}
}

func calculateSaturatedFatPoints(saturatedFat float64) int {
	switch {
	case saturatedFat <= 1:
		return 0
	case saturatedFat <= 2:
		return 1
	case saturatedFat <= 3:
		return 2
	case saturatedFat <= 4:
		return 3
	case saturatedFat <= 5:
		return 4
	case saturatedFat <= 6:
		return 5
	case saturatedFat <= 7:
		return 6
	case saturatedFat <= 8:
		return 7
	case saturatedFat <= 9:
		return 8
	case saturatedFat <= 10:
		return 9
	default:
		return 10
	}
}

func calculateSodiumPoints(sodium float64) int {
	sodiumMg := sodium * 1000 // Convert sodium to mg
	switch {
	case sodiumMg <= 90:
		return 0
	case sodiumMg <= 180:
		return 1
	case sodiumMg <= 270:
		return 2
	case sodiumMg <= 360:
		return 3
	case sodiumMg <= 450:
		return 4
	case sodiumMg <= 540:
		return 5
	case sodiumMg <= 630:
		return 6
	case sodiumMg <= 720:
		return 7
	case sodiumMg <= 810:
		return 8
	case sodiumMg <= 900:
		return 9
	default:
		return 10
	}
}

func calculateFruitsVegetablesLegumesPoints(fvl float64) int {
	switch {
	case fvl <= 40:
		return 0
	case fvl <= 60:
		return 2
	case fvl <= 80:
		return 4
	default:
		return 10
	}
}

func calculateFiberPoints(fiber float64) int {
	switch {
	case fiber <= 0.9:
		return 0
	case fiber <= 1.9:
		return 1
	case fiber <= 2.8:
		return 2
	case fiber <= 3.7:
		return 3
	case fiber <= 4.7:
		return 4
	default:
		return 5
	}
}

func calculateProteinPoints(protein float64) int {
	switch {
	case protein <= 1.6:
		return 0
	case protein <= 3.2:
		return 1
	case protein <= 4.8:
		return 2
	case protein <= 6.4:
		return 3
	case protein <= 8:
		return 4
	default:
		return 5
	}
}

func calculateScore(pointsA, pointsC int) int {
	if pointsA >= 0 && pointsA < 11 {
		return pointsA
	}
	if pointsA >= 11 && pointsC == 10 {
		return pointsA - pointsC
	}
	return pointsA - pointsC
}
