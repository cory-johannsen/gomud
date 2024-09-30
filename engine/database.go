package engine

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

type Database struct {
	conn *pgx.Conn
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
	return &Database{conn: conn}, nil
}
