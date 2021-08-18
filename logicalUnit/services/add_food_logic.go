package services

import (
	"../globals"
	"../models"
	"sync"
)

//Logika za dodavanje hrane
func feedThePopulation(foodSources []*models.FoodSource) {
	for _, foodSource := range foodSources {
		giveFoodToPleb(foodSource)
	}
}

func feedThePopulationParallel(foodSources []*models.FoodSource, wg *sync.WaitGroup) {
	var timesLoop = int(len(foodSources) / globals.ParallelNightInterval)
	for i := 0; i < timesLoop; i++ {
		wg.Add(globals.ParallelNightInterval)
		for j := 0; j < globals.ParallelNightInterval; j++ {
			//fmt.Println(i*150+j)
			go giveFoodToPlebParallel(foodSources[i*globals.ParallelNightInterval+j], wg)
		}
		wg.Wait()
	}
	var restFood = len(foodSources) % globals.ParallelNightInterval
	wg.Add(restFood)
	for i := timesLoop * globals.ParallelNightInterval; i < len(foodSources); i++ {
		//fmt.Println(i)
		go giveFoodToPlebParallel(foodSources[i], wg)
	}
	wg.Wait()
	/*wg.Add(len(foodSources))
	for _, foodSource := range foodSources {
		go giveFoodToPlebParallel(foodSource, wg)
	}
	wg.Wait()*/
}

func giveFoodToPleb(foodSource *models.FoodSource) {
	if foodSource.Occupied == 0 {
		return
	}
	if foodSource.Occupied == 1 {
		foodSource.OccupiedBy[0].Health = len(foodSource.Resources) * 100
		return
	}
	if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food
		foodSource.OccupiedBy[1].Health = food
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.BAD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food + food/2
		foodSource.OccupiedBy[1].Health = food / 2
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.BAD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[1].Health = food + food/2
	} else {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[0].Health = food / 2
	}
	//clearSource(foodSource)
}

func giveFoodToPlebParallel(foodSource *models.FoodSource, wg *sync.WaitGroup) {
	if foodSource.Occupied == 0 {
		defer wg.Done()
		return
	}
	if foodSource.Occupied == 1 {
		foodSource.OccupiedBy[0].Health = len(foodSource.Resources) * 100
		defer wg.Done()
		return
	}
	if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food
		foodSource.OccupiedBy[1].Health = food
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.BAD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food + food/2
		foodSource.OccupiedBy[1].Health = food / 2
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.BAD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[1].Health = food + food/2
	} else {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[0].Health = food / 2
	}
	defer wg.Done()
	//clearSource(foodSource)
}

func clearSource(foodSource *models.FoodSource) {
	foodSource.OccupiedBy = nil
	foodSource.Occupied = 0
}

func clearSourceParallel(i int, foodSource *models.FoodSource, wg *sync.WaitGroup) {
	foodSource.OccupiedBy = nil
	foodSource.Occupied = 0
	//fmt.Println("INDEX: ", i)
	defer wg.Done()
}
