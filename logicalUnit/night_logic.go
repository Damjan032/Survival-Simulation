package main

import (
	"math/rand"
	"sort"
)

const BoostToSurvive = 10
const MutationPercent = 10
const MaxHP = 100

func night(population *Population, percentToSurvive int, percentToReproduction int) {
	nightEating(population.members, percentToSurvive, percentToReproduction)
	doOrDie(population)

}

func nightEating(population []*Individual, percentToSurvive int, percentToReproduction int) {
	for _, pleb := range population {
		pleb.nightRes = nightTimeLogic(pleb)
	}
}

func doOrDie(population *Population) {
	var blackList []int //to kill
	var whiteList []int //to make a child
	for i := 0; i < len(population.members); i++ {
		if population.members[i].nightRes == DIED {
			blackList = append(blackList, i)
		} else if population.members[i].nightRes == REPRODUCED {
			whiteList = append(whiteList, i)
		}
	}
	population.members = append(population.members, reproduceMember(population, whiteList)...)
	killMembers(population, blackList)

}

func killMembers(population *Population, blackLis []int) {
	sort.Ints(blackLis)
	for i := 0; i < len(blackLis); i++ {
		// -i jer je vec obrisalo i onda neki 5. element kad je obrisan neko pre
		if population.members[blackLis[i]-i].typeOfIndividual == GOOD {
			population.numberOfGood--
		} else {
			population.numberOfBad--
		}

		population.members = append(population.members[:blackLis[i]-i], population.members[blackLis[i]+1-i:]...)
	}
}

func reproduceMember(population *Population, whiteList []int) []*Individual {
	var newMembers []*Individual
	for i := 0; i < len(whiteList); i++ {
		var newMember = NewIndividual(population.members[whiteList[i]].typeOfIndividual)
		if calculateIsHappened(MutationPercent) {
			convertType(newMember)
		}
		if newMember.typeOfIndividual == GOOD {
			population.numberOfGood++
		} else {
			population.numberOfBad++
		}
		newMembers = append(newMembers, newMember)

	}
	return newMembers
}

func nightTimeLogic(individual *Individual) ProductOfTheNight {
	if individual.health <= 0 {
		return DIED
	}
	if individual.health < MaxHP {
		if calculateIsHappened(MaxHP - individual.health + BoostToSurvive) {
			return DIED
		}
	} else if individual.health >= MaxHP {
		if calculateIsHappened(individual.health - MaxHP + BoostToSurvive) {
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
