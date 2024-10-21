package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ConfidenceTrick struct {
  Name string
  Description string
}

func (e *ConfidenceTrick) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Guile Tests. When you succeed, it is always considered a Critical Success. Furthermore, you can influence a number of people with the Guile Skill equal to your [FB] times three – this includes using Dirty Tricks during combat.
  log.Println("applying Confidence Trick")
  return state
}
