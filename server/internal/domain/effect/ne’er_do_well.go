package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type NeerDoWell struct {
  name string
  description string
}

func NewNeerDoWell() *NeerDoWell {
  return &NeerDoWell{
    name: "Ne’er Do Well",
    description: "You cannot Assist others’ Skill Tests.",
  }
}

func (e *NeerDoWell) Name() string {
  return e.name
}

func (e *NeerDoWell) Description() string {
  return e.description
}

func (e *NeerDoWell) Applier() domain.Applier {
  return e.Apply
}

func (e *NeerDoWell) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Ne’er Do Well")
  return state
}

var _ domain.Effect = &NeerDoWell{}
