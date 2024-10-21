package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WindsofChange struct {
  Name string
  Description string
}

func (e *WindsofChange) Apply(state domain.State) domain.State {
  // - When you would potentially invoke a Chaos Manifestation, you must roll two or more face ‘6s’ on the Chaos Dice to invoke it. Otherwise, ignore the results. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Winds of Change")
  return state
}
