package utils

func CalculateNutritionScore(energy, sugar, saturates, sodium, fvl, fibre, protein float64) (int, string) {
	// Energy Points Calculation
	var energyPoints int
	switch {
	case energy <= 335:
		energyPoints = 0
	case energy <= 670:
		energyPoints = 1
	case energy <= 1005:
		energyPoints = 2
	case energy <= 1340:
		energyPoints = 3
	case energy <= 1675:
		energyPoints = 4
	case energy <= 2010:
		energyPoints = 5
	case energy <= 2345:
		energyPoints = 6
	case energy <= 2680:
		energyPoints = 7
	case energy <= 3015:
		energyPoints = 8
	case energy <= 3350:
		energyPoints = 9
	default:
		energyPoints = 10
	}

	// Sugar Points Calculation
	var sugarPoints int
	switch {
	case sugar <= 4.5:
		sugarPoints = 0
	case sugar <= 9:
		sugarPoints = 1
	case sugar <= 13.5:
		sugarPoints = 2
	case sugar <= 18:
		sugarPoints = 3
	case sugar <= 22.5:
		sugarPoints = 4
	case sugar <= 27:
		sugarPoints = 5
	case sugar <= 31:
		sugarPoints = 6
	case sugar <= 36:
		sugarPoints = 7
	case sugar <= 40:
		sugarPoints = 8
	case sugar <= 45:
		sugarPoints = 9
	default:
		sugarPoints = 10
	}

	// Saturated Fatty Acids (SFA) Points Calculation
	var sfaPoints int
	switch {
	case saturates <= 1:
		sfaPoints = 0
	case saturates <= 2:
		sfaPoints = 1
	case saturates <= 3:
		sfaPoints = 2
	case saturates <= 4:
		sfaPoints = 3
	case saturates <= 5:
		sfaPoints = 4
	case saturates <= 6:
		sfaPoints = 5
	case saturates <= 7:
		sfaPoints = 6
	case saturates <= 8:
		sfaPoints = 7
	case saturates <= 9:
		sfaPoints = 8
	case saturates <= 10:
		sfaPoints = 9
	default:
		sfaPoints = 10
	}

	// Sodium Points Calculation
	var sodiumPoints int
	switch {
	case sodium <= 90:
		sodiumPoints = 0
	case sodium <= 180:
		sodiumPoints = 1
	case sodium <= 270:
		sodiumPoints = 2
	case sodium <= 360:
		sodiumPoints = 3
	case sodium <= 450:
		sodiumPoints = 4
	case sodium <= 540:
		sodiumPoints = 5
	case sodium <= 630:
		sodiumPoints = 6
	case sodium <= 720:
		sodiumPoints = 7
	case sodium <= 810:
		sodiumPoints = 8
	case sodium <= 900:
		sodiumPoints = 9
	default:
		sodiumPoints = 10
	}

	// FVL Points Calculation
	var fvlPoints int
	switch {
	case fvl <= 40:
		fvlPoints = 0
	case fvl <= 60:
		fvlPoints = 1
	case fvl <= 80:
		fvlPoints = 2
	default:
		fvlPoints = 5
	}

	// Fibre Points Calculation
	var fibrePoints int
	switch {
	case fibre <= 0.9:
		fibrePoints = 0
	case fibre <= 1.9:
		fibrePoints = 1
	case fibre <= 2.8:
		fibrePoints = 2
	case fibre <= 3.7:
		fibrePoints = 3
	case fibre <= 4.7:
		fibrePoints = 4
	default:
		fibrePoints = 5
	}

	// Protein Points Calculation
	var proteinPoints int
	switch {
	case protein <= 1.6:
		proteinPoints = 0
	case protein <= 3.2:
		proteinPoints = 1
	case protein <= 4.8:
		proteinPoints = 2
	case protein <= 6.4:
		proteinPoints = 3
	case protein <= 8:
		proteinPoints = 4
	default:
		proteinPoints = 5
	}

	// Calculate Points A and Points C
	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints
	pointsC := fvlPoints + fibrePoints + proteinPoints

	// Calculate Nutrition Score (based on the provided formula)
	var nutritionScore int
	switch {
	case pointsA >= 0 && pointsA < 11:
		nutritionScore = pointsA - pointsC
	case pointsA >= 11 && fvlPoints == 5:
		nutritionScore = pointsA - pointsC
	default:
		nutritionScore = pointsA - fvlPoints - proteinPoints
	}

	// Determine Nutri-Score grade
	var nutriScoreGrade string
	switch {
	case nutritionScore < 0:
		nutriScoreGrade = "a"
	case nutritionScore < 3:
		nutriScoreGrade = "b"
	case nutritionScore < 11:
		nutriScoreGrade = "c"
	case nutritionScore < 19:
		nutriScoreGrade = "d"
	default:
		nutriScoreGrade = "e"
	}


	return nutritionScore, nutriScoreGrade
}
