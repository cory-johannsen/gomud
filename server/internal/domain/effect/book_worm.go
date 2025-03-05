package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BookWorm struct {
	name        string
	description string
}

func NewBookWorm() *BookWorm {
	return &BookWorm{
		name:        "Book Worm",
		description: "Multiply your [RB] by three to determine how many Focuses you may possess.",
	}
}

func (e *BookWorm) Name() string {
	return e.name
}

func (e *BookWorm) Description() string {
	return e.description
}

func (e *BookWorm) Applier() domain.Applier {
	return e.Apply
}

func (e *BookWorm) Apply(state domain.GameState) domain.GameState {
	// - Multiply your [RB] by three to determine how many Focuses you may possess.
	log.Println("applying Book Worm")
	return state
}

var _ domain.Effect = &BookWorm{}
