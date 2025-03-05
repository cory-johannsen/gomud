package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type EnduringMortality struct {
	name        string
	description string
}

func NewEnduringMortality() *EnduringMortality {
	return &EnduringMortality{
		name:        "Enduring Mortality",
		description: "You never suffer the debilitating effects of any Disease, aging or sickness – even by the hand of Magick.,",
	}
}

func (e *EnduringMortality) Name() string {
	return e.name
}

func (e *EnduringMortality) Description() string {
	return e.description
}

func (e *EnduringMortality) Applier() domain.Applier {
	return e.Apply
}

func (e *EnduringMortality) Apply(state domain.GameState) domain.GameState {
	// - You never suffer the debilitating effects of any Disease, aging or sickness – even by the hand of Magick.,
	log.Println("applying Enduring Mortality")
	return state
}

var _ domain.Effect = &EnduringMortality{}
