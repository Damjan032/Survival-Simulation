package main

func initFoodSource(numberOfFoodUnit int) *FoodSource {
	foodSource := new(FoodSource)
	for i := 0; i < numberOfFoodUnit; i++ {
		foodSource.resources = append(foodSource.resources, newFood(HEALTHY, 100))
		foodSource.occupied = 0
	}
	return foodSource
}

func initFoodSources(numberOfFoodSources int) []*FoodSource {
	var foodSources []*FoodSource
	for i := 0; i < numberOfFoodSources; i++ {
		foodSources = append(foodSources, initFoodSource(2))
	}
	return foodSources
}

//
