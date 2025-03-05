package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Gutplate struct {
	name        string
	description string
}

func NewGutplate() *Gutplate {
	return &Gutplate{
		name:        "Gut-plate",
		description: "Whenever unarmored, add +2 to your Damage Threshold as if you were wearing a suit of leather armor. You also gain the benefits of the Natural Quality in these cases.,",
	}
}

func (e *Gutplate) Name() string {
	return e.name
}

func (e *Gutplate) Description() string {
	return e.description
}

func (e *Gutplate) Applier() domain.Applier {
	return e.Apply
}

func (e *Gutplate) Apply(state domain.GameState) domain.GameState {
	// - Whenever unarmored, add +2 to your Damage Threshold as if you were wearing a suit of leather armor. You also gain the benefits of the Natural Quality in these cases.,
	log.Println("applying Gut-plate")
	return state
}

var _ domain.Effect = &Gutplate{}
