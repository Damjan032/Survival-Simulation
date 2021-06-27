package controllers

import (
	"../dto"
	"../globals"
	"../models"
	"../services"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func SetGlobals(c *gin.Context) {
	numberOfGoods, err1 := strconv.Atoi(c.PostForm("number-of-goods"))
	if err1 != nil {
		fmt.Println("numberOfGoods cant found")
	} else {
		utils.SetNumberOfGoodIndividuals(numberOfGoods)
	}

	numberOfBads, err2 := strconv.Atoi(c.PostForm("number-of-bads"))
	if err2 != nil {
		fmt.Println("numberOfBads cant found")
	} else {
		utils.SetNumberOfBadIndividuals(numberOfBads)
	}

	numberOfFoodSources, err3 := strconv.Atoi(c.PostForm("number-of-food-sources"))
	if err3 != nil {
		fmt.Println("numberOfFoodSources cant found")
	} else {
		utils.SetNumberOfFoodSources(numberOfFoodSources)
	}

	numberOfEpoches, err4 := strconv.Atoi(c.PostForm("number-of-epoches"))
	if err4 != nil {
		fmt.Println("numberOfEpoches cant found")
	} else {
		utils.SetFinalEpoch(numberOfEpoches)
	}

	lvl := c.PostForm("simulator-lvl")
	utils.SetLvl(lvl)

	SetInitData(c)
	c.JSONP(200, gin.H{
		"message": "sve ok",
	})
}

func SetInitData(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	globals.CurrentEpoch = -1
	globals.Population = models.NewPopulation(globals.NumberOfBads, globals.NumberOfGoods)
	globals.FoodSources = services.InitFoodSources(globals.NumberOfFoodSources)
	fmt.Println("Set init data.")
	c.JSON(http.StatusOK, gin.H{
		"message": "Configuration set successfully.",
	})
}

func NextEpoch(c *gin.Context) {
	if globals.CurrentEpoch == -1 {
		fmt.Println("I: ", globals.CurrentEpoch)
		fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
		fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
		fmt.Println("-------------------------------------------------------------------------")
		services.Day(globals.Population.Members, globals.FoodSources)
		globals.CurrentEpoch += 1
		sendResult(c)
		return
	}
	if globals.CurrentEpoch >= globals.FinalEpoch {
		fmt.Println("Last epoch")
		sendResult(c)
		return
	}
	fmt.Println("I: ", globals.CurrentEpoch)
	services.ClearFoodSources(globals.FoodSources)
	services.Night(globals.Population)
	services.Day(globals.Population.Members, globals.FoodSources)
	//printPlebs(population.members)

	fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
	fmt.Println("SUM ", globals.Population.NumberOfBad+globals.Population.NumberOfGood)
	fmt.Println("-------------------------------------------------------------------------")
	globals.CurrentEpoch += 1
	sendResult(c)

}

func ToTheEnd(c *gin.Context) {
	if globals.CurrentEpoch >= globals.FinalEpoch {
		fmt.Println("Last epoch")
		sendResult(c)
		return
	}
	for i := globals.CurrentEpoch; i <= globals.FinalEpoch; i++ {
		fmt.Println("I: ", i)
		services.ClearFoodSources(globals.FoodSources)
		services.Night(globals.Population)
		services.Day(globals.Population.Members, globals.FoodSources)
	}
	fmt.Println("BAD", globals.Population.NumberOfBad, " GOOD", globals.Population.NumberOfGood)
	globals.CurrentEpoch = globals.FinalEpoch
	sendResult(c)

}

func CurrentData(c *gin.Context) {
	sendResult(c)
}
func sendResult(c *gin.Context) {
	foodSourcesDto := new(dto.FoodSourcesDto)
	foodSourcesDto.FoodSources = globals.FoodSources
	foodSourcesDto.CurrentEpoch = globals.CurrentEpoch
	foodSourcesDto.FinalEpoch = globals.FinalEpoch
	foodSourcesDto.NumberOfGood = globals.Population.NumberOfGood
	foodSourcesDto.NumberOfBad = globals.Population.NumberOfBad
	btResult, _ := json.Marshal(foodSourcesDto)
	c.Data(http.StatusOK, "application/json", btResult)
}
