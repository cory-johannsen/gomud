package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Volatile struct {
  name string
  description string
}

func NewVolatile() *Volatile {
  return &Volatile{
    name: "Volatile",
    description: "When you Critically Fail an Attack Action or Perilous Stunt when using weapons of this Quality, roll a 1D6 Chaos Die. On a result of 1 to 5, the weapon misfires, requiring an hour to clean and repair. On a result of face ‘6’, it explodes, dealing 2D10+2 Damage from fire to you and destroying the weapon. This Damage cannot be Dodged, Parried or Resisted.",
  }
}

func (e *Volatile) Name() string {
  return e.name
}

func (e *Volatile) Description() string {
  return e.description
}

func (e *Volatile) Applier() domain.Applier {
  return e.Apply
}

func (e *Volatile) Apply(state domain.State) domain.State {
  // - When you Critically Fail an Attack Action or Perilous Stunt when using weapons of this Quality, roll a 1D6 Chaos Die. On a result of 1 to 5, the weapon misfires, requiring an hour to clean and repair. On a result of face ‘6’, it explodes, dealing 2D10+2 Damage from fire to you and destroying the weapon. This Damage cannot be Dodged, Parried or Resisted.
  log.Println("applying Volatile")
  return state
}

var _ domain.Effect = &Volatile{}
