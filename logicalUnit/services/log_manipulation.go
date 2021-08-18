package services

import (
	"../globals"
	"log"
	"os"
)

func SetupLog() {
	logFile, err := os.OpenFile(globals.LogFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	err3 := os.Truncate(globals.LogFilePath, 0)
	if err3 != nil {
		log.Fatalf("error truncate file: %v", err)
	}
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(logFile)
}

func WriteStartInLog() {

}

func WriteCurrentEpochInLog() {

}
