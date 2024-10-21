package domain

type State interface {
	Player() *Player
	Property(string) Property
}

type StateProvider func() State

type StateConstructor func(player *Player) State
