package fitnessgrpc

import (
	"context"
	"fmt"

	fitness_v1 "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetProfile(ctx context.Context, req *fitness_v1.ProfileRequest) (*fitness_v1.ProfileResponse, error) {
	if req.UserId < 1 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	//TODO Send request to Service
	user, err := s.user.GetUser(ctx, int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}

	return &fitness_v1.ProfileResponse{
		UserFirstname:     user.User_firstName,
		UserLastname:      user.User_lastName,
		UserMiddlename:    user.User_middleName,
		UserBirthday:      user.User_birthday,
		UserHeight:        int64(user.User_height),
		UserWeight:        user.User_weight,
		UserFitnessTarget: user.User_fitness_target,
		UserSex:           user.User_sex,
		UserHypertain:     user.User_hypertain,
		UserDiabet:        user.User_diabet,
		UserLevel:         &(int32(*user.User_level)),
	}, nil
}
func (s *serverAPI) CreateProfile(ctx context.Context, req *fitness_v1.CreateProfileRequest) (*fitness_v1.ProfileResponse, error) {
	//TODO
	return &fitness_v1.ProfileResponse{}, nil
}
func (s *serverAPI) UpdateProfile(ctx context.Context, req *fitness_v1.UpdateProfileRequest) (*fitness_v1.ProfileResponse, error) {
	//TODO
	return &fitness_v1.ProfileResponse{}, nil
}
