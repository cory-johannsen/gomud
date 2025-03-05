package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RuneMarkedGlory struct {
	name        string
	description string
}

func NewRuneMarkedGlory() *RuneMarkedGlory {
	return &RuneMarkedGlory{
		name:        "Rune-Marked Glory",
		description: "Whenever you suffer Damage or Peril from Magick, spend a Fortune Point to ignore it entirely. Note that this does not alleviate any additional effects the Magick spell may also cause.,",
	}
}

func (e *RuneMarkedGlory) Name() string {
	return e.name
}

func (e *RuneMarkedGlory) Description() string {
	return e.description
}

func (e *RuneMarkedGlory) Applier() domain.Applier {
	return e.Apply
}

func (e *RuneMarkedGlory) Apply(state domain.GameState) domain.GameState {
	// - Whenever you suffer Damage or Peril from Magick, spend a Fortune Point to ignore it entirely. Note that this does not alleviate any additional effects the Magick spell may also cause.,
	log.Println("applying Rune-Marked Glory")
	return state
}

var _ domain.Effect = &RuneMarkedGlory{}
