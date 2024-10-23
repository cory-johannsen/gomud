package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HungerPangs struct {
  name string
  description string
}

func NewHungerPangs() *HungerPangs {
  return &HungerPangs{
    name: "Hunger Pangs",
    description: "Whenever you successfully Resist the effects of Stress, Fear, Terror or other effects that force you to make a Resolve Test, you immediately gain the benefits of one dose of mandrake root. If this is a Critical Success, you gain the benefits of three doses instead. When under Hunger Pang’s effects, you do not suffer the drawbacks normally associated with taking mandrake root.,",
  }
}

func (e *HungerPangs) Name() string {
  return e.name
}

func (e *HungerPangs) Description() string {
  return e.description
}

func (e *HungerPangs) Applier() domain.Applier {
  return e.Apply
}

func (e *HungerPangs) Apply(state domain.State) domain.State {
  // - Whenever you successfully Resist the effects of Stress, Fear, Terror or other effects that force you to make a Resolve Test, you immediately gain the benefits of one dose of mandrake root. If this is a Critical Success, you gain the benefits of three doses instead. When under Hunger Pang’s effects, you do not suffer the drawbacks normally associated with taking mandrake root.,
  log.Println("applying Hunger Pangs")
  return state
}

var _ domain.Effect = &HungerPangs{}
