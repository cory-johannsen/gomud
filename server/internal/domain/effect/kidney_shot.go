package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type KidneyShot struct {
  name string
  description string
}

func NewKidneyShot() *KidneyShot {
  return &KidneyShot{
    name: "Kidney Shot",
    description: "You penalize your foe’s ability to Resist a Stunning Blow by a -10 Base Chance to their Skill Test. In addition, when you use a Stunning Blow, your foe loses 2 Action Points, instead of 1 Action Point.",
  }
}

func (e *KidneyShot) Name() string {
  return e.name
}

func (e *KidneyShot) Description() string {
  return e.description
}

func (e *KidneyShot) Applier() domain.Applier {
  return e.Apply
}

func (e *KidneyShot) Apply(state domain.State) domain.State {
  // - You penalize your foe’s ability to Resist a Stunning Blow by a -10 Base Chance to their Skill Test. In addition, when you use a Stunning Blow, your foe loses 2 Action Points, instead of 1 Action Point.
  log.Println("applying Kidney Shot")
  return state
}

var _ domain.Effect = &KidneyShot{}
