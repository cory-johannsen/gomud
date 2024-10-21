package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type EldritchSigns struct {
  name string
  description string
}

func NewEldritchSigns() *EldritchSigns {
  return &EldritchSigns{
    name: "Eldritch Signs",
    description: "You automatically gain all of the following Focuses in the Incantation Skill whenever you enter this Job - Astrology, Card Reading, Dream Interpretation, Hypnotize, Palmistry, Scapulimancy and Scrying. You also understand how to use the Ritual of Magick Circle.",
  }
}

func (e *EldritchSigns) Name() string {
  return e.name
}

func (e *EldritchSigns) Description() string {
  return e.description
}

func (e *EldritchSigns) Applier() domain.Applier {
  return e.Apply
}

func (e *EldritchSigns) Apply(state domain.State) domain.State {
  // - You automatically gain all of the following Focuses in the Incantation Skill whenever you enter this Job - Astrology, Card Reading, Dream Interpretation, Hypnotize, Palmistry, Scapulimancy and Scrying. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Eldritch Signs")
  return state
}

var _ domain.Effect = &EldritchSigns{}
