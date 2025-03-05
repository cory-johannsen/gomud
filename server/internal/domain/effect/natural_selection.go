package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type NaturalSelection struct {
	name        string
	description string
}

func NewNaturalSelection() *NaturalSelection {
	return &NaturalSelection{
		name:        "Natural Selection",
		description: "You may permanently change any one Primary Attribute to a 55%.,",
	}
}

func (e *NaturalSelection) Name() string {
	return e.name
}

func (e *NaturalSelection) Description() string {
	return e.description
}

func (e *NaturalSelection) Applier() domain.Applier {
	return e.Apply
}

func (e *NaturalSelection) Apply(state domain.GameState) domain.GameState {
	// - You may permanently change any one Primary Attribute to a 55%.,
	log.Println("applying Natural Selection")
	return state
}

var _ domain.Effect = &NaturalSelection{}
