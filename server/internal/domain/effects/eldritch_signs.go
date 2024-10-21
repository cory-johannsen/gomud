package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EldritchSigns struct {
  Name string
  Description string
}

func (e *EldritchSigns) Apply(state domain.State) domain.State {
  // - You automatically gain all of the following Focuses in the Incantation Skill whenever you enter this Job - Astrology, Card Reading, Dream Interpretation, Hypnotize, Palmistry, Scapulimancy and Scrying. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Eldritch Signs")
  return state
}
