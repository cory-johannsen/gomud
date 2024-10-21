package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MagnificentBastard struct {
  name string
  description string
}

func NewMagnificentBastard() *MagnificentBastard {
  return &MagnificentBastard{
    name: "Magnificent Bastard",
    description: "Select one foe in combat. That foe is left Defenseless to all Perilous Stunts you make until they are defeated. You may select another foe afterwards.",
  }
}

func (e *MagnificentBastard) Name() string {
  return e.name
}

func (e *MagnificentBastard) Description() string {
  return e.description
}

func (e *MagnificentBastard) Applier() domain.Applier {
  return e.Apply
}

func (e *MagnificentBastard) Apply(state domain.State) domain.State {
  // - Select one foe in combat. That foe is left Defenseless to all Perilous Stunts you make until they are defeated. You may select another foe afterwards.
  log.Println("applying Magnificent Bastard")
  return state
}

var _ domain.Effect = &MagnificentBastard{}
