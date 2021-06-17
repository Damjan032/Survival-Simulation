package main

import (
	"./models"
	"./services"
	"fmt"
	"math/rand"
	"time"
)

func makeArray(len int) []int {
	newArray := make([]int, len)
	for i := 0; i < len; i++ {
		newArray[i] = i
	}
	return newArray
}

func isFreeFoodSource(source *models.FoodSource) bool {
	return source.Occupied < 2
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func simulation() {
	fmt.Println("SIMULATION")
	rand.Seed(time.Now().UnixNano())

	var population = models.NewPopulation(1, 5)

	/*for _, animal := range population {
		fmt.Println("My animal is:", animal)
	}
	fmt.Println(population[0].resources[0].quantity)
	fmt.Println("des")
	fmt.Println(rand.Intn(100))*/

	var foodSources = services.InitFoodSources(1000)
	for i := 0; i < 1000; i++ {
		fmt.Println("I: ", i)
		services.Day(population.Members, foodSources)
		//printPlebs(population.members)
		services.Night(population, 10, 10)
		fmt.Println("BAD", population.NumberOfBad, " GOOD", population.NumberOfGood)
		fmt.Println("SUM ", population.NumberOfBad+population.NumberOfGood)
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

}

func main() {
	//rand.Seed(time.Now().UnixNano())
	//
	//var population = NewPopulation(1, 5)
	//
	///*for _, animal := range population {
	//	fmt.Println("My animal is:", animal)
	//}
	//fmt.Println(population[0].resources[0].quantity)
	//fmt.Println("des")
	//fmt.Println(rand.Intn(100))*/
	//
	//var foodSources = initFoodSources(1000)
	//for i := 0; i < 1000; i++ {
	//	fmt.Println("I: ", i)
	//	day(population.members, foodSources)
	//	//printPlebs(population.members)
	//	night(population, 10, 10)
	//	fmt.Println("BAD", population.numberOfBad, " GOOD", population.numberOfGood)
	//	fmt.Println("SUM ", population.numberOfBad+population.numberOfGood)
	//}

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
	simulation()
	//printPlebs(population.members)
	initRouter()
	/*var population = NewPopulation(1, 5)
	fmt.Println(json.Marshal(population))
	btResult, _ := json.Marshal(population)
	fmt.Println(btResult)
	fmt.Println("\nUsing Marshal:\n" + string(btResult))
	fmt.Println("Pleb food:", population)*/
}

func printPlebs(plebs []*models.Individual) {
	for _, animal := range plebs {
		fmt.Println("Pleb food:", animal.Health)
	}
}
