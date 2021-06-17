package main

//Logika za dodavanje hrane
func feedThePopulation(foodSources []*FoodSource) {
	for _, foodSource := range foodSources {
		giveFoodToPleb(foodSource)
	}
}

func giveFoodToPleb(foodSource *FoodSource) {
	if foodSource.occupied == 0 {
		clearSource(foodSource)
		return
	}
	if foodSource.occupied == 1 {
		foodSource.occupiedBy[0].Health = len(foodSource.resources) * 100
		clearSource(foodSource)
		return
	}
	if foodSource.occupiedBy[0].TypeOfIndividual == GOOD && foodSource.occupiedBy[1].TypeOfIndividual == GOOD {
		var food = len(foodSource.resources) / 2.0 * 100.0
		foodSource.occupiedBy[0].Health = food
		foodSource.occupiedBy[1].Health = food
	} else if foodSource.occupiedBy[0].TypeOfIndividual == BAD && foodSource.occupiedBy[1].TypeOfIndividual == GOOD {
		var food = len(foodSource.resources) / 2.0 * 100.0
		foodSource.occupiedBy[0].Health = food + food/2
		foodSource.occupiedBy[1].Health = food / 2
	} else if foodSource.occupiedBy[0].TypeOfIndividual == GOOD && foodSource.occupiedBy[1].TypeOfIndividual == BAD {
		var food = len(foodSource.resources) / 2.0 * 100.0
		foodSource.occupiedBy[0].Health = food / 2
		foodSource.occupiedBy[1].Health = food + food/2
	} else {
		var food = len(foodSource.resources) / 2.0 * 100.0
		foodSource.occupiedBy[0].Health = food / 2
		foodSource.occupiedBy[0].Health = food / 2
	}
	clearSource(foodSource)
}

func clearSource(foodSource *FoodSource) {
	foodSource.occupiedBy = nil
	foodSource.occupied = 0
}
