package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Craven struct {
  name string
  description string
}

func NewCraven() *Craven {
  return &Craven{
    name: "Craven",
    description: "You may flip the results to succeed at Combat-based Skill Tests to Parry attacks made with melee weapons.,",
  }
}

func (e *Craven) Name() string {
  return e.name
}

func (e *Craven) Description() string {
  return e.description
}

func (e *Craven) Applier() domain.Applier {
  return e.Apply
}

func (e *Craven) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Combat-based Skill Tests to Parry attacks made with melee weapons.,
  log.Println("applying Craven")
  return state
}

var _ domain.Effect = &Craven{}
