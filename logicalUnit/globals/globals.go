package globals

import "../models"

var Population *models.Population
var FoodSources []*models.FoodSource
var CurrentEpoch = 0
var FinalEpoch = 100
var NumberOfGoods = 5
var NumberOfBads = 1
var NumberOfFoodSources = 100
var BoostToSurvive = 10
var MutationPercent = 10
var MutationToGoodBoost = 10
var Lvls = []string{"easy", "normal", "hard"}
var Lvl = "easy"
