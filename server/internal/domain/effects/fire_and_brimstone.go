package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FireandBrimstone struct {
  Name string
  Description string
}

func (e *FireandBrimstone) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Leadership Tests. When you succeed, it is always considered a Critical Success. Furthermore, you always influence a number of people with the Leadership Skill equal to three times your [FB] – this includes use of Inspiring Words during combat. You also understand how to use the Ritual of Blessed Sacrament.
  log.Println("applying Fire and Brimstone")
  return state
}
