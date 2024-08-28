package utils

func CalculateGeneralFoodScore(
	energy int, sugar float64, saturates float64, sodium int,
	fvl float64, fiber float64, protein float64,
) int {

	energyPoints := 0
	if energy > 335 {
		energyPoints = (energy - 335) / 335
		if energyPoints > 10 {
			energyPoints = 10
		}
	}

	sugarPoints := 0
	if sugar > 4.5 {
		sugarPoints = int((sugar - 4.5) / 4.5)
		if sugarPoints > 10 {
			sugarPoints = 10
		}
	}

	sfaPoints := 0
	if saturates > 1 {
		sfaPoints = int((saturates - 1) / 1)
		if sfaPoints > 10 {
			sfaPoints = 10
		}
	}

	sodiumPoints := 0
	if sodium > 90 {
		sodiumPoints = (sodium - 90) / 90
		if sodiumPoints > 10 {
			sodiumPoints = 10
		}
	}

	fvlPoints := 0
	if fvl > 40 {
		fvlPoints = int((fvl - 40) / 20)
		if fvlPoints > 5 {
			fvlPoints = 5
		}
	}

	fiberPoints := 0
	if fiber > 0.9 {
		fiberPoints = int((fiber - 0.9) / 1)
		if fiberPoints > 5 {
			fiberPoints = 5
		}
	}

	proteinPoints := 0
	if protein > 1.6 {
		proteinPoints = int((protein - 1.6) / 1.6)
		if proteinPoints > 5 {
			proteinPoints = 5
		}
	}

	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints
	pointsC := fvlPoints + fiberPoints + proteinPoints

	score := pointsA - pointsC
	if score < 0 {
		return 0
	}
	if score > 40 {
		return 40
	}
	return score
}

func CalculateBeverageScore(
	energy int, sugar float64, saturates float64, sodium int,
	fvl float64, fiber float64, protein float64,
	nonNutritiveSweetener bool,
) int {

	energyPoints := 0
	if energy > 30 {
		energyPoints = (energy - 30) / 30
		if energyPoints > 10 {
			energyPoints = 10
		}
	}

	sugarPoints := 0
	if sugar > 0 {
		sugarPoints = int((sugar - 0) / 1.5)
		if sugarPoints > 10 {
			sugarPoints = 10
		}
	}

	sfaPoints := 0
	if saturates > 1 {
		sfaPoints = int((saturates - 1) / 1)
		if sfaPoints > 10 {
			sfaPoints = 10
		}
	}

	sodiumPoints := 0
	if sodium > 90 {
		sodiumPoints = (sodium - 90) / 90
		if sodiumPoints > 10 {
			sodiumPoints = 10
		}
	}

	fvlPoints := 0
	if fvl > 40 {
		fvlPoints = int((fvl - 40) / 20)
		if fvlPoints > 10 {
			fvlPoints = 10
		}
	}

	fiberPoints := 0
	if fiber > 0.9 {
		fiberPoints = int((fiber - 0.9) / 1)
		if fiberPoints > 5 {
			fiberPoints = 5
		}
	}

	proteinPoints := 0
	if protein > 1.6 {
		proteinPoints = int((protein - 1.6) / 1.6)
		if proteinPoints > 5 {
			proteinPoints = 5
		}
	}

	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints
	pointsC := fvlPoints + fiberPoints + proteinPoints

	score := pointsA - pointsC
	if score < 0 {
		return 0
	}
	if score > 40 {
		return 40
	}
	if nonNutritiveSweetener {
		score = score + 1
	}
	if score < 0 {
		return 0
	}
	return score
}

