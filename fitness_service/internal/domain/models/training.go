package models

import "time"

// Сама тренировка
type Training struct {
	Training_id         int
	Training_title      string
	Training_desc       string
	Training_user_level int
}

// Связующая таблица Пользователь-Тренировка (Many to many)
type TrainingPlan struct {
	Training_id int
	User_id     int
	Date        time.Time
}

// Инструкции к каждой тренировке (Many to one)
type TrainingInstructions struct {
	Training_id    int
	Training_order int
	Training_instr string
	Training_img   string
}
