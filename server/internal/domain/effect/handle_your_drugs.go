package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HandleYourDrugs struct {
  name string
  description string
}

func NewHandleYourDrugs() *HandleYourDrugs {
  return &HandleYourDrugs{
    name: "Handle Your Drugs",
    description: "When you would potentially invoke a Chaos Manifestation, you must roll two or more face ‘6s’ on the Chaos Dice to invoke it. Otherwise, ignore the results. You also understand how to use the Ritual of Magick Circle.",
  }
}

func (e *HandleYourDrugs) Name() string {
  return e.name
}

func (e *HandleYourDrugs) Description() string {
  return e.description
}

func (e *HandleYourDrugs) Applier() domain.Applier {
  return e.Apply
}

func (e *HandleYourDrugs) Apply(state domain.State) domain.State {
  // - When you would potentially invoke a Chaos Manifestation, you must roll two or more face ‘6s’ on the Chaos Dice to invoke it. Otherwise, ignore the results. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Handle Your Drugs")
  return state
}

var _ domain.Effect = &HandleYourDrugs{}
