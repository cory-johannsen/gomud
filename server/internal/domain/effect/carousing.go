package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Carousing struct {
  name string
  description string
}

func NewCarousing() *Carousing {
  return &Carousing{
    name: "Carousing",
    description: "When you are Intoxicated, will you be a friendly drunk or a mean drunk? When you’re a friendly drunk, gain  a +10 Base Chance to Charm Tests while Intoxicated. When  you’re a mean drunk, gain a +10 Base Chance to Intimidate  Tests while Intoxicated. You can make this choice every time you become Intoxicated.",
  }
}

func (e *Carousing) Name() string {
  return e.name
}

func (e *Carousing) Description() string {
  return e.description
}

func (e *Carousing) Applier() domain.Applier {
  return e.Apply
}

func (e *Carousing) Apply(state domain.State) domain.State {
  // - When you are Intoxicated, will you be a friendly drunk or a mean drunk? When you’re a friendly drunk, gain  a +10 Base Chance to Charm Tests while Intoxicated. When  you’re a mean drunk, gain a +10 Base Chance to Intimidate  Tests while Intoxicated. You can make this choice every time you become Intoxicated.
  log.Println("applying Carousing")
  return state
}

var _ domain.Effect = &Carousing{}
