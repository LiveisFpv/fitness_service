package fitnessgrpc

import (
	"context"
	"fitness_service/internal/domain/models"

	fitness_v1 "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
	"google.golang.org/grpc"
)

type serverAPI struct {
	fitness_v1.UnimplementedFitnessServer
	user User
}

// Methods needed for handlers on Service
type User interface {
	GetProfile(
		ctx context.Context,
		user_id int,
	) (
		*models.Profile,
		error,
	)
	CreateProfile(
		ctx context.Context,
		profile *models.Profile,
	) (
		*models.Profile,
		error,
	)
	UpdateProfile(
		ctx context.Context,
		user_birthday *string,
		user_height *int,
		user_weight *float64,
		user_fitness_target *string,
		user_sex *bool,
		user_hypertain *bool,
		user_diabet *bool,
		user_level *int,
	) (
		profile *models.Profile,
		err error,
	)
}

// It how constructor but not constructor:В
func Register(gRPCServer *grpc.Server, user User) {
	fitness_v1.RegisterFitnessServer(gRPCServer, &serverAPI{user: user})
}
