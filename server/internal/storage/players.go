package storage

import (
	"context"
	"encoding/json"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type Players struct {
	database *Database
}

func NewPlayers(database *Database) *Players {
	return &Players{database: database}
}

func (p *Players) CreatePlayer(ctx context.Context, name string, password string) (*domain.Player, error) {
	var id int
	err := p.database.Conn.QueryRow(ctx, "INSERT INTO players (name, password) VALUES (?, ?) RETURNING id", name, password).Scan()
	if err != nil {
		return nil, err
	}
	return domain.NewPlayer(id, name, password), nil
}

func (p *Players) FetchPlayer(ctx context.Context, id int) (*domain.Player, error) {
	var name, password string
	err := p.database.Conn.QueryRow(ctx, "SELECT name, password FROM players WHERE id = ?", id).Scan(&name, &password)
	if err != nil {
		return nil, err
	}
	return domain.NewPlayer(id, name, password), nil
}

func (p *Players) Exists(ctx context.Context, name string) (bool, error) {
	var count int
	err := p.database.Conn.QueryRow(ctx, "SELECT count(*) as count from FROM players WHERE name = ?", name).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (p *Players) StorePlayer(ctx context.Context, player *domain.Player) error {
	encoded, err := json.Marshal(player.Data)
	if err != nil {
		return err
	}
	_, err = p.database.Conn.Exec(ctx, "UPDATE players SET data = ? WHERE id = ?", encoded, player.Id)
	if err != nil {
		return err
	}
	return nil
}
