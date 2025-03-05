package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Nighteyes struct {
	name        string
	description string
}

func NewNighteyes() *Nighteyes {
	return &Nighteyes{
		name:        "Nighteyes",
		description: "You can see completely in the dark above ground as if it were full daylight, provided there is starlight or moonlight in the sky.,",
	}
}

func (e *Nighteyes) Name() string {
	return e.name
}

func (e *Nighteyes) Description() string {
	return e.description
}

func (e *Nighteyes) Applier() domain.Applier {
	return e.Apply
}

func (e *Nighteyes) Apply(state domain.GameState) domain.GameState {
	// - You can see completely in the dark above ground as if it were full daylight, provided there is starlight or moonlight in the sky.,
	log.Println("applying Nighteyes")
	return state
}

var _ domain.Effect = &Nighteyes{}
