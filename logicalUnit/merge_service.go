package main

import (
	"fmt"
	"math/rand"
)

func merge(population []*Individual, foodSources []*FoodSource) {
	var populationFreeArray = makeArray(len(population))   //Lista id-ova
	var foodSourcesFreeArray = makeArray(len(foodSources)) //Lista id-ova
	//fmt.Printf("%v\n", populationFreeArray)
	//fmt.Printf("%v\n", foodSourcesFreeArray)
	// prolazim kroz listu dok svi ne nadju izvoe za hranu ili dok se ne popune svi izvori hrane
	for len(populationFreeArray) != 0 {
		// posto lista sadrzi sve id-jove uzmemo random iz nje
		if len(foodSourcesFreeArray) == 0 {
			break
		}
		var foodRandomId = rand.Intn(len(foodSourcesFreeArray))
		var foodSourceId = foodSourcesFreeArray[foodRandomId] //random id samo
		var individualRandomId = rand.Intn(len(populationFreeArray))
		var plebId = populationFreeArray[individualRandomId]
		var pleb = population[plebId]
		if foodSources[foodSourceId].occupied == 0 {
			foodSources[foodSourceId].occupiedBy = append(foodSources[foodSourceId].occupiedBy, pleb)
			//pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[0])
			foodSources[foodSourceId].occupied++
		} else {
			foodSources[foodSourceId].occupiedBy = append(foodSources[foodSourceId].occupiedBy, pleb)
			//pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[1])
			foodSources[foodSourceId].occupied++
			foodSourcesFreeArray = append(foodSourcesFreeArray[:foodRandomId], foodSourcesFreeArray[foodRandomId+1:]...)
		}
		populationFreeArray = append(populationFreeArray[:individualRandomId], populationFreeArray[individualRandomId+1:]...)
	}
}

func merge2(population []*Individual, foodSources []*FoodSource) {
	var populationFreeArray = makeArray(len(population))   //Lista id-ova
	var foodSourcesFreeArray = makeArray(len(foodSources)) //Lista id-ova
	fmt.Printf("%v\n", populationFreeArray)
	fmt.Printf("%v\n", foodSourcesFreeArray)
	for _, pleb := range population {
		// posto lista sadrzi sve id-jove uzmemo random iz nje
		if len(foodSourcesFreeArray) == 0 {
			break
		}
		fmt.Println("DUZINAAAAA", len(foodSourcesFreeArray))
		var randomFoodId = rand.Intn(len(foodSourcesFreeArray))
		var foodSourceId = foodSourcesFreeArray[randomFoodId]
		fmt.Println(foodSourceId)
		if foodSources[foodSourceId].occupied == 0 {
			pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[0])
			foodSources[foodSourceId].occupied++
		} else {
			pleb.resources = append(pleb.resources, foodSources[foodSourceId].resources[1])
			foodSources[foodSourceId].occupied++
			fmt.Println("UDJES LI ODJE")
			fmt.Println("ID:", foodSourceId)
			fmt.Println(len(foodSourcesFreeArray))
			fmt.Printf("Niz: %v \n", foodSourcesFreeArray)
			foodSourcesFreeArray = append(foodSourcesFreeArray[:randomFoodId], foodSourcesFreeArray[randomFoodId+1:]...)
			fmt.Printf("Niz: %v \n", foodSourcesFreeArray)
			fmt.Println(len(foodSourcesFreeArray))
		}
		//pleb.resources = append(pleb.resources, foodSources[foodSourceId])
		//DJE CESO OVO DOLE, MACINJI AKO JE PUN
		/*for !isFreeFoodSource(foodSources[foodSourceId]){
			fmt.Println("USO MICO")
			foodSourceId = rand.Intn(len(foodSources))
		}*/
		//fmt.Print(rand.Intn(len(foodSources)), ",")
		//fmt.Println("My animal is:", pleb)
	}
}
