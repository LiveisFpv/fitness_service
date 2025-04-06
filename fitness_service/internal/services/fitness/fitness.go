package fitness

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

// All methods
type UserController interface {
	GetUser(
		ctx context.Context,
		user_id int,
	) (
		*models.User,
		error,
	)
	CreateUser(
		ctx context.Context,
		profile *models.User,
	) (
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
	GetPlanTrain(ctx context.Context, user_id int, date time.Time) ([]*models.TrainingProgramm, error)
	GetPlanDishes(ctx context.Context, user_id int, date time.Time) ([]*models.DishProgramm, error)
	GetWeightHistoryList(ctx context.Context, user_id int, date time.Time) ([]*models.WeightHistory, error)
	GetRecipesList(ctx context.Context, dish_id int) ([]*models.Recipe, error)
	GetTrainingInsrtuctionsList(ctx context.Context, training_id int) ([]*models.TrainingInstructions, error)
}

type UserService struct {
	log            *logrus.Logger
	userController UserController
	tokenTTL       time.Duration
}

// Constructor service of User
func New(
	log *logrus.Logger,
	userController UserController,
	tokenTTL time.Duration,
) *UserService {
	return &UserService{
		log:            log,
		userController: userController,
		tokenTTL:       tokenTTL,
	}
}
func (u *UserService) GetUser(
	ctx context.Context,
	user_id int) (
	*models.User,
	error,
) {
	const op = "UserService.GetUser"
	log := u.log.WithFields(
		logrus.Fields{
			"op": op,
			"id": user_id,
		},
	)
	log.Info("Start Get by ID User")
	user, err := u.userController.GetUser(ctx, user_id)
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to get user with id %d", user_id), err)
		return nil, err
	}

	return user, nil
}

func (u *UserService) UpdateUser(
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
	const op = "UserService.UpdateUser"
	log := u.log.WithFields(
		logrus.Fields{
			"op":                  op,
			"user_id":             user_id,
			"user_firstName":      user_firstName,
			"user_lastName":       user_lastName,
			"user_middleName":     user_middleName,
			"user_birthday":       user_birthday,
			"user_height":         user_height,
			"user_weight":         user_weight,
			"user_fitness_target": user_fitness_target,
			"user_sex":            user_sex,
			"user_level":          user_level,
		},
	)
	log.Info("Start Update User")
	user, err := u.userController.UpdateUser(
		ctx,
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
		u.log.Error(fmt.Sprintf("failed to update user with id %d", user_id), err)
		return nil, err
	}

	return user, nil
}
func (u *UserService) CreateUser(
	ctx context.Context,
	user *models.User) (
	*models.User,
	error,
) {
	const op = "UserService.CreateUser"
	log := u.log.WithFields(
		logrus.Fields{
			"op":                  op,
			"user_id":             user.User_id,
			"user_firstName":      user.User_firstName,
			"user_lastName":       user.User_lastName,
			"user_middleName":     user.User_middleName,
			"user_birthday":       user.User_birthday,
			"user_height":         user.User_height,
			"user_weight":         user.User_weight,
			"user_fitness_target": user.User_fitness_target,
			"user_sex":            user.User_sex,
			"user_level":          user.User_level,
		},
	)
	log.Info("Start Create User")
	resp_user, err := u.userController.CreateUser(
		ctx,
		user,
	)
	if err != nil {
		u.log.Error(fmt.Sprintf("failed to create user"), err)
		return nil, err
	}

	return resp_user, nil
}

// GetPlanDishes implements fitnessgrpc.UserRepository.
func (u *UserService) GetPlanDishes(ctx context.Context, user_id int, date time.Time) ([]*models.DishProgramm, error) {
	const op = "UserService.GetPlanDishes"
	log := u.log.WithFields(logrus.Fields{})
	log.Info("Start GetPlanDishes")
	return u.userController.GetPlanDishes(ctx, user_id, date)
}

// GetPlanTrain implements fitnessgrpc.UserRepository.
func (u *UserService) GetPlanTrain(ctx context.Context, user_id int, date time.Time) ([]*models.TrainingProgramm, error) {
	const op = "UserService.GetPlanTrain"
	log := u.log.WithFields(logrus.Fields{})
	log.Info("Start GetPlanTrain")
	return u.userController.GetPlanTrain(ctx, user_id, date)
}

func (u *UserService) GetWeightHistoryList(ctx context.Context, user_id int, date time.Time) ([]*models.WeightHistory, error) {
	const op = "UserService.GetHistory"
	log := u.log.WithFields(logrus.Fields{})
	log.Info("Start GetHistory")
	return u.userController.GetWeightHistoryList(ctx, user_id, date)
}

// GetRecipe implements fitnessgrpc.UserRepository.
func (u *UserService) GetRecipe(ctx context.Context, dishes_id int) ([]*models.Recipe, error) {
	const op = "UserService.GetRecipe"
	log := u.log.WithFields(logrus.Fields{})
	log.Info("Start GetRecipe")
	return u.userController.GetRecipesList(ctx, dishes_id)
}

// GetTrainInstr implements fitnessgrpc.UserRepository.
func (u *UserService) GetTrainInstr(ctx context.Context, train_id int) ([]*models.TrainingInstructions, error) {
	const op = "UserService.TrainInstr"
	log := u.log.WithFields(logrus.Fields{})
	log.Info("Start TrainInstr")
	return u.userController.GetTrainingInsrtuctionsList(ctx, train_id)
}
