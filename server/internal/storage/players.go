package storage

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/io"
	"github.com/cory-johannsen/gomud/internal/loader"
	log "github.com/sirupsen/logrus"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
)

type Players struct {
	database *Database
	loaders  *loader.Loaders
	players  map[string]*domain.Player
	eventBus goeventbus.EventBus
}

func NewPlayers(database *Database, loaders *loader.Loaders, eventBus goeventbus.EventBus) *Players {
	return &Players{
		database: database,
		players:  make(map[string]*domain.Player),
		loaders:  loaders,
		eventBus: eventBus,
	}
}

func (p *Players) CreatePlayer(ctx context.Context, name string, password string, data map[string]domain.Property, conn io.Connection) (*domain.Player, error) {
	specData := propertiesToData(data)
	encoded, err := json.Marshal(specData)
	if err != nil {
		log.Errorf("failed to marshal player data: %s", err)
		return nil, err
	}
	var id int
	err = p.database.Conn.QueryRow(ctx, "INSERT INTO players (name, password, data) VALUES ($1, $2, $3) RETURNING id", name, password, encoded).Scan(&id)
	if err != nil {
		log.Errorf("failed to insert player: %s", err)
		return nil, err
	}
	player := domain.NewPlayer(nil, name, password, data, p.eventBus, conn)
	player.Id = &id
	player.Data = data
	p.players[name] = player
	return player, nil
}

func (p *Players) FetchPlayerById(ctx context.Context, id int, conn io.Connection) (*domain.Player, error) {
	for _, player := range p.players {
		if player.Id != nil && *player.Id == id {
			return player, nil
		}
	}
	var name, password, data string
	err := p.database.Conn.QueryRow(ctx, "SELECT name, password FROM players WHERE id = $1", id).Scan(&name, &password, &data)
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
		Name:     name,
		Password: password,
		Data:     specProps,
	}
	player := p.PlayerFromSpec(spec, conn)

	return player, nil
}

func (p *Players) FetchPlayerByName(ctx context.Context, name string, conn io.Connection) (*domain.Player, error) {
	if player, ok := p.players[name]; ok {
		return player, nil
	}
	var id int
	var password string
	var data string
	err := p.database.Conn.QueryRow(ctx, "SELECT id, password, data FROM players WHERE name = $1", name).Scan(&id, &password, &data)
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
	props := p.dataToProperties(specProps)
	player := domain.NewPlayer(&id, name, password, props, p.eventBus, conn)
	// Peril threshold is calculated from Grit Bonus
	player.Peril().Threshold = player.StatBonuses().Grit + 3
	p.players[name] = player
	return player, nil
}

