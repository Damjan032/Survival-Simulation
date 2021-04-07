package main

import (
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

func isFreeFoodSource(source *FoodSource) bool {
	return source.occupied < 2
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//fmt.Println(randomArray(10))
	//niz := make([]int,5)
	//niz = makeArray(5)
	//niz = append(niz, 5)
	//fmt.Printf("\n Niz: %v \n", niz)
	//niz = append(niz[:0], niz[1:]...)
	//fmt.Printf("%v", niz)
	var population = NewPopulation(1, 5)

	/*for _, animal := range population {
		fmt.Println("My animal is:", animal)
	}
	fmt.Println(population[0].resources[0].quantity)
	fmt.Println("des")
	fmt.Println(rand.Intn(100))*/

	var foodSources = initFoodSources(10)
	for i := 0; i < 100; i++ {
		day(population.members, foodSources)
		//printPlebs(population.members)
		night(population, 10, 10)
		fmt.Println("BAD", population.numberOfBad, " GOOD", population.numberOfGood)
		fmt.Println("SUM ", population.numberOfBad+population.numberOfGood)
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

func printPlebs(plebs []*Individual) {
	for _, animal := range plebs {
		fmt.Println("Pleb food:", animal.health)
	}
}
