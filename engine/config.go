package engine

import "os"

const DefaultPort = "7000"

type Config struct {
	Port string `json:"port"`
}

func NewConfigFromEnv() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	return &Config{port}, nil
}
