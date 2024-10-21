package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SplinteredElbow struct {
  name string
  description string
}

func NewSplinteredElbow() *SplinteredElbow {
  return &SplinteredElbow{
    name: "Splintered Elbow",
    description: "Until fully Recuperated, you cannot use one of your arms as in pain. You must undergo a successful surgery or suffer the consequences. Once a Splintered Elbow has undergone a failed surgery, you can no longer use two-handed ranged weapons. If you already have this Drawback, you permanently lose 9% Combat.",
  }
}

func (e *SplinteredElbow) Name() string {
  return e.name
}

func (e *SplinteredElbow) Description() string {
  return e.description
}

func (e *SplinteredElbow) Applier() domain.Applier {
  return e.Apply
}

func (e *SplinteredElbow) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot use one of your arms as in pain. You must undergo a successful surgery or suffer the consequences. Once a Splintered Elbow has undergone a failed surgery, you can no longer use two-handed ranged weapons. If you already have this Drawback, you permanently lose 9% Combat.
  log.Println("applying Splintered Elbow")
  return state
}

var _ domain.Effect = &SplinteredElbow{}
