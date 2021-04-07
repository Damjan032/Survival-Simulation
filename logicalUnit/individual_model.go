package main

type UnitType int

const (
	GOOD UnitType = 0
	BAD  UnitType = 1
)

type ProductOfTheNight int

const (
	DIED       ProductOfTheNight = -1
	NEUTRAL    ProductOfTheNight = 0
	REPRODUCED ProductOfTheNight = 1
)

type Individual struct {
	resources        []*Food
	typeOfIndividual UnitType
	health           int
	nightRes         ProductOfTheNight
}

func NewIndividual(typeOf UnitType) *Individual {
	i := new(Individual)
	i.health = 100
	i.typeOfIndividual = typeOf
	i.nightRes = NEUTRAL
	return i
}
