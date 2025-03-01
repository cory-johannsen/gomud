package domain

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type State interface {
	Player() *Player
	Sensor(string) any
	Property(string) Property
}

type StateProvider func() State

type StateConstructor func(player *Player, sensors htn.Sensors) State
