package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
)

// TODO
func (r *Queries) GetTrainingById(ctx context.Context, training_id int) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddTraining(ctx context.Context, training *models.Training) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateTraining(ctx context.Context, training *models.Training) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteTraining(ctx context.Context, training_id int) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) GetTrainingInsrtuctionsList(ctx context.Context, training_id int) ([]*models.TrainingInstructions, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddTrainingInstruction(ctx context.Context, training *models.Training) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateTrainingInstruction(ctx context.Context, training *models.Training) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteTrainingInstruction(ctx context.Context, training_id int) (*models.Training, error) {
	return nil, fmt.Errorf("not implemented")
}
