package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ButcheredLeg struct {
  name string
  description string
}

func NewButcheredLeg() *ButcheredLeg {
  return &ButcheredLeg{
    name: "Butchered Leg",
    description: "Until fully Recuperated, you cannot move as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Butchered Leg has undergone a failed surgery, you gain the Veteran’s Leg Drawback. If you already have this Drawback, you permanently lose 9% Agility.",
  }
}

func (e *ButcheredLeg) Name() string {
  return e.name
}

func (e *ButcheredLeg) Description() string {
  return e.description
}

func (e *ButcheredLeg) Applier() domain.Applier {
  return e.Apply
}

func (e *ButcheredLeg) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot move as you’re in pain. You must undergo a successful surgery or suffer the consequences. Once a Butchered Leg has undergone a failed surgery, you gain the Veteran’s Leg Drawback. If you already have this Drawback, you permanently lose 9% Agility.
  log.Println("applying Butchered Leg")
  return state
}

var _ domain.Effect = &ButcheredLeg{}
