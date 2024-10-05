package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"log"
)

type Players struct {
	database         *Database
	backgroundLoader *loader.BackgroundLoader
	teamLoader       *loader.TeamLoader
	players          map[string]*domain.Player
}

func NewPlayers(database *Database, backgroundLoader *loader.BackgroundLoader, teamLoader *loader.TeamLoader) *Players {
	return &Players{
		database:         database,
		players:          make(map[string]*domain.Player),
		backgroundLoader: backgroundLoader,
		teamLoader:       teamLoader,
	}
}

func (p *Players) CreatePlayer(ctx context.Context, name string, password string, data map[string]domain.Property) (*domain.Player, error) {
	encoded, err := json.Marshal(propertiesToData(data))
	if err != nil {
		return nil, err
	}
	var id int
	err = p.database.Conn.QueryRow(ctx, "INSERT INTO players (name, password, data) VALUES ($1, $2, $3) RETURNING id", name, password, encoded).Scan(&id)
	if err != nil {
		return nil, err
	}
	player := domain.NewPlayer(nil, name, password, data)
	player.Id = &id
	player.Data = data
	p.players[name] = player
	return player, nil
}

func (p *Players) FetchPlayerById(ctx context.Context, id int) (*domain.Player, error) {
	for _, player := range p.players {
		if player.Id != nil && *player.Id == id {
			return player, nil
		}
	}
	var name, password, data string
	err := p.database.Conn.QueryRow(ctx, "SELECT name, password FROM players WHERE id = $1", id).Scan(&name, &password, &data)
	if err != nil {
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Printf("failed to unmarshal player data: %s", err)
		return nil, err
	}
	spec := &PlayerSpec{
		Id:       &id,
		Name:     name,
		Password: password,
		Data:     specProps,
	}
	player := p.PlayerFromSpec(spec)

	return player, nil
}

func (p *Players) FetchPlayerByName(ctx context.Context, name string) (*domain.Player, error) {
	if player, ok := p.players[name]; ok {
		return player, nil
	}
	var id int
	var password string
	var data string
	err := p.database.Conn.QueryRow(ctx, "SELECT id, password, data FROM players WHERE name = $1", name).Scan(&id, &password, &data)
	if err != nil {
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Printf("failed to unmarshal player data: %s", err)
		return nil, err
	}
	props := p.dataToProperties(specProps)
	player := domain.NewPlayer(&id, name, password, props)
	p.players[name] = player
	return player, nil
}

func (p *Players) Exists(ctx context.Context, name string) (bool, error) {
	var count int
	row := p.database.Conn.QueryRow(ctx, "SELECT count(*) FROM players WHERE name = $1", name)
	err := row.Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

type PlayerSpec struct {
	Id       *int
	Name     string
	Password string
	Data     map[string]interface{}
}

func SpecFromPlayer(player *domain.Player) *PlayerSpec {
	data := propertiesToData(player.Data)
	p := &PlayerSpec{
		Id:       player.Id,
		Name:     player.Name,
		Password: player.Password,
		Data:     data,
	}
	return p
}

func propertiesToData(props map[string]domain.Property) map[string]interface{} {
	data := make(map[string]interface{})
	for k, v := range props {
		// background needs to be serialized to a spec to remove the traits
		if k == domain.BackgroundProperty {
			data[k] = domain.SpecFromBackground(v.(*domain.Background))
			continue
		}
		// team needs to be serialized to a spec to remove the traits and jobs
		if k == domain.TeamProperty {
			data[k] = domain.SpecFromTeam(v.(*domain.Team))
			continue
		}
		// Stats is directly serializable
		data[k] = v.Value()
	}
	return data
}

func (p *Players) dataToProperties(data map[string]interface{}) map[string]domain.Property {
	props := make(map[string]domain.Property)
	for k, v := range data {
		switch k {
		case domain.BackgroundProperty:
			background, err := p.backgroundLoader.GetBackground(v.(string))
			if err != nil {
				log.Printf("failed to load background %s: %s", v.(string), err)
				continue
			}
			if background == nil {
				log.Printf("background %s not found", v.(string))
				continue
			}
			props[k] = background
		case domain.TeamProperty:
			teamName := v.(map[string]interface{})["name"].(string)
			team, err := p.teamLoader.GetTeam(teamName)
			if err != nil {
				log.Printf("failed to load team %s: %s", v.(string), err)
				continue
			}
			if team == nil {
				log.Printf("team %s not found", v.(string))
				continue
			}
			props[k] = team
		case domain.StatsProperty:
			stats := &domain.Stats{
				Fighting: v.(map[string]interface{})["fighting"].(int),
				Muscle:   v.(map[string]interface{})["muscle"].(int),
				Speed:    v.(map[string]interface{})["speed"].(int),
				Savvy:    v.(map[string]interface{})["savvy"].(int),
				Smarts:   v.(map[string]interface{})["smarts"].(int),
				Grit:     v.(map[string]interface{})["grit"].(int),
				Flair:    v.(map[string]interface{})["flair"].(int),
			}
			props[k] = stats
		default:
			log.Printf("unknown property %s: %v", k, v)
		}
	}
	return props
}

func (p *Players) PlayerFromSpec(spec *PlayerSpec) *domain.Player {
	data := p.dataToProperties(spec.Data)
	return domain.NewPlayer(spec.Id, spec.Name, spec.Password, data)
}

func (p *Players) StorePlayer(ctx context.Context, player *domain.Player) (*domain.Player, error) {
	if player.Id == nil {
		return p.CreatePlayer(ctx, player.Name, player.Password, player.Data)
	}
	encoded, err := json.Marshal(SpecFromPlayer(player).Data)
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