func (p *Players) Exists(ctx context.Context, name string) (bool, error) {
	var count int
	row := p.database.Conn.QueryRow(ctx, "SELECT count(*) FROM players WHERE name = $1", name)
	err := row.Scan(&count)
	if err != nil {
		log.Errorf("failed to check if player exists: %s", err)
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
		case domain.ArchetypeProperty:
			data[k] = v.(*domain.Archetype).Name
		case domain.BackgroundProperty:
			data[k] = v.(*domain.Background).Name
		case domain.BackgroundTraitProperty:
			data[k] = v.(*domain.Trait).Name
		case domain.ConditionProperty:
			data[k] = v.(domain.Condition)
		case domain.ConsumedAdvancesProperty:
			data[k] = v
		case domain.DrawbackProperty:
			data[k] = v.(*domain.Drawback).Name
		case domain.JobProperty:
			data[k] = v.(*domain.Job).Name
		case domain.RoomProperty:
			data[k] = v.(*domain.Room).Name
		case domain.SkillRanksProperty:
			skillRanks := make(map[string][]string)
			for _, skill := range v.(domain.SkillRanks) {
				jobName := skill.Job.Name
				if _, ok := skillRanks[jobName]; !ok {
					skillRanks[jobName] = make([]string, 0)
				}
				skillRanks[jobName] = append(skillRanks[jobName], skill.Skill.Name)
			}
			data[k] = skillRanks
		case domain.TeamProperty:
			data[k] = domain.SpecFromTeam(v.(*domain.Team))
		case domain.TalentsProperty:
			talents := make([]string, 0)
			for _, talent := range v.(domain.Talents) {
				talents = append(talents, talent.Name)
			}
			data[k] = talents
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
			name := v.(string)
			background, err := p.loaders.BackgroundLoader.GetBackground(name)
			if err != nil {
				log.Printf("failed to load background %s: %s", v.(string), err)
				continue
			}
			if background == nil {
				log.Printf("background %s not found", v.(string))
				continue
			}
			props[k] = background
		case domain.BackgroundTraitProperty:
			name := v.(string)
			trait, err := p.loaders.TraitLoader.GetTrait(name)
			if err != nil {
				log.Printf("failed to load trait %s: %s", v.(string), err)
				continue
			}
			if trait == nil {
				log.Printf("trait %s not found", v.(string))
				continue
			}
			props[k] = trait
		case domain.BirthSeasonProperty:
			props[k] = domain.Season(v.(string))
		case domain.ConditionProperty:
			props[k] = domain.Condition(v.(string))
		case domain.ConsumedAdvancesProperty:
			consumedAdvances := make(domain.ConsumedAdvances)
			for _, advances := range v.(map[string]interface{}) {
				for _, advance := range advances.([]interface{}) {
					job := advance.(map[string]interface{})["Job"].(string)
					stat := advance.(map[string]interface{})["Stat"].(string)
					amount := advance.(map[string]interface{})["Amount"].(float64)
					if _, ok := consumedAdvances[job]; !ok {
						consumedAdvances[job] = make([]*domain.ConsumedAdvance, 0)
					}

					var consumedAdvance *domain.ConsumedAdvance
					for _, ca := range consumedAdvances[job] {
						if ca.Stat == stat {
							consumedAdvance = ca
							break
						}
					}
					if consumedAdvance == nil {
						consumedAdvances[job] = append(consumedAdvances[job], &domain.ConsumedAdvance{
							Job:    job,
							Stat:   stat,
							Amount: int(amount),
						})
					} else {
						consumedAdvance.Amount += int(amount)
					}
				}
			}
			props[k] = consumedAdvances
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
		case domain.ExperienceProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
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
		case domain.PerilProperty:
			peril := v.(map[string]interface{})
			threshold := peril["Threshold"].(float64)
			perilCondition := peril["Condition"].(float64)
			props[k] = &domain.Peril{
				Threshold: int(threshold),
				Condition: domain.PerilCondition(perilCondition),
			}
		case domain.RoomProperty:
			room := p.loaders.RoomLoader.GetRoom(v.(string))
			props[k] = room
		case domain.SkillRanksProperty:
			skillRanks := make(domain.SkillRanks, 0)
			for jobName, skills := range v.(map[string]interface{}) {
				job, err := p.loaders.JobLoader.GetJob(jobName)
				if err != nil {
					log.Printf("failed to load job %s: %s", jobName, err)
					continue
				}
				if job == nil {
					log.Printf("job %s not found", jobName)
					continue
				}
				for _, skillName := range skills.([]interface{}) {
					skill, err := p.loaders.SkillLoader.GetSkill(skillName.(string))
					if err != nil {
						log.Printf("failed to load skill %s: %s", skillName.(string), err)
						continue
					}
					if skill == nil {
						log.Printf("skill %s not found", skillName.(string))
						continue
					}
					skillRanks = append(skillRanks, &domain.SkillRank{
						Job:   job,
						Skill: skill,
					})
				}
			}
			props[k] = skillRanks
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
		case domain.TalentsProperty:
			talents := make(domain.Talents, 0)
			for _, talentName := range v.([]interface{}) {
				talent, err := p.loaders.TalentLoader.GetTalent(talentName.(string))
				if err != nil {
					log.Printf("failed to load talent %s: %s", talentName.(string), err)
					continue
				}
				if talent == nil {
					log.Printf("talent %s not found", talentName.(string))
					continue
				}
				talents = append(talents, talent)
			}
			props[k] = talents
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

func (p *Players) PlayerFromSpec(spec *PlayerSpec, conn io.Connection) *domain.Player {
	data := p.dataToProperties(spec.Data)
	return domain.NewPlayer(spec.Id, spec.Name, spec.Password, data, p.eventBus, conn)
}

func (p *Players) StorePlayer(ctx context.Context, player *domain.Player, conn io.Connection) (*domain.Player, error) {
	if player.Id == nil {
		return p.CreatePlayer(ctx, player.Name, player.Password, player.Data, conn)
	}
	data := SpecFromPlayer(player).Data
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
