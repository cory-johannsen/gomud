package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TortuousInquisition struct {
  name string
  description string
}

func NewTortuousInquisition() *TortuousInquisition {
  return &TortuousInquisition{
    name: "Tortuous Inquisition",
    description: "You do not suffer the ill-effect of Peril, until you are at “Ignore 3 Skill Ranks” on the Peril Condition Track.",
  }
}

func (e *TortuousInquisition) Name() string {
  return e.name
}

func (e *TortuousInquisition) Description() string {
  return e.description
}

func (e *TortuousInquisition) Applier() domain.Applier {
  return e.Apply
}

func (e *TortuousInquisition) Apply(state domain.State) domain.State {
  // - You do not suffer the ill-effect of Peril, until you are at “Ignore 3 Skill Ranks” on the Peril Condition Track.
  log.Println("applying Tortuous Inquisition")
  return state
}

var _ domain.Effect = &TortuousInquisition{}
