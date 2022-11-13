package bootstrap

import (
	"context"

	"gochat-witai/pkg/config"

	server "gochat-witai/internal/platform"
)

func Run() error {
	err := config.LoadConfig()
	if err != nil {
		return err
	}

	ctx, srv := server.New(context.Background(), config.Cfg.Host, config.Cfg.Port, config.Cfg.ShutdownTimeout)
	return srv.Run(ctx)
}
