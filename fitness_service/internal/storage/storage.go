package storage

import (
	"context"
	"fitness_service/internal/domain/models"
	postgresql "fitness_service/internal/storage/postgreSQL"
	"fmt"
	"time"

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

	// TODO
	// Планы тренировки и диеты
	GetPlan(ctx context.Context, user_id int, date time.Time) (*models.Plan, error)
	GetDayPlan(ctx context.Context, user_id int, date time.Time) (*models.DayPlan, error)

	GetTrainingById(ctx context.Context, training_id int) (*models.Training, error)
	AddTraining(ctx context.Context, training *models.Training) (*models.Training, error)
	UpdateTraining(ctx context.Context, training *models.Training) (*models.Training, error)
	DeleteTraining(ctx context.Context, training_id int) (*models.Training, error)

	AddTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error)
	UpdateTrainingPlan(ctx context.Context, training_plan *models.TrainingPlan) (*models.TrainingPlan, error)
	DeleteTrainingPlan(ctx context.Context, training_id, user_id int) (*models.TrainingPlan, error)

	// TODO
	GetDishById(ctx context.Context, dish_id int) (*models.Dish, error)
	AddDish(ctx context.Context, dish *models.Dish) (*models.Dish, error)
	UpdateDish(ctx context.Context, dish *models.Dish) (*models.Dish, error)
	DeleteDish(ctx context.Context, dish_id int) (*models.Dish, error)

	//TODO
	AddDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error)
	UpdateDietPlan(ctx context.Context, diet_plan *models.DietPlan) (*models.DietPlan, error)
	DeleteDietPlan(ctx context.Context, dish_id, user_id int) (*models.DietPlan, error)

	// Список рецептов для блюда
	// Per dish_id
	// TODO
	GetRecipesList(ctx context.Context, dish_id int) ([]*models.Recipe, error)
	AddRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	UpdateRecipe(ctx context.Context, recipe *models.Recipe) (*models.Recipe, error)
	DeleteRecipe(ctx context.Context, dish_id, recipe_order int) (*models.Recipe, error)

	// Список упражнений для тренировки
	// Per Training
	GetTrainingInsrtuctionsList(ctx context.Context, training_id int) ([]*models.TrainingInstructions, error)
	AddTrainingInstruction(ctx context.Context, instruction *models.TrainingInstructions) (*models.TrainingInstructions, error)
	UpdateTrainingInstruction(ctx context.Context, instruction *models.TrainingInstructions) (*models.TrainingInstructions, error)
	DeleteTrainingInstruction(ctx context.Context, training_id, training_order int) (*models.TrainingInstructions, error)

	// История весов
	// TODO
	GetWeightHistoryList(ctx context.Context) ([]*models.WeightHistory, error)
	AddWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error)
	UpdateWeightHistory(ctx context.Context, weight *models.WeightHistory) (*models.WeightHistory, error)
	DeleteWightHistory(ctx context.Context, user_id int, date time.Time) (*models.WeightHistory, error)

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
