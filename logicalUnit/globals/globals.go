package globals

import (
	"../models"
	"go.mongodb.org/mongo-driver/mongo"
)

var Population *models.Population
var FoodSources []*models.FoodSource
var CurrentEpoch = -1
var FinalEpoch = 100
var NumberOfGoods = 5
var NumberOfBads = 1
var NumberOfFoodSources = 100
var BoostToSurvive = 10
var MutationPercent = 10
var MutationToGoodBoost = 10
var ErrorNormal = 0
var ErrorParallel = 0
var ParallelNightInterval = 16
var ParallelDayInterval = 16
var Lvls = []string{"easy", "normal", "hard"}
var MongoDBConnection *mongo.Database
var Lvl = "easy"
