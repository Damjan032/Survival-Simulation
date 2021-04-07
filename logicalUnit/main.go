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
	var population []*Individual
	population = initPopulation(25, 5)

	/*for _, animal := range population {
		fmt.Println("My animal is:", animal)
	}
	fmt.Println(population[0].resources[0].quantity)
	fmt.Println("des")
	fmt.Println(rand.Intn(100))*/

	var foodSources = initFoodSources(16)

	merge(population, foodSources)
	for _, foodSource := range foodSources {
		fmt.Println("Size:", len(foodSource.resources))
		fmt.Println("Occupied", foodSource.occupied)
		if foodSource.occupied == 1 {
			foodSource.occupiedBy[0].health = 1000
			fmt.Println("TO BREEEEEEEEEEEE")
		}
		if foodSource.occupied == 0 {
			//foodSource.occupiedBy[0].health=1000
			fmt.Println("TO BREEEEEEEEEEEE")
		}
		for _, foodUnit := range foodSource.resources {
			fmt.Println("Quantity", foodUnit.quantity, ", type", foodUnit.foodType)
		}
	}

	for _, animal := range population {
		fmt.Println("Pleb food:", animal.health)
		fmt.Println("Pleb food:", len(animal.resources))
	}
}
