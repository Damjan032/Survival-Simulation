package main

type UnitType int

const (
	GOOD UnitType = 0
	BAD  UnitType = 1
)

type Individual struct {
	resources        []*Food
	typeOfIndividual UnitType
	health           int
}

func NewIndividual(typeOf UnitType) *Individual {
	i := new(Individual)
	i.health = 100
	i.typeOfIndividual = typeOf
	return i
}
