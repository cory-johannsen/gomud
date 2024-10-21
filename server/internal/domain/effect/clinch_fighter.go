package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ClinchFighter struct {
  name string
  description string
}

func NewClinchFighter() *ClinchFighter {
  return &ClinchFighter{
    name: "Clinch Fighter",
    description: "You penalize your foe’s ability to Resist a Chokehold and Dirty Tricks by a -10 Base Chance to their Skill Test.  In addition, when you use a Chokehold, add an additional 1D10 to determine how much physical Peril you inflict.",
  }
}

func (e *ClinchFighter) Name() string {
  return e.name
}

func (e *ClinchFighter) Description() string {
  return e.description
}

func (e *ClinchFighter) Applier() domain.Applier {
  return e.Apply
}

func (e *ClinchFighter) Apply(state domain.State) domain.State {
  // - You penalize your foe’s ability to Resist a Chokehold and Dirty Tricks by a -10 Base Chance to their Skill Test.  In addition, when you use a Chokehold, add an additional 1D10 to determine how much physical Peril you inflict.
  log.Println("applying Clinch Fighter")
  return state
}

var _ domain.Effect = &ClinchFighter{}
