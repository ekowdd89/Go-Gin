package utils

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type LoadEnvOptions struct {
	Dotenv    bool
	EnvPrefix string
}

func LoadEnv(v any, opt LoadEnvOptions) (err error) {
	if opt.Dotenv {
		err = godotenv.Load()
	}
	envOpt := env.Options{
		Prefix: opt.EnvPrefix,
	}
	if err = env.ParseWithOptions(v, envOpt); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	return nil
}
