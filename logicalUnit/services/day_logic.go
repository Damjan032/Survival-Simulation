package services

import (
	"../globals"
	"../models"
	"sync"
)

func Day(population []*models.Individual, foodSources []*models.FoodSource) {
	Merge(population, foodSources)
	feedThePopulation(foodSources)
}

func DayParallel(population []*models.Individual, foodSources []*models.FoodSource) {
	Merge(population, foodSources)
	if len(globals.FoodSources) > 2500 {
		var wg sync.WaitGroup
		feedThePopulationParallel(foodSources, &wg)
		wg.Wait()
	} else {
		feedThePopulation(foodSources)
	}
}

func ClearFoodSourcesParallel(foodSources []*models.FoodSource, wg *sync.WaitGroup) {
	var intreval = 16
	//fmt.Println(len(foodSources))
	var timesLoop = int(len(foodSources) / intreval)
	//fmt.Println("FOOD SOURCES ", timesLoop)
	for i := 0; i < timesLoop; i++ {
		wg.Add(intreval)
		for j := 0; j < intreval; j++ {
			//fmt.Println(i*150+j)
			go clearSourceParallel(i*intreval+j, foodSources[i*intreval+j], wg)
		}
		wg.Wait()
	}
	var restFood = len(foodSources) % intreval
	wg.Add(restFood)
	for i := timesLoop * intreval; i < len(foodSources); i++ {
		//fmt.Println(i)
		go clearSourceParallel(i, foodSources[i], wg)
	}
	wg.Wait()
	/*for _, foodSource := range foodSources {
		clearSource(foodSource)
	}*/

}

func ClearFoodSources(foodSources []*models.FoodSource) {
	for _, foodSource := range foodSources {
		clearSource(foodSource)
	}
}
