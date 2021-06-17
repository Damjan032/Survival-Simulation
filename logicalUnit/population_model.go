package main

type Population2 struct {
	NumberOfBad  int `json:"number-of-bad"`
	NumberOfGood int `json:"number-of-good"`
}

type Population struct {
	Members      []*Individual `json:"members"`
	NumberOfBad  int           `json:"numberOfBad"`
	NumberOfGood int           `json:"numberOfGood"`
}

func NewPopulation2(numberOfBad int, numberOfGood int) *Population2 {
	population := new(Population2)
	population.NumberOfBad = numberOfBad
	population.NumberOfGood = numberOfGood
	return population
}

func NewPopulation(numberOfBad int, numberOfGood int) *Population {
	population := new(Population)
	population.Members = initMembersOfPopulation(numberOfGood, numberOfBad)
	population.NumberOfBad = numberOfBad
	population.NumberOfGood = numberOfGood
	return population
}
