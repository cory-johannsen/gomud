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
	database *Database
	loaders  *loader.Loaders
	players  map[string]*domain.Player
}

func NewPlayers(database *Database, loaders *loader.Loaders) *Players {
	return &Players{
		database: database,
		players:  make(map[string]*domain.Player),
		loaders:  loaders,
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
		switch k {
		case domain.AlignmentProperty:
			data[k] = domain.SpecFromAlignment(v.(*domain.Alignment))
			continue
		case domain.ArchetypeProperty:
			data[k] = v.(*domain.Archetype).Name
			continue
		case domain.BackgroundProperty:
			data[k] = domain.SpecFromBackground(v.(*domain.Background))
			continue
		case domain.DrawbackProperty:
			data[k] = v.(*domain.Drawback).Name
			continue
		case domain.JobProperty:
			data[k] = v.(*domain.Job).Name
			continue
		case domain.TeamProperty:
			data[k] = domain.SpecFromTeam(v.(*domain.Team))
			continue
		case domain.TattooProperty:
			fallthrough
		case domain.DistinguishingMarkProperty:
			fallthrough
		case domain.BirthSeasonProperty:
			fallthrough
		case domain.StatsProperty:
			fallthrough
		default:
			data[k] = v.Value()
		}
	}
	return data
}

func (p *Players) dataToProperties(data map[string]interface{}) map[string]domain.Property {
	props := make(map[string]domain.Property)
	for k, v := range data {
		switch k {
		case domain.AgeProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
			continue
		case domain.AlignmentProperty:
			alignmentMap := v.(map[string]interface{})
			alignment, err := p.loaders.AlignmentLoader.GetAlignment(alignmentMap["order"].(string))
			if err != nil {
				log.Printf("failed to load alignment %s: %s", alignmentMap["order"].(string), err)
				continue
			}
			if alignment == nil {
				log.Printf("alignment %s not found", alignmentMap["order"].(string))
				continue
			}
			props[k] = alignment
		case domain.ArchetypeProperty:
			archetype, err := p.loaders.ArchetypeLoader.GetArchetype(v.(string))
			if err != nil {
				log.Printf("failed to load archetype %s: %s", v.(string), err)
				continue
			}
			if archetype == nil {
				log.Printf("archetype %s not found", v.(string))
				continue
			}
			props[k] = archetype
		case domain.BackgroundProperty:
			background, err := p.loaders.BackgroundLoader.GetBackground(v.(map[string]interface{})["Name"].(string))
			if err != nil {
				log.Printf("failed to load background %s: %s", v.(string), err)
				continue
			}
			if background == nil {
				log.Printf("background %s not found", v.(string))
				continue
			}
			props[k] = background
		case domain.BirthSeasonProperty:
			props[k] = domain.Season(v.(string))
		case domain.DistinguishingMarkProperty:
			marks := make(domain.DistinguishingMarks, 0)
			for _, mark := range v.([]interface{}) {
				marks = append(marks, domain.DistinguishingMark(mark.(string)))
			}
			props[k] = marks
		case domain.DrawbackProperty:
			drawback, err := p.loaders.AppearanceLoader.GetDrawback(v.(string))
			if err != nil {
				log.Printf("failed to load drawback %s: %s", v.(string), err)
				continue
			}
			if drawback == nil {
				log.Printf("drawback %s not found", v.(string))
				continue
			}
			props[k] = drawback
		case domain.JobProperty:
			job, err := p.loaders.JobLoader.GetJob(v.(string))
			if err != nil {
				log.Printf("failed to load job %s: %s", v.(string), err)
				continue
			}
			if job == nil {
				log.Printf("job %s not found", v.(string))
				continue
			}
			props[k] = job
		case domain.TeamProperty:
			teamName := v.(map[string]interface{})["Name"].(string)
			team, err := p.loaders.TeamLoader.GetTeam(teamName)
			if err != nil {
				log.Printf("failed to load team %s: %s", v.(string), err)
				continue
			}
			if team == nil {
				log.Printf("team %s not found", v.(string))
				continue
			}
			props[k] = team
		case domain.TattooProperty:
			description := v.(map[string]interface{})["Description"].(string)
			location := v.(map[string]interface{})["Location"].(string)
			tat := &domain.Tattoo{
				Description: description,
				Location:    domain.TattooLocation(location),
				Season:      domain.Season(v.(map[string]interface{})["Season"].(string)),
			}
			props[k] = tat
		case domain.StatsProperty:
			fighting := int(v.(map[string]interface{})["fighting"].(float64))
			muscle := int(v.(map[string]interface{})["muscle"].(float64))
			speed := int(v.(map[string]interface{})["speed"].(float64))
			savvy := int(v.(map[string]interface{})["savvy"].(float64))
			smarts := int(v.(map[string]interface{})["smarts"].(float64))
			grit := int(v.(map[string]interface{})["grit"].(float64))
			flair := int(v.(map[string]interface{})["flair"].(float64))
			stats := &domain.Stats{
				Fighting: fighting,
				Muscle:   muscle,
				Speed:    speed,
				Savvy:    savvy,
				Smarts:   smarts,
				Grit:     grit,
				Flair:    flair,
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
