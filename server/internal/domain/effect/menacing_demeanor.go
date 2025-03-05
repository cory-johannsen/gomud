package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MenacingDemeanor struct {
	name        string
	description string
}

func NewMenacingDemeanor() *MenacingDemeanor {
	return &MenacingDemeanor{
		name:        "Menacing Demeanor",
		description: "When you succeed at an Intimidate Test, you inflict 1D10+1 mental Peril.",
	}
}

func (e *MenacingDemeanor) Name() string {
	return e.name
}

func (e *MenacingDemeanor) Description() string {
	return e.description
}

func (e *MenacingDemeanor) Applier() domain.Applier {
	return e.Apply
}

func (e *MenacingDemeanor) Apply(state domain.GameState) domain.GameState {
	// - When you succeed at an Intimidate Test, you inflict 1D10+1 mental Peril.
	log.Println("applying Menacing Demeanor")
	return state
}

var _ domain.Effect = &MenacingDemeanor{}
