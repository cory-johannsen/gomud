package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type PuncturedLung struct {
	name        string
	description string
}

func NewPuncturedLung() *PuncturedLung {
	return &PuncturedLung{
		name:        "Punctured Lung",
		description: "Until fully Recuperated, you remain unconscious. You must undergo a successful surgery or suffer the consequences. Once a Punctured Lung has undergone a failed surgery, you permanently gain the Weak Lungs Drawback. If you already have this Drawback, you permanently lose 9% Willpower.",
	}
}

func (e *PuncturedLung) Name() string {
	return e.name
}

func (e *PuncturedLung) Description() string {
	return e.description
}

func (e *PuncturedLung) Applier() domain.Applier {
	return e.Apply
}

func (e *PuncturedLung) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you remain unconscious. You must undergo a successful surgery or suffer the consequences. Once a Punctured Lung has undergone a failed surgery, you permanently gain the Weak Lungs Drawback. If you already have this Drawback, you permanently lose 9% Willpower.
	log.Println("applying Punctured Lung")
	return state
}

var _ domain.Effect = &PuncturedLung{}
