package grpc

import (
	"fmt"
	"github.com/mirage-deadline/sso/internal/grpc/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	log        *zap.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *zap.Logger, port int) *App {
	grpcServer := grpc.NewServer()
	auth.Register(grpcServer)

	return &App{
		log:        log,
		gRPCServer: grpcServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	log := a.log.With(
		zap.String("op", op),
		zap.Int("port", a.port),
	)
	log.Info("Starting grpc Server")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	log.Info("grpc server is running", zap.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (a *App) Stop() {
	const op = "grpacpp.Stop"
	a.log.With(zap.String("op", op)).Info("stopping grpc server", zap.Int("port", a.port))
	a.gRPCServer.GracefulStop()
}
