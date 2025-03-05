package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Finesse struct {
	name        string
	description string
}

func NewFinesse() *Finesse {
	return &Finesse{
		name:        "Finesse",
		description: "Weapons of this Quality always reference [AB] whenever dealing Damage, instead of [BB].",
	}
}

func (e *Finesse) Name() string {
	return e.name
}

func (e *Finesse) Description() string {
	return e.description
}

func (e *Finesse) Applier() domain.Applier {
	return e.Apply
}

func (e *Finesse) Apply(state domain.GameState) domain.GameState {
	// - Weapons of this Quality always reference [AB] whenever dealing Damage, instead of [BB].
	log.Println("applying Finesse")
	return state
}

var _ domain.Effect = &Finesse{}
