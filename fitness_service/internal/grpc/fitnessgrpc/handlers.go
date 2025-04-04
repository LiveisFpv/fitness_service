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
	fitness, err := s.fitness.Get_Profile(ctx, int(req.UserId))
	if err != nil {
		return nil, status.Error(codes.NotFound, fmt.Sprint(err))
	}
	return &fitness_v1.ProfileResponse{}, nil
}
func (s *serverAPI) CreateProfile(ctx context.Context, req *fitness_v1.CreateProfileRequest) (*fitness_v1.ProfileResponse, error) {
	//TODO
	return &fitness_v1.ProfileResponse{}, nil
}
func (s *serverAPI) UpdateProfile(ctx context.Context, req *fitness_v1.UpdateProfileRequest) (*fitness_v1.ProfileResponse, error) {
	//TODO
	return &fitness_v1.ProfileResponse{}, nil
}
