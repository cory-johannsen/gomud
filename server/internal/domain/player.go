package domain

import "errors"

const PropertyNotFound = "property not found"

type Player struct {
	Id       int
	Name     string
	password string
	Data     map[string]interface{}
}

func NewPlayer(id int, name string, password string) *Player {
	return &Player{
		Id:       id,
		Name:     name,
		password: password,
		Data:     make(map[string]interface{}),
	}
}

func (p *Player) ValidPassword(password string) bool {
	return p.password == password
}

func (p *Player) GetProperty(key string) (interface{}, error) {
	result, ok := p.Data[key]
	if !ok {
		return nil, errors.New(PropertyNotFound)
	}
	return result, nil
}
