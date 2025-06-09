package utils

import (
	"fmt"
	"log"
	"os"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type LoadEnvOptions struct {
	Dotenv    bool
	EnvPrefix string
}

func LoadEnv(v any, opt LoadEnvOptions) (err error) {
	if opt.Dotenv {
		err = godotenv.Load(".env")
		if err != nil {
			err = fmt.Errorf("failed to load env: %w", err)
			log.Println(err)
			return
		}
	}
	envOpt := env.Options{
		Prefix: opt.EnvPrefix,
	}
	if err = env.ParseWithOptions(v, envOpt); err != nil {
		return fmt.Errorf("failed to parse env: %w", err)
	}
	log.Println("env loaded", os.Getenv("GO_GIN_DATABASE_URL"))
	return nil
}
