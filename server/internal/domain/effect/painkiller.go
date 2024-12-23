package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Painkiller struct {
  name string
  description string
}

func NewPainkiller() *Painkiller {
  return &Painkiller{
    name: "Painkiller",
    description: "Work with your GM to select a single Addiction your Character begins play with.",
  }
}

func (e *Painkiller) Name() string {
  return e.name
}

func (e *Painkiller) Description() string {
  return e.description
}

func (e *Painkiller) Applier() domain.Applier {
  return e.Apply
}

func (e *Painkiller) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Painkiller")
  return state
}

var _ domain.Effect = &Painkiller{}
