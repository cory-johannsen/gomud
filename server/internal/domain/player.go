package domain

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Property interface {
	Value() interface{}
	String() string
}

const PropertyNotFound = "property not found"

const (
	AgeProperty                = "age"
	AlignmentProperty          = "alignment"
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

type BaseProperty struct {
	Val interface{}
}

func (p *BaseProperty) Value() interface{} {
	return p.Val
}
func (p *BaseProperty) String() string {
	return fmt.Sprintf("%v", p.Val)
}

var _ Property = &BaseProperty{}

type Character struct {
	Name string
	Data map[string]Property
}

type Player struct {
	Character
	Id       *int
	Password string
}

func NewPlayer(id *int, name string, password string, data map[string]Property) *Player {
	p := &Player{
		Character: Character{
			Name: name,
			Data: data,
		},
		Id:       id,
		Password: password,
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
	// Enforce the ordering of the character properties
	properties := []string{
		StatsProperty,
		AlignmentProperty,
		BirthSeasonProperty,
		AgeProperty,
		BackgroundProperty,
		DistinguishingMarkProperty,
		TattooProperty,
		DrawbackProperty,
		ArchetypeProperty,
		JobProperty,
		TeamProperty}
	for _, k := range properties {
		v, ok := p.Data[k]
		if !ok {
			log.Printf("property %s not found", k)
			continue
		}
		switch k {
		case AgeProperty:
			msg += fmt.Sprintf("  Age - %d\n", v.(*BaseProperty).Val.(int))
			continue
		case AlignmentProperty:
			msg += fmt.Sprintf("  Alignment - %s/%s\n\tOrder: %s (rank: %d)\n\tChaos: %s (rank: %d)\n\tCorruption: %d\n", v.(*Alignment).Order.Name, v.(*Alignment).Chaos.Name, v.(*Alignment).Order.Name, v.(*Alignment).Order.Rank, v.(*Alignment).Chaos.Name, v.(*Alignment).Chaos.Rank, v.(*Alignment).Corruption)
			continue
		case ArchetypeProperty:
			msg += fmt.Sprintf("  Archetype - %s\n", v.(*Archetype).Name)
			for _, trait := range v.(*Archetype).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name, effect.Description)
				}
			}
			continue
		case BackgroundProperty:
			msg += fmt.Sprintf("  Background - \n\t%s\n\t%s\n", v.(*Background).Name, v.(*Background).Description)
			for _, trait := range v.(*Background).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name, effect.Description)
				}
			}
			continue
		case BirthSeasonProperty:
			msg += fmt.Sprintf("  Birth Season - %s\n", v.(Season))
			continue
		case DistinguishingMarkProperty:
			msg += "  Distinguishing Marks: \n"
			for _, mark := range v.(DistinguishingMarks) {
				msg += fmt.Sprintf("\t%s\n", mark)
			}
			continue
		case DrawbackProperty:
			msg += fmt.Sprintf("  Drawback - \n\t%s\n\tDescription: %s\n\tEffect: \n\t\t%s\n\t\tdesc\n\t\teffect\n", v.(*Drawback).Name, v.(*Drawback).Description, v.(*Drawback).Effect)

			continue
		case JobProperty:
			msg += fmt.Sprintf("  Job - \n\t%s\n\tDescription: %s\n\tArchetype: %s\n\tTier: %s\n\tExperience Cost: %d\n", v.(*Job).Name, v.(*Job).Description, v.(*Job).Archetype.Name, v.(*Job).Tier, v.(*Job).ExperienceCost)
			for _, trait := range v.(*Job).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name, effect.Description)
				}
			}
			continue
		case StatsProperty:
			msg += fmt.Sprintf("  Stats - \n\tFighting: %d\n\tMuscle: %d\n\tSpeed: %d\n\tSavvy: %d\n\tSmarts: %d\n\tGrit: %d\n\tFlair: %d\n", v.(*Stats).Fighting, v.(*Stats).Muscle, v.(*Stats).Speed, v.(*Stats).Savvy, v.(*Stats).Smarts, v.(*Stats).Grit, v.(*Stats).Flair)
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
