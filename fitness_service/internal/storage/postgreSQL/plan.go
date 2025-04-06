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
	sqlStatement := `SELECT d.dish_id, dish_time, dish_title, dish_kcal, dish_proteins, dish_fats, dish_carbs, dish_desc, p.dish_weight, p.date from diet_plan p JOIN dishes d on d.dish_id = p.dish_id WHERE p.user_id=$1 and (TO_DATE(date, 'YYYY-MM-DD') between $2::date and ($2::date + INTERVAL '1 week')) ORDER BY date ASC`
	rows, err := r.pool.Query(ctx, sqlStatement, user_id, date)
	if err != nil {
		return nil, fmt.Errorf("can`t consturct dishes plan: %w", err)
	}
	response := []*models.DishProgramm{}
	for rows.Next() {
		resp := &models.DishProgramm{}
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
			&resp.Date,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		response = append(response, resp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}

func (r *Queries) GetPlanTrain(ctx context.Context, user_id int, date time.Time) ([]*models.TrainingProgramm, error) {
	sqlStatement := `SELECT t.training_id, training_title, training_desc, training_user_level, p.date from train_plan p JOIN training t on t.training_id = p.training_id WHERE p.user_id=$1 and (TO_DATE(date, 'YYYY-MM-DD') between $2::date and ($2::date + INTERVAL '1 week')) ORDER BY date ASC`
	rows, err := r.pool.Query(ctx, sqlStatement, user_id, date)
	if err != nil {
		return nil, fmt.Errorf("can`t consturct training plan: %w", err)
	}
	response := []*models.TrainingProgramm{}
	for rows.Next() {
		resp := &models.TrainingProgramm{}
		err := rows.Scan(
			&resp.Training_id,
			&resp.Training_title,
			&resp.Training_desc,
			&resp.Training_user_level,
			&resp.Date,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		response = append(response, resp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return response, nil
}
