package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HardToKill struct {
  name string
  description string
}

func NewHardToKill() *HardToKill {
  return &HardToKill{
    name: "Hard To Kill",
    description: "When you are Grievously Wounded, temporarily add 3 to your Damage Threshold.",
  }
}

func (e *HardToKill) Name() string {
  return e.name
}

func (e *HardToKill) Description() string {
  return e.description
}

func (e *HardToKill) Applier() domain.Applier {
  return e.Apply
}

func (e *HardToKill) Apply(state domain.State) domain.State {
  // - When you are Grievously Wounded, temporarily add 3 to your Damage Threshold.
  log.Println("applying Hard To Kill")
  return state
}

var _ domain.Effect = &HardToKill{}
