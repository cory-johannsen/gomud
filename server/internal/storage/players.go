package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/io"
	"github.com/cory-johannsen/gomud/internal/loader"
	log "github.com/sirupsen/logrus"
)

type Players struct {
	database  *Database
	loaders   *loader.Loaders
	equipment *Equipment
	npcs      *NPCs
	players   map[string]*domain.Player
}

func NewPlayers(database *Database, npcs *NPCs, loaders *loader.Loaders, equipment *Equipment) *Players {
	return &Players{
		database:  database,
		players:   make(map[string]*domain.Player),
		loaders:   loaders,
		equipment: equipment,
		npcs:      npcs,
	}
}

func (p *Players) CreatePlayer(ctx context.Context, name string, password string, data map[string]domain.Property, conn io.Connection) (*domain.Player, error) {
	log.Debugf("creating player %s", name)
	specData := p.npcs.PropertiesToData(data)
	encoded, err := json.Marshal(specData)
	if err != nil {
		log.Errorf("failed to marshal player data: %s", err)
		return nil, err
	}
	var id int
	err = p.database.Conn.QueryRow(ctx, "INSERT INTO players (name, password, data) VALUES ($1, $2, $3) RETURNING id", name, password, encoded).
		Scan(&id)
	if err != nil {
		log.Errorf("failed to insert player: %s", err)
		return nil, err
	}
	player := domain.NewPlayer(nil, name, password, data, conn)
	player.Id = &id
	player.Data = data
	p.players[name] = player
	return player, nil
}

func (p *Players) FetchPlayerById(ctx context.Context, id int, conn io.Connection) (*domain.Player, error) {
	log.Debugf("fetching player %d", id)

	for _, player := range p.players {
		if player.Id != nil && *player.Id == id {
			return player, nil
		}
	}
	var name, password, data string
	err := p.database.Conn.QueryRow(ctx, "SELECT name, password FROM players WHERE id = $1", id).
		Scan(&name, &password, &data)
	if err != nil {
		log.Errorf("failed to fetch player: %s", err)
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Errorf("failed to unmarshal player data: %s", err)
		return nil, err
	}
	spec := &PlayerSpec{
		Id:       &id,
		Password: password,
		CharacterSpec: domain.CharacterSpec{
			Name: name,
			Data: specProps,
		},
	}
	player := p.PlayerFromSpec(ctx, spec, conn)

	// todo: load equipment

	return player, nil
}

func (p *Players) FetchPlayerByName(ctx context.Context, name string, conn io.Connection) (*domain.Player, error) {
	log.Debugf("fetching player %s", name)

	if player, ok := p.players[name]; ok {
		return player, nil
	}
	var id int
	var password string
	var data string
	err := p.database.Conn.QueryRow(ctx, "SELECT id, password, data FROM players WHERE name = $1", name).
		Scan(&id, &password, &data)
	if err != nil {
		log.Errorf("failed to fetch player: %s", err)
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Errorf("failed to unmarshal player data: %s", err)
		return nil, err
	}
	props := p.npcs.DataToProperties(ctx, specProps)
	player := domain.NewPlayer(&id, name, password, props, conn)
	// Peril threshold is calculated from Grit Bonus
	player.Peril().Threshold = player.StatBonuses().Grit + 3
	p.players[name] = player
	return player, nil
}

func (p *Players) Exists(ctx context.Context, name string) (bool, error) {
	log.Debugf("checking if player %s exists", name)

	var count int
	row := p.database.Conn.QueryRow(ctx, "SELECT count(*) FROM players WHERE name = $1", name)
	err := row.Scan(&count)
	if err != nil {
		log.Errorf("failed to check if player exists: %s", err)
		return false, err
	}
	return count > 0, nil
}

func (p *Players) IsLoggedIn(ctx context.Context, name string, conn io.Connection) (bool, error) {
	log.Debugf("checking if player %s is logged in", name)
	player, err := p.FetchPlayerByName(ctx, name, conn)
	if err != nil {
		return false, nil
	}
	return player.LoggedIn, nil
}

type PlayerSpec struct {
	domain.CharacterSpec
	Id       *int
	Password string
}

func (p *Players) SpecFromPlayer(player *domain.Player) *PlayerSpec {
	data := p.npcs.PropertiesToData(player.Data)
	spec := &PlayerSpec{
		Id:       player.Id,
		Password: player.Password,
		CharacterSpec: domain.CharacterSpec{
			Name: player.Name,
			Data: data,
		},
	}
	return spec
}

func (p *Players) PlayerFromSpec(ctx context.Context, spec *PlayerSpec, conn io.Connection) *domain.Player {
	data := p.npcs.DataToProperties(ctx, spec.Data)
	return domain.NewPlayer(spec.Id, spec.Name, spec.Password, data, conn)
}

func (p *Players) StorePlayer(ctx context.Context, player *domain.Player, conn io.Connection) (*domain.Player, error) {
	log.Debugf("storing player %s", player.Name)
	if player.Id == nil {
		return p.CreatePlayer(ctx, player.Name, player.Password, player.Data, conn)
	}
	data := p.SpecFromPlayer(player).Data
	encoded, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	tag, err := p.database.Conn.Exec(ctx, "UPDATE players SET data = $1 WHERE id = $2", encoded, player.Id)
	if err != nil {
		return nil, err
	}
	if tag.RowsAffected() != 1 {
		return nil, errors.New("failed to store player")
	}
	return player, nil
}
