package models

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
	Resources        []*Food           `json:"resources"`
	TypeOfIndividual UnitType          `json:"typeOfIndividual"`
	Health           int               `json:"health"`
	NightRes         ProductOfTheNight `json:"night-res"`
}

func NewIndividual(typeOf UnitType) *Individual {
	i := new(Individual)
	i.Health = 100
	i.TypeOfIndividual = typeOf
	i.NightRes = NEUTRAL
	return i
}
