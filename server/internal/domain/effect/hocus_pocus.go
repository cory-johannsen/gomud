package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type HocusPocus struct {
	name        string
	description string
}

func NewHocusPocus() *HocusPocus {
	return &HocusPocus{
		name:        "Hocus Pocus",
		description: "Whenever a foe fails to Resist your Magicks, they Critically Fail instead.,",
	}
}

func (e *HocusPocus) Name() string {
	return e.name
}

func (e *HocusPocus) Description() string {
	return e.description
}

func (e *HocusPocus) Applier() domain.Applier {
	return e.Apply
}

func (e *HocusPocus) Apply(state domain.GameState) domain.GameState {
	// - Whenever a foe fails to Resist your Magicks, they Critically Fail instead.,
	log.Println("applying Hocus Pocus")
	return state
}

var _ domain.Effect = &HocusPocus{}
