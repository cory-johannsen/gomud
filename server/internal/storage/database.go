package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	url2 "net/url"
)

type Database struct {
	Conn *pgx.Conn
}

func NewDatabase(config *config.Config) (*Database, error) {
	pw := url2.QueryEscape(config.DatabasePassword)
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", config.DatabaseUser, pw, config.DatabaseHost, config.DatabasePort, config.DatabaseName, config.DatabaseSchema)
	conn, err := pgx.Connect(
		context.Background(),
		connectionString,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("connected to database %s", config.DatabaseName)
	log.Println("running migrations")

	url := fmt.Sprintf("file://%s", config.DatabaseMigrationsPath)
	migration, err := migrate.New(url, connectionString)
	if err != nil {
		return nil, err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return nil, err
	}

	return &Database{Conn: conn}, nil
}
