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

	return nil, fmt.Errorf("not implemented")
}

// TODO
func (r *Queries) GetDayPlan(ctx context.Context, user_id int, date time.Time) (*models.DayPlan, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Queries) GetTrainingListByWeek(ctx context.Context, user_id int, date time.Time) {
	sqlStatement := "SELECT * from training t JOIN train_plan p on p.training_id = t.training_id WHERE p.user_id = $1 and t.date BETWEEN $2::date AND ($2::date + INTERVAL '6 days')"

}
