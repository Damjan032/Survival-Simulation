package main

func addFood(individual *Individual, foodType FoodType, quantity int) {
	individual.Resources = append(individual.Resources, newFood(foodType, quantity))
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
	if individual.TypeOfIndividual == GOOD {
		individual.TypeOfIndividual = BAD
	} else {
		individual.TypeOfIndividual = GOOD
	}
}
