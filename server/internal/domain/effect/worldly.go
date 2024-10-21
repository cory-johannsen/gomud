package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Worldly struct {
  name string
  description string
}

func NewWorldly() *Worldly {
  return &Worldly{
    name: "Worldly",
    description: "When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.",
  }
}

func (e *Worldly) Name() string {
  return e.name
}

func (e *Worldly) Description() string {
  return e.description
}

func (e *Worldly) Applier() domain.Applier {
  return e.Apply
}

func (e *Worldly) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Worldly")
  return state
}

var _ domain.Effect = &Worldly{}
