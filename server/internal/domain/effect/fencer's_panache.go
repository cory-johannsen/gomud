package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type FencersPanache struct {
  name string
  description string
}

func NewFencersPanache() *FencersPanache {
  return &FencersPanache{
    name: "Fencer's Panache",
    description: "When you Take Aim and then make a successful Melee Attack, you force a foe to Resist a Disarm. You must attack with weapons possessing the Fast or Finesse Quality to utilize this Talent.",
  }
}

func (e *FencersPanache) Name() string {
  return e.name
}

func (e *FencersPanache) Description() string {
  return e.description
}

func (e *FencersPanache) Applier() domain.Applier {
  return e.Apply
}

func (e *FencersPanache) Apply(state domain.State) domain.State {
  // - When you Take Aim and then make a successful Melee Attack, you force a foe to Resist a Disarm. You must attack with weapons possessing the Fast or Finesse Quality to utilize this Talent.
  log.Println("applying Fencer's Panache")
  return state
}

var _ domain.Effect = &FencersPanache{}
