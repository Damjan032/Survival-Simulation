package services

import (
	"../globals"
	"../models"
	"math/rand"
	"sort"
	"sync"
)

const MaxHP = 100

func Night(population *models.Population) {
	nightEating(population.Members)
	doOrDie(population)

}

func NightParallel(population *models.Population) {
	if len(globals.FoodSources) > 2500 {
		var wg sync.WaitGroup
		nightEatingParallel(population.Members, &wg)
		wg.Wait()
	} else {
		nightEating(population.Members)
	}

	doOrDie(population)

}

func nightEatingParallel(population []*models.Individual, wg *sync.WaitGroup) {
	//fmt.Println(len(foodSources))
	var timesLoop = int(len(population) / globals.ParallelNightInterval)
	//fmt.Println("FOOD SOURCES ", timesLoop)
	for i := 0; i < timesLoop; i++ {
		wg.Add(globals.ParallelNightInterval)
		for j := 0; j < globals.ParallelNightInterval; j++ {
			//fmt.Println(i*150+j)
			go nightTimeLogicParallel(population[i*globals.ParallelNightInterval+j], wg)
		}
		wg.Wait()
	}
	var restPopulation = len(population) % globals.ParallelNightInterval
	wg.Add(restPopulation)
	for i := timesLoop * globals.ParallelNightInterval; i < len(population); i++ {
		//fmt.Println(i)
		go nightTimeLogicParallel(population[i], wg)
	}
	wg.Wait()

	/*for _, pleb := range population {
		pleb.NightRes = nightTimeLogic(pleb)
	}
	defer wg.Done()*/
}

func nightEating(population []*models.Individual) {
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
		var newMemberMutation = globals.MutationPercent
		if newMember.TypeOfIndividual == models.BAD {
			newMemberMutation += globals.MutationToGoodBoost
		}
		if calculateIsHappened(newMemberMutation) {
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

func nightTimeLogicParallel(individual *models.Individual, wg *sync.WaitGroup) models.ProductOfTheNight {
	if individual.Health <= 0 {
		individual.NightRes = models.DIED
		defer wg.Done()
		return models.DIED
	}
	if individual.Health < MaxHP {
		if calculateIsHappened(MaxHP - individual.Health + globals.BoostToSurvive) {
			individual.NightRes = models.DIED
			defer wg.Done()
			return models.DIED
		}
	} else if individual.Health >= MaxHP {
		if calculateIsHappened(individual.Health - MaxHP + globals.BoostToSurvive) {
			//fmt.Println("UMNOZIO SE")

			individual.NightRes = models.REPRODUCED
			defer wg.Done()
			return models.REPRODUCED
		}

	}

	individual.NightRes = models.NEUTRAL
	defer wg.Done()
	return models.NEUTRAL

}

func nightTimeLogic(individual *models.Individual) models.ProductOfTheNight {
	if individual.Health <= 0 {
		return models.DIED
	}
	if individual.Health < MaxHP {
		if calculateIsHappened(MaxHP - individual.Health + globals.BoostToSurvive) {
			return models.DIED
		}
	} else if individual.Health >= MaxHP {
		if calculateIsHappened(individual.Health - MaxHP + globals.BoostToSurvive) {
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
