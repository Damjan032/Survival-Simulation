package services

import (
	"../models"
)

func Day(population []*models.Individual, foodSources []*models.FoodSource) {
	Merge(population, foodSources)
	feedThePopulation(foodSources)
}
