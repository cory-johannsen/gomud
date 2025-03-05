package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type AppallingMien struct {
	name        string
	description string
}

func NewAppallingMien() *AppallingMien {
	return &AppallingMien{
		name:        "Appalling Mien",
		description: "When you succeed at an Intimidate Test against one foe, they cannot attack you until they succeed at a Resolve Test. However, if you or your allies harm them in any way, they immediately shake off this effect.",
	}
}

func (e *AppallingMien) Name() string {
	return e.name
}

func (e *AppallingMien) Description() string {
	return e.description
}

func (e *AppallingMien) Applier() domain.Applier {
	return e.Apply
}

func (e *AppallingMien) Apply(state domain.GameState) domain.GameState {
	// - When you succeed at an Intimidate Test against one foe, they cannot attack you until they succeed at a Resolve Test. However, if you or your allies harm them in any way, they immediately shake off this effect.
	log.Println("applying Appalling Mien")
	return state
}

var _ domain.Effect = &AppallingMien{}
