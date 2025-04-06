package models

import "time"

// Блюдо
type Dish struct {
	Dish_id       int
	Dish_time     string // Утро/день/вечер
	Dish_title    string
	Dish_kcal     float64
	Dish_proteins float64
	Dish_fats     float64
	Dish_carbs    float64
	Dish_desc     string
}

// Связь пользовател-блюдо (Many to many)
type DietPlan struct {
	Dish_id     int
	User_id     int
	Dish_weight float64
	Date        time.Time
}

type DishProgramm struct {
	Dish_id       int
	Dish_time     string
	Dish_title    string
	Dish_kcal     float64
	Dish_proteins float64
	Dish_fats     float64
	Dish_carbs    float64
	Dish_desc     string
	Dish_weight   float64
	Date          time.Time
}

// Инструкции к блюду (many to one)
type Recipe struct {
	Dish_id         int
	Recipe_order    int
	Recipe_instruct string
	Recipe_img      string
}
