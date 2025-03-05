package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Eureka struct {
	name        string
	description string
}

func NewEureka() *Eureka {
	return &Eureka{
		name:        "Eureka!",
		description: "You may flip the results to succeed at Tradecraft Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Skill Tests to locate resources needed for construction.",
	}
}

func (e *Eureka) Name() string {
	return e.name
}

func (e *Eureka) Description() string {
	return e.description
}

func (e *Eureka) Applier() domain.Applier {
	return e.Apply
}

func (e *Eureka) Apply(state domain.GameState) domain.GameState {
	// - You may flip the results to succeed at Tradecraft Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Skill Tests to locate resources needed for construction.
	log.Println("applying Eureka!")
	return state
}

var _ domain.Effect = &Eureka{}
