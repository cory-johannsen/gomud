package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ISelltheDead struct {
  name string
  description string
}

func NewISelltheDead() *ISelltheDead {
  return &ISelltheDead{
    name: "I Sell the Dead",
    description: "When facing Supernatural creatures, you always succeed at Resolve Tests to save yourself from Stress and Fear provoked by them. In addition, you are immune to specific Diseases such as Red Death and Tomb Rot.",
  }
}

func (e *ISelltheDead) Name() string {
  return e.name
}

func (e *ISelltheDead) Description() string {
  return e.description
}

func (e *ISelltheDead) Applier() domain.Applier {
  return e.Apply
}

func (e *ISelltheDead) Apply(state domain.State) domain.State {
  // - When facing Supernatural creatures, you always succeed at Resolve Tests to save yourself from Stress and Fear provoked by them. In addition, you are immune to specific Diseases such as Red Death and Tomb Rot.
  log.Println("applying I Sell the Dead")
  return state
}

var _ domain.Effect = &ISelltheDead{}
