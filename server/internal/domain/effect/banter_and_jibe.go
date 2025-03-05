package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BanterandJibe struct {
	name        string
	description string
}

func NewBanterandJibe() *BanterandJibe {
	return &BanterandJibe{
		name:        "Banter and Jibe",
		description: "You may flip the results to succeed at Coordination Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Coordination Tests to perform acrobatics.",
	}
}

func (e *BanterandJibe) Name() string {
	return e.name
}

func (e *BanterandJibe) Description() string {
	return e.description
}

func (e *BanterandJibe) Applier() domain.Applier {
	return e.Apply
}

func (e *BanterandJibe) Apply(state domain.GameState) domain.GameState {
	// - You may flip the results to succeed at Coordination Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Coordination Tests to perform acrobatics.
	log.Println("applying Banter and Jibe")
	return state
}

var _ domain.Effect = &BanterandJibe{}
