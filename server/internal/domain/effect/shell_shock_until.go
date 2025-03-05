package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ShellShockUntil struct {
	name        string
	description string
}

func NewShellShockUntil() *ShellShockUntil {
	return &ShellShockUntil{
		name:        "Shell Shock Until",
		description: "fully Recuperated, you cannot add Fury Dice to Damage.",
	}
}

func (e *ShellShockUntil) Name() string {
	return e.name
}

func (e *ShellShockUntil) Description() string {
	return e.description
}

func (e *ShellShockUntil) Applier() domain.Applier {
	return e.Apply
}

func (e *ShellShockUntil) Apply(state domain.GameState) domain.GameState {
	// - fully Recuperated, you cannot add Fury Dice to Damage.
	log.Println("applying Shell Shock Until")
	return state
}

var _ domain.Effect = &ShellShockUntil{}
