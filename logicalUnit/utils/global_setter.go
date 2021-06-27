package utils

import "../globals"

func SetNumberOfGoodIndividuals(numberOfGood int) {
	if numberOfGood > 0 {
		globals.NumberOfGoods = numberOfGood
	}
}

func SetNumberOfBadIndividuals(numberOfBad int) {
	if numberOfBad > 0 {
		globals.NumberOfBads = numberOfBad
	}
}

func SetNumberOfFoodSources(numberOfFoodSources int) {
	if numberOfFoodSources > 0 {
		globals.NumberOfFoodSources = numberOfFoodSources
	}
}

func SetFinalEpoch(finalEpoch int) {
	if finalEpoch > 0 {
		globals.FinalEpoch = finalEpoch
	}
}

func SetLvl(lvl string) {
	var foundIndex = Find(globals.Lvls, lvl)
	if foundIndex == -1 {
		return
	}
	globals.MutationToGoodBoost = (len(globals.Lvls) - foundIndex - 1) * 10
	globals.BoostToSurvive = 10 + (len(globals.Lvls)-foundIndex-1)*1
}

func Find(source []string, value string) int {
	for i := 0; i < 5; i++ {
		if value == source[i] {
			return i
		}
	}
	return -1
}
