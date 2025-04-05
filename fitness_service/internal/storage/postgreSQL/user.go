package postgresql

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
)

// TODO sql

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
	err := r.pool.QueryRow(ctx, sqlStatement, user_id).Scan(
		&user.User_id,
		&user.User_firstName,
		&user.User_lastName,
		&user.User_middleName,
		&user.User_birthday,
		&user.User_height,
		&user.User_weight,
		&user.User_fitness_target,
		&user.User_sex,
		&user.User_hypertain,
		&user.User_diabet,
		&user.User_level,
	)

	if err != nil {
		return nil, fmt.Errorf("Couldn`t find user: %w", err)
	}

	return user, nil
}

// UpdateUser implements Repository
func (r *Queries) UpdateUser(
	ctx context.Context,
	user *models.User,
) (
	*models.User,
	error,
) {
	sqlStatement := `UPDATE users SET user_firstName=$2, user_lastName=$3, user_middleName=$4, user_birthday=$5, user_height=$6, user_weight=$7, user_fitness_target=$8, user_sex=$9, user_hypertain=$10, user_diabet=$11, user_level=$12 WHERE user_id=$1`

	_, err := r.pool.Exec(
		ctx,
		sqlStatement,
		user.User_id,
		user.User_firstName,
		user.User_lastName,
		user.User_middleName,
		user.User_birthday,
		user.User_height,
		user.User_weight,
		user.User_fitness_target,
		user.User_sex,
		user.User_hypertain,
		user.User_diabet,
		user.User_level,
	)
	if err != nil {
		return nil, fmt.Errorf("can`t update user profile: %w", err)
	}
	user, err = r.GetUser(ctx, user.User_id)
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
	sqlStatement := `INSERT INTO users (user_firstName, user_lastName, user_middleName, user_birthday, user_height, user_weight, user_fitness_target, user_sex, user_hypertain, user_diabet, user_level) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING user_id`

	user_id := 0
	err := r.pool.QueryRow(ctx, sqlStatement,
		user.User_firstName,
		user.User_lastName,
		user.User_middleName,
		user.User_birthday,
		user.User_height,
		user.User_weight,
		user.User_fitness_target,
		user.User_sex,
		user.User_hypertain,
		user.User_diabet,
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
