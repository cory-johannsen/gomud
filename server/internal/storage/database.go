package engine

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"log"
)

type Database struct {
	Conn *pgx.Conn
}

func NewDatabase(config *Config) (*Database, error) {
	conn, err := pgx.Connect(
		context.Background(),
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.DatabaseHost,
			config.DatabasePort,
			config.DatabaseUser,
			config.DatabasePassword,
			config.DatabaseName,
		),
	)
	if err != nil {
		return nil, err
	}

	log.Printf("connected to database %s\n", config.DatabaseName)
	log.Println("running migrations")

	url := fmt.Sprintf("file://%s", config.DatabaseMigrationsPath)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DatabaseUser, config.DatabasePassword, config.DatabaseHost, config.DatabasePort, config.DatabaseName)
	migration, err := migrate.New(url, connectionString)
	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &Database{Conn: conn}, nil
}
