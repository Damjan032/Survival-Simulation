package services

import "../models"

func addFood(individual *models.Individual, foodType models.FoodType, quantity int) {
	individual.Resources = append(individual.Resources, models.NewFood(foodType, quantity))
}

func convertIndividualType(individual *models.Individual) {
	if individual.TypeOfIndividual == models.GOOD {
		individual.TypeOfIndividual = models.BAD
	} else {
		individual.TypeOfIndividual = models.GOOD
	}
}
