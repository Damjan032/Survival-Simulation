package simulator

import (
	"../globals"
	"../models"
	"../services"
	"../utils"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func setData() {
	fmt.Println("SIMULATION")
	rand.Seed(time.Now().UnixNano())
	globals.CurrentEpoch = 0
	utils.SetNumberOfFoodSources(500)
	utils.SetFinalEpoch(100)
	utils.SetNumberOfBadIndividuals(1)
	utils.SetNumberOfGoodIndividuals(10)
	globals.Population = models.NewPopulation(globals.NumberOfBads, globals.NumberOfGoods)
	globals.FoodSources = services.InitFoodSources(globals.NumberOfFoodSources)
	utils.SetLvl("easy")
}

func SimulateParallel() {

	setData()

	globals.ParallelNightInterval = 16
	fmt.Println(len(globals.FoodSources))
	if len(globals.FoodSources) < 6000 {
		globals.ParallelNightInterval = int(len(globals.FoodSources) / 4)
	}
	if len(globals.FoodSources) < 2001 {
		globals.ParallelNightInterval = 500
	}
	fmt.Println(globals.ParallelNightInterval)
	for i := globals.CurrentEpoch; i <= globals.FinalEpoch; i++ {

		//services.ClearFoodSources(globals.FoodSources)

		var wg sync.WaitGroup
		//go services.ClearFoodSources(globals.FoodSources)
		services.ClearFoodSourcesParallel(globals.FoodSources, &wg)

		services.NightParallel(globals.Population)

		//services.Day(globals.Population.Members, globals.FoodSources)
		services.DayParallel(globals.Population.Members, globals.FoodSources)

		//printInfo(i)

		//fmt.Println("I: ", i)
		//fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		//fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)

		if globals.Population.NumberOfBad+globals.Population.NumberOfGood > globals.NumberOfFoodSources*5 {
			fmt.Println("ERROR PARALLEL")
			globals.ErrorParallel += 1
			break
		}
	}

	fmt.Println(models.UNHEALTHY)

}

func SimulateParallel2() {

	setData()

	globals.ParallelNightInterval = 16
	fmt.Println(len(globals.FoodSources))
	if len(globals.FoodSources) < 6000 {
		globals.ParallelNightInterval = int(len(globals.FoodSources) / 4)
	}
	if len(globals.FoodSources) < 2001 {
		globals.ParallelNightInterval = 500
	}
	fmt.Println(globals.ParallelNightInterval)
	for i := globals.CurrentEpoch; i <= globals.FinalEpoch; i++ {

		//services.ClearFoodSources(globals.FoodSources)

		//var wg sync.WaitGroup
		services.ClearFoodSources(globals.FoodSources)

		//services.ClearFoodSourcesParallel(globals.FoodSources, &wg)

		services.NightParallel(globals.Population)

		services.Day(globals.Population.Members, globals.FoodSources)
		//services.DayParallel(globals.Population.Members, globals.FoodSources)

		//printInfo(i)

		//fmt.Println("I: ", i)
		//fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		//fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)

		if globals.Population.NumberOfBad+globals.Population.NumberOfGood > globals.NumberOfFoodSources*5 {
			fmt.Println("ERROR PARALLEL")
			globals.ErrorNormal += 1
			break
		}
	}

	fmt.Println(models.UNHEALTHY)

}

func Simulate() {

	setData()

	for i := globals.CurrentEpoch; i <= globals.FinalEpoch; i++ {

		services.ClearFoodSources(globals.FoodSources)
		//var wg sync.WaitGroup
		//wg.Add(1)
		////go services.ClearFoodSources(globals.FoodSources)
		//go services.ClearFoodSourcesParallel(globals.FoodSources, &wg)
		//wg.Wait()
		//

		services.Night(globals.Population)
		services.Day(globals.Population.Members, globals.FoodSources)
		//printInfo(i)

		//fmt.Println("I: ", i)
		//fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		//fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
		if globals.Population.NumberOfBad+globals.Population.NumberOfGood > globals.NumberOfFoodSources*10 {
			fmt.Println("ERROR NORMAL")
			globals.ErrorNormal += 1
			break
		}
	}

	fmt.Println(models.UNHEALTHY)

}

func printPlebs(plebs []*models.Individual) {
	for _, animal := range plebs {
		fmt.Println("Pleb food:", animal.Health)
	}
}

func printInfo(i int) {
	fmt.Println("I: ", i)
	fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
	fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
}
