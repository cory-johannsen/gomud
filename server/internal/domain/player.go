package domain

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

const PropertyNotFound = "property not found"

const (
	BackgroundProperty = "background"
	TeamProperty       = "team"
	StatsProperty      = "stats"
)

type Player struct {
	Id       int
	Name     string
	password string
	Data     map[string]interface{}
}

func NewPlayer(id *int, name string, password string) *Player {
	p := &Player{
		Name:     name,
		password: password,
		Data:     make(map[string]interface{}),
	}
	if id != nil {
		p.Id = *id
	}
	return p
}

func (p *Player) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (p *Player) GetProperty(key string) (interface{}, error) {
	result, ok := p.Data[key]
	if !ok {
		return nil, errors.New(PropertyNotFound)
	}
	return result, nil
}

func (p *Player) String() string {
	msg := fmt.Sprintf("Name: %s\n", p.Name)
	for k, v := range p.Data {
		msg += fmt.Sprintf("  %s: %s\n", k, v)
	}
	return msg
}
