package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VeteransLeg struct {
  name string
  description string
}

func NewVeteransLeg() *VeteransLeg {
  return &VeteransLeg{
    name: "Veteran’s Leg",
    description: "You must reduce your Movement by 3 and cannot Run.",
  }
}

func (e *VeteransLeg) Name() string {
  return e.name
}

func (e *VeteransLeg) Description() string {
  return e.description
}

func (e *VeteransLeg) Applier() domain.Applier {
  return e.Apply
}

func (e *VeteransLeg) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Veteran’s Leg")
  return state
}

var _ domain.Effect = &VeteransLeg{}
