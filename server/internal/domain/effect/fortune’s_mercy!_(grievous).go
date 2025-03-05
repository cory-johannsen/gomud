package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FortunesMercyGrievous struct {
	name        string
	description string
}

func NewFortunesMercyGrievous() *FortunesMercyGrievous {
	return &FortunesMercyGrievous{
		name:        "Fortune’s Mercy! (Grievous)",
		description: "Ignore Injury, keep your Fate Point and continue fighting onwards!",
	}
}

func (e *FortunesMercyGrievous) Name() string {
	return e.name
}

func (e *FortunesMercyGrievous) Description() string {
	return e.description
}

func (e *FortunesMercyGrievous) Applier() domain.Applier {
	return e.Apply
}

func (e *FortunesMercyGrievous) Apply(state domain.GameState) domain.GameState {
	// - Ignore Injury, keep your Fate Point and continue fighting onwards!
	log.Println("applying Fortune’s Mercy! (Grievous)")
	return state
}

var _ domain.Effect = &FortunesMercyGrievous{}
