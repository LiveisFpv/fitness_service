package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
)

// TODO
func (r *Queries) GetDishById(ctx context.Context, dish_id int) (*models.Dish, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddDish(ctx context.Context, dish *models.Dish) (*models.Dish, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateDish(ctx context.Context, dish *models.Dish) (*models.Dish, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteDish(ctx context.Context, dish_id int) (*models.Dish, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteDietPlan(ctx context.Context, dish_id, user_id int) (*models.DietPlan, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) GetRecipesList(ctx context.Context, dish_id int) ([]*models.Recipe, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) AddRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) UpdateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	return nil, fmt.Errorf("not implemented")
}
func (r *Queries) DeleteRecipe(ctx context.Context, recipe_id int) (*models.Recipe, error) {
	return nil, fmt.Errorf("not implemented")
}
