package domain

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Property interface {
	Value() interface{}
	String() string
}

const PropertyNotFound = "property not found"

const (
	ArchetypeProperty          = "archetype"
	BackgroundProperty         = "background"
	BirthSeasonProperty        = "birthSeason"
	DistinguishingMarkProperty = "distinguishingMark"
	DrawbackProperty           = "drawback"
	JobProperty                = "job"
	TeamProperty               = "team"
	TattooProperty             = "tattoo"
	StatsProperty              = "stats"
)

type Player struct {
	Id       *int
	Name     string
	Password string
	Data     map[string]Property
}

func NewPlayer(id *int, name string, password string, data map[string]Property) *Player {
	p := &Player{
		Id:       id,
		Name:     name,
		Password: password,
		Data:     data,
	}
	if p.Data == nil {
		p.Data = make(map[string]Property)
	}
	return p
}

func (p *Player) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (p *Player) GetProperty(key string) (Property, error) {
	result, ok := p.Data[key]
	if !ok {
		return nil, errors.New(PropertyNotFound)
	}
	return result, nil
}

func (p *Player) String() string {
	msg := fmt.Sprintf("Name: %s\n", p.Name)
	for k, v := range p.Data {
		switch k {
		case ArchetypeProperty:
			msg += fmt.Sprintf("  Archetype - %s\n", v.(*Archetype).Name)
			continue
		case BackgroundProperty:
			msg += fmt.Sprintf("  Background - \n\t%s\n\t%s\n", v.(*Background).Name, v.(*Background).Description)
			continue
		case BirthSeasonProperty:
			msg += fmt.Sprintf("  Birth Season - %s\n", v.(Season))
			continue
		case DistinguishingMarkProperty:
			msg += fmt.Sprintf("  Distinguishing Mark - %s\n", v.(DistinguishingMark))
			continue
		case DrawbackProperty:
			msg += fmt.Sprintf("  Drawback - \n\t%s\n\t%s\n\t%s", v.(*Drawback).Name, v.(*Drawback).Description, v.(*Drawback).Effect)
			continue
		case JobProperty:
			msg += fmt.Sprintf("  Job - \n\t%s\n\t%s\n\t%s\n\t%d\n\t%s", v.(*Job).Name, v.(*Job).Description, v.(*Job).Archetype.Name, v.(*Job).Tier, v.(*Job).ExperienceCost)
			continue
		case StatsProperty:
			msg += fmt.Sprintf("  Stats - \n\tFighting: %d\n\tMuscle: %d\n\tSpeed: %d\n\tSavvy: %d\n\tSmarts: %d\n\tGrit: %d\n\tFlair: %d", v.(*Stats).Fighting, v.(*Stats).Muscle, v.(*Stats).Speed, v.(*Stats).Savvy, v.(*Stats).Smarts, v.(*Stats).Grit, v.(*Stats).Flair)
			continue
		case TattooProperty:
			msg += fmt.Sprintf("  Tattoo - \n\t%s\n", v.(*Tattoo).Description)
			continue
		case TeamProperty:
			msg += fmt.Sprintf("  Team - %s\n", v.(*Team).Name)
			continue
		default:
			msg += fmt.Sprintf("  %s: %s\n", k, v)
		}
	}
	return msg
}
