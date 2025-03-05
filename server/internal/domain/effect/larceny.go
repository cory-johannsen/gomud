package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Larceny struct {
	name        string
	description string
}

func NewLarceny() *Larceny {
	return &Larceny{
		name:        "Larceny",
		description: "When fencing black market goods or procuring illegal information, you gain a +20 Base Chance to Bargain Tests.",
	}
}

func (e *Larceny) Name() string {
	return e.name
}

func (e *Larceny) Description() string {
	return e.description
}

func (e *Larceny) Applier() domain.Applier {
	return e.Apply
}

func (e *Larceny) Apply(state domain.GameState) domain.GameState {
	// - When fencing black market goods or procuring illegal information, you gain a +20 Base Chance to Bargain Tests.
	log.Println("applying Larceny")
	return state
}

var _ domain.Effect = &Larceny{}
