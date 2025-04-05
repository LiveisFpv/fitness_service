package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"
)

// TODO
func (r *Queries) GetWeightHistoryList(ctx context.Context) ([]*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteWightHistory(ctx context.Context, user_id int, date time.Time) (*models.WeightHistory, error) {
	return nil, fmt.Errorf("not implemented")
}
