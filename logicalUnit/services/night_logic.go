package services

import (
	"../models"
	"math/rand"
	"sort"
)

const BoostToSurvive = 10
const MutationPercent = 10
const MaxHP = 100

func Night(population *models.Population, percentToSurvive int, percentToReproduction int) {
	nightEating(population.Members, percentToSurvive, percentToReproduction)
	doOrDie(population)

}

func nightEating(population []*models.Individual, percentToSurvive int, percentToReproduction int) {
	for _, pleb := range population {
		pleb.NightRes = nightTimeLogic(pleb)
	}
}

func doOrDie(population *models.Population) {
	var blackList []int //to kill
	var whiteList []int //to make a child
	for i := 0; i < len(population.Members); i++ {
		if population.Members[i].NightRes == models.DIED {
			blackList = append(blackList, i)
		} else if population.Members[i].NightRes == models.REPRODUCED {
			whiteList = append(whiteList, i)
		}
	}
	population.Members = append(population.Members, reproduceMember(population, whiteList)...)
	killMembers(population, blackList)

}

func killMembers(population *models.Population, blackLis []int) {
	sort.Ints(blackLis)
	for i := 0; i < len(blackLis); i++ {
		// -i jer je vec obrisalo i onda neki 5. element kad je obrisan neko pre
		if population.Members[blackLis[i]-i].TypeOfIndividual == models.GOOD {
			population.NumberOfGood--
		} else {
			population.NumberOfBad--
		}

		population.Members = append(population.Members[:blackLis[i]-i], population.Members[blackLis[i]+1-i:]...)
	}
}

func reproduceMember(population *models.Population, whiteList []int) []*models.Individual {
	var newMembers []*models.Individual
	for i := 0; i < len(whiteList); i++ {
		var newMember = models.NewIndividual(population.Members[whiteList[i]].TypeOfIndividual)
		if calculateIsHappened(MutationPercent) {
			convertIndividualType(newMember)
		}
		if newMember.TypeOfIndividual == models.GOOD {
			population.NumberOfGood++
		} else {
			population.NumberOfBad++
		}
		newMembers = append(newMembers, newMember)

	}
	return newMembers
}

func nightTimeLogic(individual *models.Individual) models.ProductOfTheNight {
	if individual.Health <= 0 {
		return models.DIED
	}
	if individual.Health < MaxHP {
		if calculateIsHappened(MaxHP - individual.Health + BoostToSurvive) {
			return models.DIED
		}
	} else if individual.Health >= MaxHP {
		if calculateIsHappened(individual.Health - MaxHP + BoostToSurvive) {
			//fmt.Println("UMNOZIO SE")
			return models.REPRODUCED
		}

	}
	return models.NEUTRAL

}

func calculateIsHappened(marginalValue int) bool {
	if rand.Intn(MaxHP) > marginalValue {
		return false
	}
	return true
}
