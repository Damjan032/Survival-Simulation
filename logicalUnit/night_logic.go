package main

import (
	"math/rand"
	"sort"
)

const BoostToSurvive = 10
const MutationPercent = 10
const MaxHP = 100

func night(population *Population, percentToSurvive int, percentToReproduction int) {
	nightEating(population.Members, percentToSurvive, percentToReproduction)
	doOrDie(population)

}

func nightEating(population []*Individual, percentToSurvive int, percentToReproduction int) {
	for _, pleb := range population {
		pleb.NightRes = nightTimeLogic(pleb)
	}
}

func doOrDie(population *Population) {
	var blackList []int //to kill
	var whiteList []int //to make a child
	for i := 0; i < len(population.Members); i++ {
		if population.Members[i].NightRes == DIED {
			blackList = append(blackList, i)
		} else if population.Members[i].NightRes == REPRODUCED {
			whiteList = append(whiteList, i)
		}
	}
	population.Members = append(population.Members, reproduceMember(population, whiteList)...)
	killMembers(population, blackList)

}

func killMembers(population *Population, blackLis []int) {
	sort.Ints(blackLis)
	for i := 0; i < len(blackLis); i++ {
		// -i jer je vec obrisalo i onda neki 5. element kad je obrisan neko pre
		if population.Members[blackLis[i]-i].TypeOfIndividual == GOOD {
			population.NumberOfGood--
		} else {
			population.NumberOfBad--
		}

		population.Members = append(population.Members[:blackLis[i]-i], population.Members[blackLis[i]+1-i:]...)
	}
}

func reproduceMember(population *Population, whiteList []int) []*Individual {
	var newMembers []*Individual
	for i := 0; i < len(whiteList); i++ {
		var newMember = NewIndividual(population.Members[whiteList[i]].TypeOfIndividual)
		if calculateIsHappened(MutationPercent) {
			convertType(newMember)
		}
		if newMember.TypeOfIndividual == GOOD {
			population.NumberOfGood++
		} else {
			population.NumberOfBad++
		}
		newMembers = append(newMembers, newMember)

	}
	return newMembers
}

func nightTimeLogic(individual *Individual) ProductOfTheNight {
	if individual.Health <= 0 {
		return DIED
	}
	if individual.Health < MaxHP {
		if calculateIsHappened(MaxHP - individual.Health + BoostToSurvive) {
			return DIED
		}
	} else if individual.Health >= MaxHP {
		if calculateIsHappened(individual.Health - MaxHP + BoostToSurvive) {
			//fmt.Println("UMNOZIO SE")
			return REPRODUCED
		}

	}
	return NEUTRAL

}

func calculateIsHappened(marginalValue int) bool {
	if rand.Intn(MaxHP) > marginalValue {
		return false
	}
	return true
}
