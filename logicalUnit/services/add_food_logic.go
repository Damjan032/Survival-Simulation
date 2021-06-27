package services

import "../models"

//Logika za dodavanje hrane
func feedThePopulation(foodSources []*models.FoodSource) {
	for _, foodSource := range foodSources {
		giveFoodToPleb(foodSource)
	}
}

func giveFoodToPleb(foodSource *models.FoodSource) {
	if foodSource.Occupied == 0 {
		//clearSource(foodSource)
		return
	}
	if foodSource.Occupied == 1 {
		foodSource.OccupiedBy[0].Health = len(foodSource.Resources) * 100
		//clearSource(foodSource)
		return
	}
	if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food
		foodSource.OccupiedBy[1].Health = food
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.BAD && foodSource.OccupiedBy[1].TypeOfIndividual == models.GOOD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food + food/2
		foodSource.OccupiedBy[1].Health = food / 2
	} else if foodSource.OccupiedBy[0].TypeOfIndividual == models.GOOD && foodSource.OccupiedBy[1].TypeOfIndividual == models.BAD {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[1].Health = food + food/2
	} else {
		var food = len(foodSource.Resources) / 2.0 * 100.0
		foodSource.OccupiedBy[0].Health = food / 2
		foodSource.OccupiedBy[0].Health = food / 2
	}
	//clearSource(foodSource)
}

func clearSource(foodSource *models.FoodSource) {
	foodSource.OccupiedBy = nil
	foodSource.Occupied = 0
}
