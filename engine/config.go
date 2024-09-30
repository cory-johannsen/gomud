package engine

import "os"

const DefaultPort = "7000"

type Config struct {
	Port             string `json:"port"`
	DatabaseHost     string `json:"database_host" default:"localhost"`
	DatabasePort     string `json:"database_port" default:"5432"`
	DatabaseName     string `json:"database_name" default:"postgres"`
	DatabaseUser     string `json:"database_user" default:"postgres"`
	DatabasePassword string `json:"database_password" default:"password"`
}

func NewConfigFromEnv() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	return &Config{
		Port:             port,
		DatabaseHost:     dbHost,
		DatabasePort:     dbPort,
		DatabaseName:     dbName,
		DatabaseUser:     dbUser,
		DatabasePassword: dbPassword,
	}, nil
}
