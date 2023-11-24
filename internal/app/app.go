package app

import (
	grpcapp "github.com/mirage-deadline/sso/internal/app/grpc"
	"go.uber.org/zap"
	"time"
)

type App struct {
	GRPCServ *grpcapp.App
}

func New(log *zap.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	grpcApp := grpcapp.New(log, grpcPort)
	return &App{
		GRPCServ: grpcApp,
	}
}
