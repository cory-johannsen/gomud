package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Beguiler struct {
  name string
  description string
}

func NewBeguiler() *Beguiler {
  return &Beguiler{
    name: "Beguiler",
    description: "Reference either your [FB] or [BB] for the purposes of determining Damage Threshold. Additionally, you may substitute Guile in place of any Combat-based Skill when Parrying attacks.,",
  }
}

func (e *Beguiler) Name() string {
  return e.name
}

func (e *Beguiler) Description() string {
  return e.description
}

func (e *Beguiler) Applier() domain.Applier {
  return e.Apply
}

func (e *Beguiler) Apply(state domain.State) domain.State {
  // - Reference either your [FB] or [BB] for the purposes of determining Damage Threshold. Additionally, you may substitute Guile in place of any Combat-based Skill when Parrying attacks.,
  log.Println("applying Beguiler")
  return state
}

var _ domain.Effect = &Beguiler{}
