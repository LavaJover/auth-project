package app

import (
	"log/slog"
	"time"

	grpcapp "github.com/LavaJover/auth-project/sso/app/grpc"
)


type App struct{
	GRPCSrv *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App{
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCSrv: grpcApp,
	}
}