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
		user_hypertain *bool,
		user_diabet *bool,
		user_level *int,
	) (
		*models.User,
		error,
	)
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
	user_hypertain *bool,
	user_diabet *bool,
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
			"user_hypertain":      user_hypertain,
			"user_diabet":         user_diabet,
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
		user_hypertain,
		user_diabet,
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
			"user_hypertain":      user.User_hypertain,
			"user_diabet":         user.User_diabet,
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
