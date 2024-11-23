package grpcapp

import (
	"fmt"
	"log/slog"
	"net"

	authgrpc "github.com/LavaJover/auth-project/sso/internal/grpc/auth"
	"google.golang.org/grpc"
)

type App struct{
	log *slog.Logger
	gRPCServer *grpc.Server
	port int
}

func New(
	log *slog.Logger,
	port int,
) *App{
	gRPCServer := grpc.NewServer()
	authgrpc.Register(gRPCServer)

	return &App{
		log: log,
		gRPCServer: gRPCServer,
		port: port,
	}
}

func (app *App) MustRun(){
	if err:=app.Run(); err != nil{
		panic(err)
	}
}

func (app *App) Run() error{
	const op = "grpcapp.Run"
	log := app.log.With(slog.String("op", op))
	log.Info("Starting gRPC server...", slog.Int("port", app.port))
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", app.port))
	if err != nil{
		return err
	}
	log.Info("gRPC server is running", slog.String("address", lis.Addr().String()))
	if err := app.gRPCServer.Serve(lis); err != nil{
		return err
	}

	return nil
}

func (app *App) Stop() error{
	const op = "grpcapp.Stop"
	log := app.log.With(slog.String("op", op))

	log.Info("Server is stopping...")
	app.gRPCServer.GracefulStop()
	
	return nil
}