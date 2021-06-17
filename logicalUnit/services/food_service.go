package services

import "../models"

func initFoodSource(numberOfFoodUnit int) *models.FoodSource {
	foodSource := new(models.FoodSource)
	for i := 0; i < numberOfFoodUnit; i++ {
		foodSource.Resources = append(foodSource.Resources, models.NewFood(models.HEALTHY, 100))
		foodSource.Occupied = 0
	}
	return foodSource
}

func InitFoodSources(numberOfFoodSources int) []*models.FoodSource {
	var foodSources []*models.FoodSource
	for i := 0; i < numberOfFoodSources; i++ {
		foodSources = append(foodSources, initFoodSource(2))
	}
	return foodSources
}

//
