package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"
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
func (r *Queries) CreateUser(
	ctx context.Context,
	user *models.User) (
	*models.User,
	error,
) {
	sqlStatement := `INSERT INTO users (user_id, user_firstName, user_lastName, user_middleName, user_birthday, user_height, user_weight, user_fitness_target, user_sex, user_level) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING user_id`

	user_id := 0
	err := r.pool.QueryRow(ctx, sqlStatement,
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
		return nil, fmt.Errorf("can`t create user with profile: %w", err)
	}

	user, err = r.GetUser(ctx, user_id)
	if err != nil {
		return nil, fmt.Errorf("can`t find user: %w", err)
	}

	return user, nil
}
