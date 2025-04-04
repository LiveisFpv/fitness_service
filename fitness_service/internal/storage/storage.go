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

// Repository constructor
func NewRepository(
	pgxpool *pgxpool.Pool,
	log *logrus.Logger,
) Repository {
	return &repo{
		Queries: postgresql.New(pgxpool),
		log:     log,
		pool:    pgxpool,
	}
}

// Func for work with DB
type Repository interface {
	GetProfile(
		ctx context.Context,
		user_id int) (
		*models.Profile,
		error,
	)
	UpdateProfile(
		ctx context.Context,
		profile *models.Profile) (
		*models.Profile,
		error,
	)
	CreateProfile(
		ctx context.Context,
		profile *models.Profile) (
		*models.Profile,
		error,
	)
	Stop()
}

func NewStorage(ctx context.Context, dsn string, log *logrus.Logger) (Repository, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Проверяем подключение
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return NewRepository(pool, log), nil
}

func (r *repo) Stop() {
	r.Queries.Stop()
}
