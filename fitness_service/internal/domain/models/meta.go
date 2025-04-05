package models

import "time"

type Pagination struct {
	Current int
	Total   int
	Limit   int
}

type Filter struct {
	Field string
	Value string
}

type Sort struct {
	Direction string
	By        string
}

// Один прием пищи (тупо из-за нагрузки веса в связь)
type DishRation struct {
	Dish        Dish
	Dish_weight float64
}

// План на один день
type DayPlan struct {
	Dishes    []Dish
	Trainings []Training
	Date      time.Time
}

// План на (условно) неделю
type Plan struct {
	Days []DayPlan
}
