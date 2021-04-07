package main

type Population struct {
	members      []*Individual
	numberOfBad  int
	numberOfGood int
}

func NewPopulation(numberOfBad int, numberOfGood int) *Population {
	population := new(Population)
	population.members = initMembersOfPopulation(numberOfGood, numberOfBad)
	population.numberOfBad = numberOfBad
	population.numberOfGood = numberOfGood
	return population
}
