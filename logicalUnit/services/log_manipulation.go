package services

import (
	"../const"
	"../globals"
	"fmt"
	"log"
	"os"
)

func SetupLog() {
	logFile, err := os.OpenFile(_const.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	err3 := os.Truncate(_const.LogFilePath, 0)
	if err3 != nil {
		log.Fatalf("error truncate file: %v", err)
	}
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
}

func WriteStartInLog() {
	fmt.Printf("[INFO] NumberOfFoodSources: %d FinalEpochs: %d\n", globals.NumberOfFoodSources, globals.FinalEpoch)
	log.Printf("[INFO] CurrentEpoch: %d, FinalEpochs: %d, NumberOfFoodSources: %d, NumberOfGoods: %d, NumberOfBads: %d, Level: %s\n",
		globals.CurrentEpoch, globals.FinalEpoch, globals.NumberOfFoodSources, globals.Population.NumberOfGood, globals.Population.NumberOfBad,
		globals.Lvl)
}

func WriteCurrentEpochInLog() {
	log.Printf("[INFO] CurrentEpoch: %d, FinalEpochs: %d, NumberOfFoodSources: %d, NumberOfGoods: %d, NumberOfBads: %d, Level: %s\n",
		globals.CurrentEpoch, globals.FinalEpoch, globals.NumberOfFoodSources, globals.Population.NumberOfGood, globals.Population.NumberOfBad, globals.Lvl)
}
