package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TheFullMonty struct {
  name string
  description string
}

func NewTheFullMonty() *TheFullMonty {
  return &TheFullMonty{
    name: "The Full Monty",
    description: "All Skill Ranks you acquire in the Charm Skill modify your Base Chance by +20, instead of +10. In addition, when you perform acts of carnal knowledge with another, you both ‘get lucky’, gaining 1 Fortune Point each to personally use. This benefit cannot be given or gained more than once per day.",
  }
}

func (e *TheFullMonty) Name() string {
  return e.name
}

func (e *TheFullMonty) Description() string {
  return e.description
}

func (e *TheFullMonty) Applier() domain.Applier {
  return e.Apply
}

func (e *TheFullMonty) Apply(state domain.State) domain.State {
  // - All Skill Ranks you acquire in the Charm Skill modify your Base Chance by +20, instead of +10. In addition, when you perform acts of carnal knowledge with another, you both ‘get lucky’, gaining 1 Fortune Point each to personally use. This benefit cannot be given or gained more than once per day.
  log.Println("applying The Full Monty")
  return state
}

var _ domain.Effect = &TheFullMonty{}
