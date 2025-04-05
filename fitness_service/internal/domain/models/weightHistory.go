package models

import "time"

type WeightHistory struct {
	User_id int
	Weight  float64
	Date    time.Time
}
