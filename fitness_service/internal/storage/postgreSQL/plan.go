package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"
)

// TODO
func (r *Queries) GetPlan(ctx context.Context, user_id int, date time.Time) (*models.Plan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) GetDayPlan(ctx context.Context, user_id int, date time.Time) (*models.DayPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
