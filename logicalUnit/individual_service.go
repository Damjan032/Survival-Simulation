package main

func addFood(individual *Individual, foodType FoodType, quantity int) {
	individual.resources = append(individual.resources, newFood(foodType, quantity))
}

func initMembersOfPopulation(numberOfGood int, numberOfBad int) []*Individual {
	var population []*Individual
	for i := 0; i < numberOfGood; i++ {
		population = append(population, NewIndividual(GOOD))
	}
	for i := 0; i < numberOfBad; i++ {
		population = append(population, NewIndividual(BAD))
	}
	return population
}

func convertType(individual *Individual) {
	if individual.typeOfIndividual == GOOD {
		individual.typeOfIndividual = BAD
	} else {
		individual.typeOfIndividual = GOOD
	}
}
