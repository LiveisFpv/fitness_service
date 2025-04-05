package app

import (
	"context"
	"fitness_service/internal/grpc/fitnessgrpc"
	"fitness_service/internal/services/fitness"
	"fitness_service/internal/storage"
	"time"

	"github.com/sirupsen/logrus"
)

type App struct {
	GRPCServer *fitnessgrpc.App
	Storage    storage.UserModel
}

// Constructor APP creates gRPCServer, storage
func New(
	ctx context.Context,
	log *logrus.Logger,
	grpcPort int,
	dsn string,
	tokenTTL time.Duration,
) *App {
	storage, err := storage.NewStorage(ctx, dsn, log)
	if err != nil {
		panic(err)
	}

	//Todo service
	userService := fitness.New(log, storage, tokenTTL)
	grpcApp := fitnessgrpc.New(log, userService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
		Storage:    storage,
	}
}
