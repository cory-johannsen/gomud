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
	BackgroundProperty = "background"
	TeamProperty       = "team"
	StatsProperty      = "stats"
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
		if k == StatsProperty {
			msg += fmt.Sprintf("Stats - \n%s", v.String())
			continue
		}
		if k == BackgroundProperty {
			msg += fmt.Sprintf("Background - \n%s", v.String())
			continue
		}
		if k == TeamProperty {
			msg += fmt.Sprintf("Team - %s\n", v.String())
			continue
		}
		msg += fmt.Sprintf("  %s: %s\n", k, v)
	}
	return msg
}
