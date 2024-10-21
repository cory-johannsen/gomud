package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FeastorFamine struct {
  Name string
  Description string
}

func (e *FeastorFamine) Apply(state domain.State) domain.State {
  // - You can attempt a Survival or Tradecraft Test to cook for others. If successful, a number of allies equal to your [FB] move one step up the Damage Condition Track positively, and recover their Peril Condition Track to Unhindered, whether they’re resting in a dangerous place or not. If your roll is a Critical Success, you affect a number of allies equal to three times your [FB] instead. A Character may only gain the benefit of this meal once per day, which takes an hour of time to prepare.
  log.Println("applying Feast or Famine")
  return state
}
