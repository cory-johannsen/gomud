package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Oathkeeper struct {
  name string
  description string
}

func NewOathkeeper() *Oathkeeper {
  return &Oathkeeper{
    name: "Oathkeeper",
    description: "All armor, shields and weapons you forge with the Tradecraft Skill automatically gains the Castle-forged Quality.,",
  }
}

func (e *Oathkeeper) Name() string {
  return e.name
}

func (e *Oathkeeper) Description() string {
  return e.description
}

func (e *Oathkeeper) Applier() domain.Applier {
  return e.Apply
}

func (e *Oathkeeper) Apply(state domain.State) domain.State {
  // - All armor, shields and weapons you forge with the Tradecraft Skill automatically gains the Castle-forged Quality.,
  log.Println("applying Oathkeeper")
  return state
}

var _ domain.Effect = &Oathkeeper{}
