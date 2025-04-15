package domain

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type GameState interface {
	Player() *Player
	Sensor(string) any
	Property(string) Property
	Domain() *htn.Domain
}

type StateProvider func() GameState

type StateConstructor func(player *Player, sensors htn.Sensors) GameState
