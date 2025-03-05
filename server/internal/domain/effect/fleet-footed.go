package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type FleetFooted struct {
	name        string
	description string
}

func NewFleetFooted() *FleetFooted {
	return &FleetFooted{
		name:        "Fleet-Footed",
		description: "Whenever you Charge or Run, add +6 yards to Movement.,",
	}
}

func (e *FleetFooted) Name() string {
	return e.name
}

func (e *FleetFooted) Description() string {
	return e.description
}

func (e *FleetFooted) Applier() domain.Applier {
	return e.Apply
}

func (e *FleetFooted) Apply(state domain.GameState) domain.GameState {
	// - Whenever you Charge or Run, add +6 yards to Movement.,
	log.Println("applying Fleet-Footed")
	return state
}

var _ domain.Effect = &FleetFooted{}
