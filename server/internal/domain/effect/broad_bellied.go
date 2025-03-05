package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BroadBellied struct {
	name        string
	description string
}

func NewBroadBellied() *BroadBellied {
	return &BroadBellied{
		name:        "Broad Bellied",
		description: "You entirely ignore the effects of Overage, unless the value is in excess of two times your Initiative or Movement.,",
	}
}

func (e *BroadBellied) Name() string {
	return e.name
}

func (e *BroadBellied) Description() string {
	return e.description
}

func (e *BroadBellied) Applier() domain.Applier {
	return e.Apply
}

func (e *BroadBellied) Apply(state domain.GameState) domain.GameState {
	// - You entirely ignore the effects of Overage, unless the value is in excess of two times your Initiative or Movement.,
	log.Println("applying Broad Bellied")
	return state
}

var _ domain.Effect = &BroadBellied{}
