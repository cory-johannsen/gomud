package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FrighteningBellow struct {
	name        string
	description string
}

func NewFrighteningBellow() *FrighteningBellow {
	return &FrighteningBellow{
		name:        "Frightening Bellow",
		description: "Whenever you succeed at an Intimidate Test, the foes you intimidated also suffer from Stress. If it is a Critical Success, they suffer from Fear instead.,",
	}
}

func (e *FrighteningBellow) Name() string {
	return e.name
}

func (e *FrighteningBellow) Description() string {
	return e.description
}

func (e *FrighteningBellow) Applier() domain.Applier {
	return e.Apply
}

func (e *FrighteningBellow) Apply(state domain.GameState) domain.GameState {
	// - Whenever you succeed at an Intimidate Test, the foes you intimidated also suffer from Stress. If it is a Critical Success, they suffer from Fear instead.,
	log.Println("applying Frightening Bellow")
	return state
}

var _ domain.Effect = &FrighteningBellow{}
