package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PhlegmaticTemperament struct {
  name string
  description string
}

func NewPhlegmaticTemperament() *PhlegmaticTemperament {
  return &PhlegmaticTemperament{
    name: "Phlegmatic Temperament",
    description: "Whenever you are suffering from Stress, Fear or Terror, your Fury Dice do not explode. This lasts until you get a good night's rest.",
  }
}

func (e *PhlegmaticTemperament) Name() string {
  return e.name
}

func (e *PhlegmaticTemperament) Description() string {
  return e.description
}

func (e *PhlegmaticTemperament) Applier() domain.Applier {
  return e.Apply
}

func (e *PhlegmaticTemperament) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Phlegmatic Temperament")
  return state
}

var _ domain.Effect = &PhlegmaticTemperament{}
