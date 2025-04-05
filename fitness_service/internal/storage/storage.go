package storage

import (
	"context"
	"fitness_service/internal/domain/models"
	postgresql "fitness_service/internal/storage/postgreSQL"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

type repo struct {
	*postgresql.Queries
	log  *logrus.Logger
	pool *pgxpool.Pool
}

// UserModel constructor
func NewUserModel(
	pgxpool *pgxpool.Pool,
	log *logrus.Logger,
) UserModel {
	return &repo{
		Queries: postgresql.New(pgxpool),
		log:     log,
		pool:    pgxpool,
	}
}

// Func for work with DB
type UserModel interface {
	GetUser(
		ctx context.Context,
		user_id int) (
		*models.User,
		error,
	)
	UpdateUser(
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
		user_hypertain *bool,
		user_diabet *bool,
		user_level *int,
	) (
		*models.User,
		error,
	)
	CreateUser(
		ctx context.Context,
		user *models.User) (
		*models.User,
		error,
	)
	Stop()
}

func NewStorage(ctx context.Context, dsn string, log *logrus.Logger) (UserModel, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return NewUserModel(pool, log), nil
}

func (r *repo) Stop() {
	r.Queries.Stop()
}
