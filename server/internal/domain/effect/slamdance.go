package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Slamdance struct {
	name        string
	description string
}

func NewSlamdance() *Slamdance {
	return &Slamdance{
		name:        "Slamdance",
		description: "Whenever you make a successful Stunning Blow, you force a foe to Resist a Takedown as well. You must attack with weapons possessing the Pummeling or Weak Quality to utilize this Talent.,",
	}
}

func (e *Slamdance) Name() string {
	return e.name
}

func (e *Slamdance) Description() string {
	return e.description
}

func (e *Slamdance) Applier() domain.Applier {
	return e.Apply
}

func (e *Slamdance) Apply(state domain.GameState) domain.GameState {
	// - Whenever you make a successful Stunning Blow, you force a foe to Resist a Takedown as well. You must attack with weapons possessing the Pummeling or Weak Quality to utilize this Talent.,
	log.Println("applying Slamdance")
	return state
}

var _ domain.Effect = &Slamdance{}
