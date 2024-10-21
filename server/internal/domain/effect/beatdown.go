package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Beatdown struct {
  name string
  description string
}

func NewBeatdown() *Beatdown {
  return &Beatdown{
    name: "Beatdown",
    description: "When you Take Aim and then make a successful Melee Attack, you force a foe to Resist a Takedown. You must attack with Brawling or Crushing types of weapons to utilize this Talent.",
  }
}

func (e *Beatdown) Name() string {
  return e.name
}

func (e *Beatdown) Description() string {
  return e.description
}

func (e *Beatdown) Applier() domain.Applier {
  return e.Apply
}

func (e *Beatdown) Apply(state domain.State) domain.State {
  // - When you Take Aim and then make a successful Melee Attack, you force a foe to Resist a Takedown. You must attack with Brawling or Crushing types of weapons to utilize this Talent.
  log.Println("applying Beatdown")
  return state
}

var _ domain.Effect = &Beatdown{}
