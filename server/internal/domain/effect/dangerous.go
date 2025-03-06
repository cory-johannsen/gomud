package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Dangerous struct {
	name        string
	description string
}

func NewDangerous() *Dangerous {
	return &Dangerous{
		name:        "Dangerous",
		description: "Clothing or armor of this Quality cannot adequately protect you from harm. Should you suffer an Injury during this time without wearing a suit of armor, you begin to Bleed.",
	}
}

func (e *Dangerous) Name() string {
	return e.name
}

func (e *Dangerous) Description() string {
	return e.description
}

func (e *Dangerous) Applier() domain.Applier {
	return e.Apply
}

func (e *Dangerous) Apply(state domain.GameState) domain.GameState {
	// - Clothing or armor of this Quality cannot adequately protect you from harm. Should you suffer an Injury during this time without wearing a suit of armor, you begin to Bleed.
	log.Println("applying Dangerous")
	return state
}

var _ domain.Effect = &Dangerous{}
