package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DislocatedShoulder struct {
  name string
  description string
}

func NewDislocatedShoulder() *DislocatedShoulder {
  return &DislocatedShoulder{
    name: "Dislocated Shoulder",
    description: "Until fully Recuperated, you start your Turn with 1 less AP.",
  }
}

func (e *DislocatedShoulder) Name() string {
  return e.name
}

func (e *DislocatedShoulder) Description() string {
  return e.description
}

func (e *DislocatedShoulder) Applier() domain.Applier {
  return e.Apply
}

func (e *DislocatedShoulder) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you start your Turn with 1 less AP.
  log.Println("applying Dislocated Shoulder")
  return state
}

var _ domain.Effect = &DislocatedShoulder{}
