package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Azimuth struct {
  name string
  description string
}

func NewAzimuth() *Azimuth {
  return &Azimuth{
    name: "Azimuth",
    description: "When you can see the stars, you gain a +20 Base Chance to Navigation Tests. In addition, you always know true north above ground, even during inclement weather.",
  }
}

func (e *Azimuth) Name() string {
  return e.name
}

func (e *Azimuth) Description() string {
  return e.description
}

func (e *Azimuth) Applier() domain.Applier {
  return e.Apply
}

func (e *Azimuth) Apply(state domain.State) domain.State {
  // - When you can see the stars, you gain a +20 Base Chance to Navigation Tests. In addition, you always know true north above ground, even during inclement weather.
  log.Println("applying Azimuth")
  return state
}

var _ domain.Effect = &Azimuth{}
