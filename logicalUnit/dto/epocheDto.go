package dto

type Epoch struct {
	CurrentEpoch int `json:"CurrentEpoch"`
	NumberOfGood int `json:"NumberOfGood"`
	NumberOfBad  int `json:"NumberOfBad"`
}
