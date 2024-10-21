package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TheFullMonty struct {
  Name string
  Description string
}

func (e *TheFullMonty) Apply(state domain.State) domain.State {
  // - All Skill Ranks you acquire in the Charm Skill modify your Base Chance by +20, instead of +10. In addition, when you perform acts of carnal knowledge with another, you both ‘get lucky’, gaining 1 Fortune Point each to personally use. This benefit cannot be given or gained more than once per day.
  log.Println("applying The Full Monty")
  return state
}
