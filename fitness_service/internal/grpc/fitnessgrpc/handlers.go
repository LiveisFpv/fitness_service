package fitnessgrpc

import (
	"context"
	"fitness_service/internal/domain/models"
	"fmt"

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

func (s *serverAPI) CreateProfile(ctx context.Context, req *fitness_v1.CreateProfileRequest) (*fitness_v1.ProfileResponse, error) {

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

func (s *serverAPI) UpdateProfile(ctx context.Context, req *fitness_v1.UpdateProfileRequest) (*fitness_v1.ProfileResponse, error) {
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
