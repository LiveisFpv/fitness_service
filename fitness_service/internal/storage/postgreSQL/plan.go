package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"
)

// TODO
func (r *Queries) GetPlan(ctx context.Context, user_id int, date time.Time) (*models.Plan, error) {
	sqlStatement_trainings := "SELECT "
	sqlStatement_trainings = sqlStatement_trainings
	return nil, fmt.Errorf("not implemented")
}

// TODO
func (r *Queries) GetDayPlan(ctx context.Context, user_id int, date time.Time) (*models.DayPlan, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Queries) GetTrainingListByWeek(ctx context.Context, user_id int, date time.Time) {
	sqlStatement := "SELECT * from training t JOIN train_plan p on p.training_id = t.training_id WHERE p.user_id = $1 and t.date BETWEEN $2::date AND ($2::date + INTERVAL '6 days')"
	sqlStatement = sqlStatement
}

func (r *Queries) GetPlanDishes(ctx context.Context, user_id int, date time.Time) ([]*models.DishProgramm, error) {
	sqlStatement := `
		SELECT d.dish_id, dish_time, dish_title, dish_kcal, dish_proteins, dish_fats, dish_carbs, dish_desc, p.dish_weight, p.date
		FROM diet_plan p
		JOIN dishes d ON d.dish_id = p.dish_id
		WHERE p.user_id = $1 
		AND TO_DATE(p.date, 'YYYY-MM-DD') BETWEEN $2::date AND ($2::date + INTERVAL '1 week')
		ORDER BY p.date ASC
	`

	// Выполняем запрос
	rows, err := r.pool.Query(ctx, sqlStatement, user_id, date)
	if err != nil {
		return nil, fmt.Errorf("can't construct dishes plan: %w", err)
	}
	defer rows.Close() // Закрываем rows после завершения работы

	var response []*models.DishProgramm

	// Чтение данных из результата запроса
	for rows.Next() {
		resp := &models.DishProgramm{}
		var dateStr string // Переменная для хранения строки даты из БД
		err := rows.Scan(
			&resp.Dish_id,
			&resp.Dish_time,
			&resp.Dish_title,
			&resp.Dish_kcal,
			&resp.Dish_proteins,
			&resp.Dish_fats,
			&resp.Dish_carbs,
			&resp.Dish_desc,
			&resp.Dish_weight,
			&dateStr, // Считываем строку с датой
		)
		if err != nil {
			return nil, fmt.Errorf("can't process query result: %w", err)
		}

		// Преобразуем строку даты в time.Time, если это нужно
		parsedDate, err := time.Parse("2006-01-02T15:04:05", dateStr) // Формат ISO 8601
		if err != nil {
			return nil, fmt.Errorf("can't parse date %s: %w", dateStr, err)
		}

		// Присваиваем преобразованную дату
		resp.Date = parsedDate

		// Добавляем в ответ
		response = append(response, resp)
	}

	// Проверка на ошибку после завершения чтения всех строк
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Queries) GetPlanTrain(ctx context.Context, user_id int, date time.Time) ([]*models.TrainingProgramm, error) {
	sqlStatement := `
		SELECT t.training_id, training_title, training_desc, training_user_level, p.date
		FROM train_plan p
		JOIN training t ON t.training_id = p.training_id
		WHERE p.user_id = $1
		AND TO_DATE(p.date, 'YYYY-MM-DD') BETWEEN $2::date AND ($2::date + INTERVAL '1 week')
		ORDER BY p.date ASC
	`

	// Выполняем запрос
	rows, err := r.pool.Query(ctx, sqlStatement, user_id, date)
	if err != nil {
		return nil, fmt.Errorf("can't construct training plan: %w", err)
	}
	defer rows.Close() // Закрываем rows после завершения работы

	var response []*models.TrainingProgramm

	// Чтение данных из результата запроса
	for rows.Next() {
		resp := &models.TrainingProgramm{}
		var dateStr string // Переменная для хранения строки даты из БД
		err := rows.Scan(
			&resp.Training_id,
			&resp.Training_title,
			&resp.Training_desc,
			&resp.Training_user_level,
			&dateStr, // Считываем строку с датой
		)
		if err != nil {
			return nil, fmt.Errorf("can't process query result: %w", err)
		}

		// Преобразуем строку даты в time.Time, если это нужно
		parsedDate, err := time.Parse("2006-01-02T15:04:05", dateStr) // Формат ISO 8601
		if err != nil {
			return nil, fmt.Errorf("can't parse date %s: %w", dateStr, err)
		}

		// Присваиваем преобразованную дату
		resp.Date = parsedDate

		// Добавляем в ответ
		response = append(response, resp)
	}

	// Проверка на ошибку после завершения чтения всех строк
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return response, nil
}
