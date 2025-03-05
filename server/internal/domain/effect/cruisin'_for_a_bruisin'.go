package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type CruisinforaBruisin struct {
	name        string
	description string
}

func NewCruisinforaBruisin() *CruisinforaBruisin {
	return &CruisinforaBruisin{
		name:        "Cruisin' for a Bruisin'",
		description: "You may reference either [BB] or [BB] with any melee weapons you wield, including those which possess the Throwing Quality.,",
	}
}

func (e *CruisinforaBruisin) Name() string {
	return e.name
}

func (e *CruisinforaBruisin) Description() string {
	return e.description
}

func (e *CruisinforaBruisin) Applier() domain.Applier {
	return e.Apply
}

func (e *CruisinforaBruisin) Apply(state domain.GameState) domain.GameState {
	// - You may reference either [BB] or [BB] with any melee weapons you wield, including those which possess the Throwing Quality.,
	log.Println("applying Cruisin' for a Bruisin'")
	return state
}

var _ domain.Effect = &CruisinforaBruisin{}