func CalculateFatsOilsNutsSeedsScore(
	energy int, totalFat float64, saturates float64, sodium int,
	sugar float64, ratioSFA float64, energyFromSFA float64,
	fvl float64, fiber float64, protein float64,
) int {

	energyPoints := 0
	if energy > 335 {
		energyPoints = (energy - 335) / 335
		if energyPoints > 10 {
			energyPoints = 10
		}
	}

	sugarPoints := 0
	if sugar > 4.5 {
		sugarPoints = int((sugar - 4.5) / 4.5)
		if sugarPoints > 10 {
			sugarPoints = 10
		}
	}

	sfaPoints := 0
	if saturates > 10 {
		sfaPoints = int((saturates - 10) / 10)
		if sfaPoints > 10 {
			sfaPoints = 10
		}
	}

	sodiumPoints := 0
	if sodium > 90 {
		sodiumPoints = (sodium - 90) / 90
		if sodiumPoints > 10 {
			sodiumPoints = 10
		}
	}

	fvlPoints := 0
	if fvl > 40 {
		fvlPoints = int((fvl - 40) / 20)
		if fvlPoints > 5 {
			fvlPoints = 5
		}
	}

	fiberPoints := 0
	if fiber > 0.9 {
		fiberPoints = int((fiber - 0.9) / 1)
		if fiberPoints > 5 {
			fiberPoints = 5
		}
	}

	proteinPoints := 0
	if protein > 1.6 {
		proteinPoints = int((protein - 1.6) / 1.6)
		if proteinPoints > 5 {
			proteinPoints = 5
		}
	}

	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints
	pointsC := fvlPoints + fiberPoints + proteinPoints

	score := pointsA - pointsC
	if score < 0 {
		return 0
	}
	if score > 40 {
		return 40
	}
	return score
}

func CalculateRedMeatScore(
	energy int, sugar float64, saturates float64, sodium int,
	salt float64, protein float64,
) int {

	energyPoints := 0
	if energy > 335 {
		energyPoints = (energy - 335) / 335
		if energyPoints > 10 {
			energyPoints = 10
		}
	}

	sugarPoints := 0
	if sugar > 4.5 {
		sugarPoints = int((sugar - 4.5) / 4.5)
		if sugarPoints > 10 {
			sugarPoints = 10
		}
	}

	sfaPoints := 0
	if saturates > 1 {
		sfaPoints = int((saturates - 1) / 1)
		if sfaPoints > 10 {
			sfaPoints = 10
		}
	}

	sodiumPoints := 0
	if sodium > 90 {
		sodiumPoints = (sodium - 90) / 90
		if sodiumPoints > 10 {
			sodiumPoints = 10
		}
	}

	saltPoints := int(salt / 0.1)
	if saltPoints > 10 {
		saltPoints = 10
	}

	proteinPoints := 0
	if protein > 1.6 {
		proteinPoints = int((protein - 1.6) / 1.6)
		if proteinPoints > 5 {
			proteinPoints = 5
		}
	}

	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints + saltPoints
	pointsC := proteinPoints

	score := pointsA - pointsC
	if score < 0 {
		return 0
	}
	if score > 40 {
		return 40
	}
	return score
}

func CalculateCheeseScore(
	energy int, sugar float64, saturates float64, sodium int,
	protein float64,
) int {
	energyPoints := 0
	if energy > 335 {
		energyPoints = (energy - 335) / 335
		if energyPoints > 10 {
			energyPoints = 10
		}
	}

	sugarPoints := 0
	if sugar > 4.5 {
		sugarPoints = int((sugar - 4.5) / 4.5)
		if sugarPoints > 10 {
			sugarPoints = 10
		}
	}

	sfaPoints := 0
	if saturates > 1 {
		sfaPoints = int((saturates - 1) / 1)
		if sfaPoints > 10 {
			sfaPoints = 10
		}
	}

	sodiumPoints := 0
	if sodium > 90 {
		sodiumPoints = (sodium - 90) / 90
		if sodiumPoints > 10 {
			sodiumPoints = 10
		}
	}

	proteinPoints := 0
	if protein > 0.6 {
		proteinPoints = int((protein - 0.6) / 0.6)
		if proteinPoints > 10 {
			proteinPoints = 10
		}
	}

	pointsA := energyPoints + sugarPoints + sfaPoints + sodiumPoints
	pointsC := proteinPoints

	score := pointsA - pointsC
	if score < 0 {
		return 0
	}
	if score > 40 {
		return 40
	}
	return score
}
