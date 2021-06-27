package services

import (
	"../models"
)

func Day(population []*models.Individual, foodSources []*models.FoodSource) {
	Merge(population, foodSources)
	feedThePopulation(foodSources)
}
func ClearFoodSources(foodSources []*models.FoodSource) {
	for _, foodSource := range foodSources {
		clearSource(foodSource)
	}
}
