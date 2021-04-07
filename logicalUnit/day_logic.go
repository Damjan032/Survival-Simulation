package main

func day(population []*Individual, foodSources []*FoodSource) {
	merge(population, foodSources)
	feedThePopulation(foodSources)
}
