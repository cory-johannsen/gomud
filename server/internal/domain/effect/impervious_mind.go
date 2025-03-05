package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ImperviousMind struct {
	name        string
	description string
}

func NewImperviousMind() *ImperviousMind {
	return &ImperviousMind{
		name:        "Impervious Mind",
		description: "When you suffer mental Peril, reduce your Peril Condition Track by one less step negatively.",
	}
}

func (e *ImperviousMind) Name() string {
	return e.name
}

func (e *ImperviousMind) Description() string {
	return e.description
}

func (e *ImperviousMind) Applier() domain.Applier {
	return e.Apply
}

func (e *ImperviousMind) Apply(state domain.GameState) domain.GameState {
	// - When you suffer mental Peril, reduce your Peril Condition Track by one less step negatively.
	log.Println("applying Impervious Mind")
	return state
}

var _ domain.Effect = &ImperviousMind{}
