package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Ambidexterity struct {
  name string
  description string
}

func NewAmbidexterity() *Ambidexterity {
  return &Ambidexterity{
    name: "Ambidexterity",
    description: "You never suffer penalties when using tools or weapons in either hand. If you ever suffer an Injury where you cannot use your primary hand, you do not suffer penalties to use your off-hand.",
  }
}

func (e *Ambidexterity) Name() string {
  return e.name
}

func (e *Ambidexterity) Description() string {
  return e.description
}

func (e *Ambidexterity) Applier() domain.Applier {
  return e.Apply
}

func (e *Ambidexterity) Apply(state domain.State) domain.State {
  // - You never suffer penalties when using tools or weapons in either hand. If you ever suffer an Injury where you cannot use your primary hand, you do not suffer penalties to use your off-hand.
  log.Println("applying Ambidexterity")
  return state
}

var _ domain.Effect = &Ambidexterity{}
