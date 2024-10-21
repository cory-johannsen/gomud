package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VeteransHand struct {
  name string
  description string
}

func NewVeteransHand() *VeteransHand {
  return &VeteransHand{
    name: "Veteran’s Hand",
    description: "You cannot hold weapons which are two-handed and must flip the results to fail any Skill Test requiring use of both hands.",
  }
}

func (e *VeteransHand) Name() string {
  return e.name
}

func (e *VeteransHand) Description() string {
  return e.description
}

func (e *VeteransHand) Applier() domain.Applier {
  return e.Apply
}

func (e *VeteransHand) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Veteran’s Hand")
  return state
}

var _ domain.Effect = &VeteransHand{}
