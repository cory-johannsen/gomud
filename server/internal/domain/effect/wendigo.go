package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Wendigo struct {
	name        string
	description string
}

func NewWendigo() *Wendigo {
	return &Wendigo{
		name:        "Wendigo",
		description: "You are immune from the effects of Frostbite, and automatically succeed at Toughness Tests to withstand cold weather.,",
	}
}

func (e *Wendigo) Name() string {
	return e.name
}

func (e *Wendigo) Description() string {
	return e.description
}

func (e *Wendigo) Applier() domain.Applier {
	return e.Apply
}

func (e *Wendigo) Apply(state domain.GameState) domain.GameState {
	// - You are immune from the effects of Frostbite, and automatically succeed at Toughness Tests to withstand cold weather.,
	log.Println("applying Wendigo")
	return state
}

var _ domain.Effect = &Wendigo{}
