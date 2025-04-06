package fitnessgrpc

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"
	"time"

	fitness_v1 "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetUser(ctx context.Context, req *fitness_v1.ProfileRequest) (*fitness_v1.ProfileResponse, error) {
	if req.UserId < 1 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	user, err := s.user.GetUser(ctx, int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}

	resp := &fitness_v1.ProfileResponse{
		UserFirstname:     user.User_firstName,
		UserLastname:      user.User_lastName,
		UserMiddlename:    user.User_middleName,
		UserBirthday:      user.User_birthday,
		UserHeight:        int32(user.User_height),
		UserWeight:        user.User_weight,
		UserFitnessTarget: user.User_fitness_target,
		UserSex:           user.User_sex,
		UserLevel:         int32(*user.User_level),
	}

	return resp, nil
}

func (s *serverAPI) CreateUser(ctx context.Context, req *fitness_v1.CreateProfileRequest) (*fitness_v1.ProfileResponse, error) {

	usrLevel := int(*req.UserLevel)
	user, err := s.user.CreateUser(ctx, &models.User{
		User_id:             int(req.UserId),
		User_firstName:      req.UserFirstname,
		User_lastName:       req.UserLastname,
		User_middleName:     req.UserMiddlename,
		User_birthday:       req.UserBirthday,
		User_height:         int(req.UserHeight),
		User_weight:         req.UserWeight,
		User_fitness_target: req.UserFitnessTarget,
		User_sex:            req.UserSex,
		User_level:          &usrLevel,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprint(err))
	}

	userLevel2 := int32(*user.User_level)
	return &fitness_v1.ProfileResponse{
		UserFirstname:     user.User_firstName,
		UserLastname:      user.User_lastName,
		UserMiddlename:    user.User_middleName,
		UserBirthday:      user.User_birthday,
		UserHeight:        int32(user.User_height),
		UserWeight:        user.User_weight,
		UserFitnessTarget: user.User_fitness_target,
		UserSex:           user.User_sex,
		UserLevel:         userLevel2,
	}, nil
}

func (s *serverAPI) UpdateUser(ctx context.Context, req *fitness_v1.UpdateProfileRequest) (*fitness_v1.ProfileResponse, error) {
	usrLevel := int(*req.UserLevel)
	usrHeight := int(*req.UserHeight)
	user, err := s.user.UpdateUser(
		ctx,
		int(req.UserId),
		req.UserFirstname,
		req.UserLastname,
		req.UserMiddlename,
		req.UserBirthday,
		&usrHeight,
		req.UserWeight,
		req.UserFitnessTarget,
		req.UserSex,
		&usrLevel,
	)
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}

	return &fitness_v1.ProfileResponse{
		UserFirstname:     user.User_firstName,
		UserLastname:      user.User_lastName,
		UserMiddlename:    user.User_middleName,
		UserBirthday:      user.User_birthday,
		UserHeight:        int32(user.User_height),
		UserWeight:        user.User_weight,
		UserFitnessTarget: user.User_fitness_target,
		UserSex:           user.User_sex,
		UserLevel:         int32(*user.User_level),
	}, nil
}

func (s *serverAPI) GetPlanDishes(ctx context.Context, req *fitness_v1.GetPlanDishesRequest) (resp *fitness_v1.PlanDishesResponse, err error) {
	parsedTime, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, err
	}
	dishProgramm, err := s.user.GetPlanDishes(ctx, int(req.UserId), parsedTime)
	if err != nil {
		return nil, err
	}

	var response []*fitness_v1.PlanDishes

	for _, dish := range dishProgramm {
		response = append(response, &fitness_v1.PlanDishes{
			DishesId:     int64(dish.Dish_id),
			DishesWeight: dish.Dish_weight,
			DishesTitle:  dish.Dish_title,
			Time:         dish.Dish_time,
			Kcal:         dish.Dish_kcal,
			Fat:          dish.Dish_fats,
			Protein:      dish.Dish_proteins,
			Carbs:        dish.Dish_carbs,
			Description:  dish.Dish_desc,
			Date:         dish.Date.Format(time.RFC3339),
		})
	}

	return &fitness_v1.PlanDishesResponse{
		Data: response,
	}, nil
}
func (s *serverAPI) GetPlanTrain(ctx context.Context, req *fitness_v1.GetPlanTrainRequest) (resp *fitness_v1.PlanTrainResponse, err error) {
	parsedTime, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, err
	}

	trainProgramm, err := s.user.GetPlanTrain(ctx, int(req.UserId), parsedTime)
	if err != nil {
		return nil, err
	}

	var response []*fitness_v1.PlanTrain
	for _, train := range trainProgramm {
		response = append(response, &fitness_v1.PlanTrain{
			TrainId:          int64(train.Training_id),
			TrainTitle:       train.Training_title,
			TrainDescription: train.Training_desc,
			UserLevel:        int64(train.Training_user_level),
			Date:             train.Date.Format(time.RFC3339),
		})
	}

	return &fitness_v1.PlanTrainResponse{
		Data: response,
	}, nil
}

func (s *serverAPI) GetHistory(ctx context.Context, req *fitness_v1.GetHistoryRequest) (resp *fitness_v1.HistoryResponse, err error) {
	parsedTime, err := time.Parse(time.RFC3339, req.Date)
	if err != nil {
		return nil, err
	}

	history, err := s.user.GetWeightHistoryList(ctx, int(req.UserId), parsedTime)
	if err != nil {
		return nil, err
	}

	var response []*fitness_v1.History
	for _, train := range history {
		response = append(response, &fitness_v1.History{
			UserWeight: float64(train.Weight),
			Date:       train.Date.Format(time.RFC3339),
		})
	}

	return &fitness_v1.HistoryResponse{
		Data: response,
	}, nil
}
