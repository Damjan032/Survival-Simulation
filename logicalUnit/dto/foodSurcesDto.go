package dto

import "../models"

type FoodSourcesDto struct {
	FoodSources  []*models.FoodSource `json:"foodSources"`
	CurrentEpoch int                  `json:"currentEpoch"`
	FinalEpoch   int                  `json:"finalEpoch"`
	NumberOfGood int                  `json:"NumberOfGood"`
	NumberOfBad  int                  `json:"NumberOfBad"`
}
