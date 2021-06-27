package simulator

import (
	"../globals"
	"../models"
	"../services"
	"../utils"
	"fmt"
	"math/rand"
	"time"
)

func Simulate() {
	fmt.Println("SIMULATION")
	rand.Seed(time.Now().UnixNano())
	globals.CurrentEpoch = 0
	utils.SetNumberOfFoodSources(1000)
	utils.SetFinalEpoch(1000)
	utils.SetNumberOfBadIndividuals(1)
	utils.SetNumberOfGoodIndividuals(10)
	globals.Population = models.NewPopulation(globals.NumberOfBads, globals.NumberOfGoods)
	globals.FoodSources = services.InitFoodSources(globals.NumberOfFoodSources)
	utils.SetLvl("easy")

	for i := globals.CurrentEpoch; i <= globals.FinalEpoch; i++ {
		fmt.Println("I: ", i)
		services.ClearFoodSources(globals.FoodSources)
		services.Night(globals.Population)
		services.Day(globals.Population.Members, globals.FoodSources)
		fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	}

	//printPlebs(population.members)
	//day(population.members, foodSources)
	//night(population, 10,10)
	//fmt.Println(population.numberOfBad, " ", population.numberOfGood)
	//for _, foodSource := range foodSources {
	//	fmt.Println("Size:", len(foodSource.resources))
	//	fmt.Println("Occupied", foodSource.occupied)
	//	/*if len(foodSource.occupiedBy) != 0 {
	//		//foodSource.occupiedBy[0].health = 1000
	//		fmt.Println("TO BREEEEEEEEEEEE")
	//	}
	//	if foodSource.occupied != 0 {
	//		//foodSource.occupiedBy[0].health=1000
	//		fmt.Println("TO BREEEEEEEEEEEE")
	//	}*/
	//	//for _, foodUnit := range foodSource.resources {
	//	//	fmt.Println("Quantity", foodUnit.quantity, ", type", foodUnit.foodType)
	//	//}
	//}

	//printPlebs(population.members)
	fmt.Println(models.UNHEALTHY)

}

func printPlebs(plebs []*models.Individual) {
	for _, animal := range plebs {
		fmt.Println("Pleb food:", animal.Health)
	}
}
