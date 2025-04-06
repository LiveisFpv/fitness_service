package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

// GetUser implements Repository
func (r *Queries) GetUser(
	ctx context.Context,
	user_id int,
) (
	*models.User,
	error,
) {
	sqlStatement := `SELECT * FROM users WHERE user_id=$1`

	user := &models.User{}
	var time_temp time.Time
	err := r.pool.QueryRow(ctx, sqlStatement, user_id).Scan(
		&user.User_id,
		&user.User_firstName,
		&user.User_lastName,
		&user.User_middleName,
		&time_temp,
		&user.User_height,
		&user.User_weight,
		&user.User_fitness_target,
		&user.User_sex,
		&user.User_level,
	)
	user.User_birthday = time_temp.String()

	if err != nil {
		return nil, fmt.Errorf("couldn`t find user: %w", err)
	}

	return user, nil
}

// UpdateUser implements Repository
func (r *Queries) UpdateUser(
	ctx context.Context,
	user_id int,
	user_firstName *string,
	user_lastName *string,
	user_middleName *string,
	user_birthday *string,
	user_height *int,
	user_weight *float64,
	user_fitness_target *string,
	user_sex *bool,
	user_level *int,
) (
	*models.User,
	error,
) {
	sqlStatement := `UPDATE users SET user_firstName=$2, user_lastName=$3, user_middleName=$4, user_birthday=$5, user_height=$6, user_weight=$7, user_fitness_target=$8, user_sex=$9, user_level=$10 WHERE user_id=$1`

	old_user := &models.User{}

	old_user, err := r.GetUser(ctx, user_id)
	if err != nil {
		return nil, err
	}

	if user_firstName == nil {
		user_firstName = &old_user.User_firstName
	}

	if user_middleName == nil {
		user_middleName = old_user.User_middleName
	}
	if user_birthday == nil {
		user_birthday = &old_user.User_birthday
	}
	if user_height == nil {
		user_height = &old_user.User_height
	}
	if user_weight == nil {
		user_weight = &old_user.User_weight
	}
	if user_fitness_target == nil {
		user_fitness_target = &old_user.User_fitness_target
	}
	if user_sex == nil {
		user_sex = &old_user.User_sex
	}
	if user_level == nil {
		user_level = old_user.User_level
	}

	_, err = r.pool.Exec(
		ctx,
		sqlStatement,
		user_id,
		user_firstName,
		user_lastName,
		user_middleName,
		user_birthday,
		user_height,
		user_weight,
		user_fitness_target,
		user_sex,
		user_level,
	)

	if err != nil {
		return nil, fmt.Errorf("can`t update user profile: %w", err)
	}
	user, err := r.GetUser(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("can`t find user: %w", err)
	}
	return user, nil
}

// CreateUser implements Repository
// func (r *Queries) CreateUser(
// 	ctx context.Context,
// 	user *models.User) (
// 	*models.User,
// 	error,
// ) {
// 	sqlStatement := `INSERT INTO users (user_id, user_firstName, user_lastName, user_middleName, user_birthday, user_height, user_weight, user_fitness_target, user_sex, user_level) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`

// 	user_id := 0
// 	err := r.pool.QueryRow(ctx, sqlStatement,
// 		user.User_id,
// 		user.User_firstName,
// 		user.User_lastName,
// 		user.User_middleName,
// 		user.User_birthday,
// 		user.User_height,
// 		user.User_weight,
// 		user.User_fitness_target,
// 		user.User_sex,
// 		user.User_level,
// 	).Scan(&user_id)

// 	if err != nil {
// 		return nil, fmt.Errorf("can`t create user with profile: %w", err)
// 	}

// 	user, err = r.GetUser(ctx, user_id)
// 	if err != nil {
// 		return nil, fmt.Errorf("can`t find user: %w", err)
// 	}

// 	return user, nil
// }

func (r *Queries) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	// Insert user
	sqlStatement := `INSERT INTO users (user_id, user_firstName, user_lastName, user_middleName, user_birthday, user_height, user_weight, user_fitness_target, user_sex, user_level) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`
	user_id := 0
	err = r.pool.QueryRow(ctx, sqlStatement,
		user.User_id,
		user.User_firstName,
		user.User_lastName,
		user.User_middleName,
		user.User_birthday,
		user.User_height,
		user.User_weight,
		user.User_fitness_target,
		user.User_sex,
		user.User_level,
	).Scan(&user_id)
	if err != nil {
		return nil, fmt.Errorf("can't create user: %w", err)
	}

	// Calculate recommended calories
	bmr, err := calculateBMR(*user)
	if err != nil {
		return nil, err
	}
	tdee := calculateTDEE(bmr, *user.User_level)
	targetCalories := adjustCalories(tdee, user.User_fitness_target)

	// Generate weekly plans
	err = generateDietPlan(ctx, tx, int64(user.User_id), targetCalories)
	if err != nil {
		return nil, fmt.Errorf("can't generate diet plan: %w", err)
	}

	err = generateTrainingPlan(ctx, tx, int64(user.User_id), *user.User_level)
	if err != nil {
		return nil, fmt.Errorf("can't generate training plan: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("can't commit transaction: %w", err)
	}

	return r.GetUser(ctx, user.User_id)
}

