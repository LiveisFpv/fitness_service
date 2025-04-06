package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
)

func (r *Queries) GetDishById(ctx context.Context, dish_id int) (*models.Dish, error) {
	sqlStatement := `SELECT * FROM dishes WHERE dish_id=$1`
	dish := &models.Dish{}
	err := r.pool.QueryRow(ctx, sqlStatement, dish_id).Scan(
		&dish.Dish_id,
		&dish.Dish_time,
		&dish.Dish_title,
		&dish.Dish_kcal,
		&dish.Dish_proteins,
		&dish.Dish_fats,
		&dish.Dish_carbs,
		&dish.Dish_desc,
	)
	if err != nil {
		return nil, fmt.Errorf("couldn`t find dish: %w", err)
	}
	return dish, nil
}
func (r *Queries) AddDish(ctx context.Context, dish *models.Dish) (*models.Dish, error) {
	sqlStatement := `INSERT INTO dishes (dish_time, dish_title, dish_kcal, dish_proteins, dish_fats, dish_carbs, dish_desc) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING dish_id`

	dish_id := 0
	err := r.pool.QueryRow(ctx, sqlStatement,
		dish.Dish_time,
		dish.Dish_title,
		dish.Dish_kcal,
		dish.Dish_proteins,
		dish.Dish_fats,
		dish.Dish_carbs,
		dish.Dish_desc,
	).Scan(&dish_id)
	if err != nil {
		return nil, fmt.Errorf("can`t create dish: %w", err)
	}

	new_dish, err := r.GetDishById(ctx, dish_id)
	if err != nil {
		return nil, err
	}

	return new_dish, nil
}
func (r *Queries) UpdateDish(ctx context.Context, dish *models.Dish) (*models.Dish, error) {
	sqlStatement := `UPDATE dishes SET dish_time=$2, dish_title=$3, dish_kcal=$4, dish_proteins=$5, dish_fats=$6, dish_carbs=$7, dish_desc=$8 WHERE dish_id=$1`
	_, err := r.pool.Exec(
		ctx, sqlStatement,
		dish.Dish_id,
		dish.Dish_time,
		dish.Dish_title,
		dish.Dish_kcal,
		dish.Dish_proteins,
		dish.Dish_fats,
		dish.Dish_carbs,
		dish.Dish_desc,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update dish: %w", err)
	}

	new_dish, err := r.GetDishById(ctx, dish.Dish_id)
	if err != nil {
		return nil, err
	}

	return new_dish, nil
}
func (r *Queries) DeleteDish(ctx context.Context, dish_id int) (*models.Dish, error) {
	sqlStatement := `DELETE FROM dishes WHERE dish_id=$1 RETURNING dish_id, dish_time, dish_title, dish_kcal, dish_proteins, dish_fats, dish_carbs, dish_desc`
	dish := &models.Dish{}
	err := r.pool.QueryRow(ctx, sqlStatement, dish_id).Scan(
		&dish.Dish_id,
		&dish.Dish_time,
		&dish.Dish_title,
		&dish.Dish_kcal,
		&dish.Dish_proteins,
		&dish.Dish_fats,
		&dish.Dish_carbs,
		&dish.Dish_desc,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete dish: %w", err)
	}
	return dish, nil
}
func (r *Queries) AddDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error) {
	sqlStatement := `INSERT INTO diet_plan (dish_id, user_id, dish_weight, date) VALUES ($1, $2, $3, $4) RETURNING dish_id, user_id, dish_weight, date`
	new_diet_plan := &models.DietPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		diet_plan.Dish_id,
		diet_plan.User_id,
		diet_plan.Dish_weight,
		diet_plan.Date,
	).Scan(
		&new_diet_plan.Dish_id,
		&new_diet_plan.User_id,
		&new_diet_plan.Dish_weight,
		&new_diet_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t add new diet plan: %w", err)
	}
	return new_diet_plan, nil
}
func (r *Queries) UpdateDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error) {
	sqlStatement := `UPDATE diet_plan SET dish_weight=$3, date=$4 WHERE dish_id=$1, user_id=$2 RETURNING dish_id, user_id, dish_weight, date`
	new_diet_plan := &models.DietPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		diet_plan.Dish_id,
		diet_plan.User_id,
		diet_plan.Dish_weight,
		diet_plan.Date,
	).Scan(
		&new_diet_plan.Dish_id,
		&new_diet_plan.User_id,
		&new_diet_plan.Dish_weight,
		&new_diet_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update diet plan: %w", err)
	}
	return new_diet_plan, nil
}
func (r *Queries) DeleteDietPlan(ctx context.Context, dish_id, user_id int) (*models.DietPlan, error) {
	sqlStatement := `DELETE FROM diet_plan WHERE dish_id=$1, user_id=$2 RETURNING dish_id, user_id, dish_weight, date`
	old_diet_plan := &models.DietPlan{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		dish_id,
		user_id,
	).Scan(
		&old_diet_plan.Dish_id,
		&old_diet_plan.User_id,
		&old_diet_plan.Dish_weight,
		&old_diet_plan.Date,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete diet plan: %w", err)
	}
	return old_diet_plan, nil
}
func (r *Queries) GetRecipesList(ctx context.Context, dish_id int) ([]*models.Recipe, error) {
	sqlStatement := `SELECT * from recipes where dish_id=$1`

	rows, err := r.pool.Query(ctx, sqlStatement, dish_id)
	if err != nil {
		return nil, fmt.Errorf("can`t find recipes list: %w", err)
	}

	recipes := []*models.Recipe{}
	for rows.Next() {
		recipe := &models.Recipe{}
		err := rows.Scan(
			&recipe.Dish_id,
			&recipe.Recipe_order,
			&recipe.Recipe_instruct,
			&recipe.Recipe_img,
		)
		if err != nil {
			return nil, fmt.Errorf("can`t process query result: %w", err)
		}
		recipes = append(recipes, recipe)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return recipes, nil
}
func (r *Queries) AddRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	sqlStatement := `INSERT INTO recipes (dish_id, recipe_order, recipe_instruct, recipe_img) VALUES ($1, $2, $3, $4) RETURNING dish_id, recipe_order, recipe_instruct, recipe_img`
	new_recipes := &models.Recipe{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		recipe.Dish_id,
		recipe.Recipe_order,
		recipe.Recipe_instruct,
		recipe.Recipe_img,
	).Scan(
		&new_recipes.Dish_id,
		&new_recipes.Recipe_order,
		&new_recipes.Recipe_instruct,
		&new_recipes.Recipe_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t add new recipe: %w", err)
	}
	return new_recipes, nil
}
func (r *Queries) UpdateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error) {
	sqlStatement := `UPDATE recipes SET recipe_instruct=$3, recipe_img=$4 WHERE dish_id=$1, recipe_order=$2 RETURNING dish_id, recipe_order, recipe_instruct, recipe_img`
	new_recipes := &models.Recipe{}
	err := r.pool.QueryRow(ctx, sqlStatement,
		recipe.Dish_id,
		recipe.Recipe_order,
		recipe.Recipe_instruct,
		recipe.Recipe_img,
	).Scan(
		&new_recipes.Dish_id,
		&new_recipes.Recipe_order,
		&new_recipes.Recipe_instruct,
		&new_recipes.Recipe_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update recipe: %w", err)
	}
	return new_recipes, nil
}
func (r *Queries) DeleteRecipe(ctx context.Context, dish_id, recipe_order int) (*models.Recipe, error) {
	sqlStatement := `DELETE FROM recipes WHERE dish_id=$1, recipe_order=$2 RETURNING dish_id, recipe_order, recipe_instruct, recipe_img`
	old_recipes := &models.Recipe{}
	err := r.pool.QueryRow(ctx, sqlStatement, dish_id, recipe_order).Scan(
		&old_recipes.Dish_id,
		&old_recipes.Recipe_order,
		&old_recipes.Recipe_instruct,
		&old_recipes.Recipe_img,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t delete recipe: %w", err)
	}
	return old_recipes, nil
}
