package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Cavesight struct {
	name        string
	description string
}

func NewCavesight() *Cavesight {
	return &Cavesight{
		name:        "Cavesight",
		description: "You can see completely in the dark below ground as if it were daylight, providing you are able to hear.,",
	}
}

func (e *Cavesight) Name() string {
	return e.name
}

func (e *Cavesight) Description() string {
	return e.description
}

func (e *Cavesight) Applier() domain.Applier {
	return e.Apply
}

func (e *Cavesight) Apply(state domain.GameState) domain.GameState {
	// - You can see completely in the dark below ground as if it were daylight, providing you are able to hear.,
	log.Println("applying Cavesight")
	return state
}

var _ domain.Effect = &Cavesight{}
