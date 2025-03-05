package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type VitreousHemorrhage struct {
	name        string
	description string
}

func NewVitreousHemorrhage() *VitreousHemorrhage {
	return &VitreousHemorrhage{
		name:        "Vitreous Hemorrhage",
		description: "Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.",
	}
}

func (e *VitreousHemorrhage) Name() string {
	return e.name
}

func (e *VitreousHemorrhage) Description() string {
	return e.description
}

func (e *VitreousHemorrhage) Applier() domain.Applier {
	return e.Apply
}

func (e *VitreousHemorrhage) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Vitreous Hemorrhage")
	return state
}

var _ domain.Effect = &VitreousHemorrhage{}