// Helper functions
func calculateBMR(user models.User) (float64, error) {
	today := time.Now()
	birthday, err := time.Parse("2006-01-02T15:04:05Z", user.User_birthday)
	if err != nil {
		return -1, err
	}

	age := today.Year() - birthday.Year()
	if today.YearDay() < birthday.YearDay() {
		age--
	}

	if user.User_sex { // Assuming true = female
		return 447.593 + (9.247 * user.User_weight) + (3.098 * float64(user.User_height)) - (4.330 * float64(age)), nil
	}
	return 88.362 + (13.397 * user.User_weight) + (4.799 * float64(user.User_height)) - (5.677 * float64(age)), nil
}

func calculateTDEE(bmr float64, activityLevel int) float64 {
	multipliers := map[int]float64{
		1: 1.2,
		2: 1.375,
		3: 1.55,
		4: 1.725,
		5: 1.9,
	}
	return bmr * multipliers[activityLevel]
}

func adjustCalories(tdee float64, target string) float64 {
	switch target {
	case "lose":
		return tdee - 500
	case "gain":
		return tdee + 500
	default:
		return tdee
	}
}

func generateDietPlan(ctx context.Context, tx pgx.Tx, userID int64, dailyCalories float64) error {
	// Get Monday of next week
	startDate := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))

	// Sample meal distribution (breakfast, lunch, snack, dinner)
	mealDistribution := []struct {
		dishTime   string
		percentage float64
	}{
		{"Завтрак", 0.25},
		{"Обед", 0.35},
		{"Полдник", 0.15},
		{"Ужин", 0.25},
	}

	for i := 0; i < 7; i++ {
		currentDate := startDate.AddDate(0, 0, i).Format("2006-01-02")

		for _, meal := range mealDistribution {
			// Get random dish for meal type
			var dish models.Dish
			err := tx.QueryRow(ctx,
				"SELECT dish_id, dish_kcal FROM dishes WHERE dish_time = $1 ORDER BY RANDOM() LIMIT 1",
				meal.dishTime,
			).Scan(&dish.Dish_id, &dish.Dish_kcal)

			if err != nil {
				return fmt.Errorf("can't select dish: %w", err)
			}

			// Calculate portion weight
			mealCalories := dailyCalories * meal.percentage
			portionWeight := (mealCalories * 100) / dish.Dish_kcal

			// Insert into diet_plan
			_, err = tx.Exec(ctx,
				"INSERT INTO diet_plan (dish_id, user_id, dish_weight, date) VALUES ($1, $2, $3, $4)",
				dish.Dish_id,
				userID,
				portionWeight,
				currentDate,
			)
			if err != nil {
				return fmt.Errorf("can't insert diet plan: %w", err)
			}
		}
	}
	return nil
}

func generateTrainingPlan(ctx context.Context, tx pgx.Tx, userID int64, userLevel int) error {
	startDate := time.Now().Truncate(24*time.Hour).AddDate(0, 0, -int(time.Now().Weekday()))

	// Get 3-5 suitable workouts based on user level
	var trainingIDs []int64
	rows, err := tx.Query(ctx,
		"SELECT training_id FROM training WHERE training_user_level BETWEEN $1 AND $2 LIMIT 5",
		userLevel-1,
		userLevel+1,
	)
	if err != nil {
		return fmt.Errorf("can't select trainings: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		if err := rows.Scan(&id); err != nil {
			return err
		}
		trainingIDs = append(trainingIDs, id)
	}

	// Distribute trainings through the week
	for i, trainingID := range trainingIDs {
		date := startDate.AddDate(0, 0, i%7).Format("2006-01-02")
		_, err = tx.Exec(ctx,
			"INSERT INTO train_plan (training_id, user_id, date) VALUES ($1, $2, $3)",
			trainingID,
			userID,
			date,
		)
		if err != nil {
			return fmt.Errorf("can't insert training plan: %w", err)
		}
	}
	return nil
}
