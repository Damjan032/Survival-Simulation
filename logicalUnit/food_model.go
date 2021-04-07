package main

type FoodType int

const (
	HEALTHY   FoodType = 0
	UNHEALTHY FoodType = 1
)

type Food struct {
	foodType FoodType
	quantity int
}

func newFood(foodType FoodType, quantity int) *Food {
	food := new(Food)
	food.foodType = foodType
	food.quantity = quantity
	return food
}

type FoodSource struct {
	resources  []*Food
	occupiedBy []*Individual
	occupied   int
}
