package main

type FoodType int

const (
	HEALTHY   FoodType = 0
	UNHEALTHY FoodType = 1
)

type Food struct {
	FoodType FoodType `json:"food-type"`
	Quantity int      `json:"quantity"`
}

func newFood(foodType FoodType, quantity int) *Food {
	food := new(Food)
	food.FoodType = foodType
	food.Quantity = quantity
	return food
}

type FoodSource struct {
	resources  []*Food
	occupiedBy []*Individual
	occupied   int
}
