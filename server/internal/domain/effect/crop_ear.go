package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type CropEar struct {
	name        string
	description string
}

func NewCropEar() *CropEar {
	return &CropEar{
		name:        "Crop Ear",
		description: "You must flip the results to fail all Skill Tests that rely on hearing.",
	}
}

func (e *CropEar) Name() string {
	return e.name
}

func (e *CropEar) Description() string {
	return e.description
}

func (e *CropEar) Applier() domain.Applier {
	return e.Apply
}

func (e *CropEar) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Crop Ear")
	return state
}

var _ domain.Effect = &CropEar{}
