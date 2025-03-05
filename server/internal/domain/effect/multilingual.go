package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Multilingual struct {
	name        string
	description string
}

func NewMultilingual() *Multilingual {
	return &Multilingual{
		name:        "Multilingual",
		description: "You can communicate simple thoughts through words and hand signals with other Backgrounds whom you do not share the same language with.",
	}
}

func (e *Multilingual) Name() string {
	return e.name
}

func (e *Multilingual) Description() string {
	return e.description
}

func (e *Multilingual) Applier() domain.Applier {
	return e.Apply
}

func (e *Multilingual) Apply(state domain.GameState) domain.GameState {
	// - You can communicate simple thoughts through words and hand signals with other Backgrounds whom you do not share the same language with.
	log.Println("applying Multilingual")
	return state
}

var _ domain.Effect = &Multilingual{}
