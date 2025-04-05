package fitnessgrpc

import (
	"context"
	"fitness_service/internal/domain/models"

	fitness_v1 "github.com/LiveisFPV/fitness_v1/gen/go/fitness"
	"google.golang.org/grpc"
)

type serverAPI struct {
	fitness_v1.UnimplementedFitnessServer
	user UserRepository
}

// Methods needed for handlers on Service
type UserRepository interface {
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
}

// It how constructor but not constructor:В
func Register(gRPCServer *grpc.Server, user UserRepository) {
	fitness_v1.RegisterFitnessServer(gRPCServer, &serverAPI{user: user})
}
