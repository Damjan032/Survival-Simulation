package controllers

import (
	"../globals"
	"../models"
	"../services"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func SetInitData(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	globals.Population = models.NewPopulation(1, 5)
	globals.FoodSources = services.InitFoodSources(1000)
	fmt.Println("Set init data.")
	fmt.Println("I: ", globals.CurrentEpoch)
	fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
	fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	fmt.Println("-------------------------------------------------------------------------")
}

func NextEpoch(c *gin.Context) {
	globals.CurrentEpoch += 1
	fmt.Println(globals.CurrentEpoch)
	fmt.Println(globals.FinalEpoch)
	fmt.Println(globals.CurrentEpoch > globals.FinalEpoch)
	if globals.CurrentEpoch > globals.FinalEpoch {
		fmt.Println("Ne moze vise")
		return
	}
	fmt.Println("I: ", globals.CurrentEpoch)
	services.Day(globals.Population.Members, globals.FoodSources)
	//printPlebs(population.members)
	services.Night(globals.Population, 10, 10)
	fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
	fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	fmt.Println("-------------------------------------------------------------------------")

}

func ToTheEnd(c *gin.Context) {
	for i := globals.CurrentEpoch + 1; i <= globals.FinalEpoch; i++ {
		fmt.Println("I: ", i)
		services.Day(globals.Population.Members, globals.FoodSources)
		//printPlebs(population.members)
		services.Night(globals.Population, 10, 10)
		fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
		fmt.Println("-------------------------------------------------------------------------")
	}

}
