package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	// Server config
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"3000"`
	ShutdownTimeout time.Duration `default:"20s"`
}

var Cfg config

func LoadConfig() error {
	err := envconfig.Process("IRIS", &Cfg)
	if err != nil {
		return err
	}
	return nil
}
