package domain

type Applier func(state GameState) GameState

type Effect interface {
	Name() string
	Description() string
	Applier() Applier
}
type Effects []Effect
