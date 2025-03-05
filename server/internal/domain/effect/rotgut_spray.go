package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RotgutSpray struct {
	name        string
	description string
}

func NewRotgutSpray() *RotgutSpray {
	return &RotgutSpray{
		name:        "Rotgut Spray",
		description: "After consuming an entire bottle of rotgut, you can belch it up as flaming breath. Treat this breath as a bottle bomb. However, you must immediately make a Toughness Test or become Intoxicated.,",
	}
}

func (e *RotgutSpray) Name() string {
	return e.name
}

func (e *RotgutSpray) Description() string {
	return e.description
}

func (e *RotgutSpray) Applier() domain.Applier {
	return e.Apply
}

func (e *RotgutSpray) Apply(state domain.GameState) domain.GameState {
	// - After consuming an entire bottle of rotgut, you can belch it up as flaming breath. Treat this breath as a bottle bomb. However, you must immediately make a Toughness Test or become Intoxicated.,
	log.Println("applying Rotgut Spray")
	return state
}

var _ domain.Effect = &RotgutSpray{}
