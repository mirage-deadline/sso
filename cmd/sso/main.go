package main

import (
	"github.com/mirage-deadline/sso/internal/app"
	"github.com/mirage-deadline/sso/internal/config"
	pkglogger "github.com/mirage-deadline/sso/pkg/logger"
)

func main() {
	cfg := config.NewConfig()
	logger := pkglogger.MustLogger(cfg.Env)
	defer logger.Sync()

	application := app.New(logger, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)
	application.GRPCServ.MustRun()

}
