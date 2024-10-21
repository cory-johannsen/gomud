package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Eureka struct {
  Name string
  Description string
}

func (e *Eureka) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Tradecraft Tests. When you succeed, it is always considered a Critical Success. In addition, you always succeed at Skill Tests to locate resources needed for construction.
  log.Println("applying Eureka!")
  return state
}
