package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FortunesMercyModerate struct {
	name        string
	description string
}

func NewFortunesMercyModerate() *FortunesMercyModerate {
	return &FortunesMercyModerate{
		name:        "Fortune’s Mercy! (Moderate)",
		description: "Ignore Injury, keep your Fate Point and continue fighting onwards!",
	}
}

func (e *FortunesMercyModerate) Name() string {
	return e.name
}

func (e *FortunesMercyModerate) Description() string {
	return e.description
}

func (e *FortunesMercyModerate) Applier() domain.Applier {
	return e.Apply
}

func (e *FortunesMercyModerate) Apply(state domain.GameState) domain.GameState {
	// - Ignore Injury, keep your Fate Point and continue fighting onwards!
	log.Println("applying Fortune’s Mercy! (Moderate)")
	return state
}

var _ domain.Effect = &FortunesMercyModerate{}
