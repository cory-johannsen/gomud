package config

import (
	"github.com/caarlos0/env/v6"
	"os"
	"strings"
)

type Config struct {
	Port                   string `env:"PORT" envDefault:"7000"`
	DatabaseHost           string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabasePort           string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabaseName           string `env:"DATABASE_NAME" envDefault:"mud"`
	DatabaseUser           string `env:"DATABASE_USER" envDefault:"postgres"`
	DatabasePassword       string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	DatabaseMigrationsPath string `env:"DATABASE_MIGRATIONS_PATH" envDefault:"migrations"`
	AssetPath              string `env:"ASSET_PATH" envDefault:"assets"`
}

func NewConfigFromEnv() (*Config, error) {
	opts := env.Options{
		Environment: make(map[string]string),
	}
	for _, existingEnv := range os.Environ() {
		envVar := strings.Split(existingEnv, "=")
		opts.Environment[envVar[0]] = os.Getenv(envVar[0])
	}
	cfg := Config{}
	err := env.Parse(&cfg, opts)
	if err != nil {
		panic(err)
	}

	return &cfg, nil
}
