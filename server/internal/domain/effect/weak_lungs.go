package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WeakLungs struct {
  name string
  description string
}

func NewWeakLungs() *WeakLungs {
  return &WeakLungs{
    name: "Weak Lungs",
    description: "Whenever you suffer physical Peril, move one additional step down the Peril Condition Track negatively while suffering 1 Corruption.",
  }
}

func (e *WeakLungs) Name() string {
  return e.name
}

func (e *WeakLungs) Description() string {
  return e.description
}

func (e *WeakLungs) Applier() domain.Applier {
  return e.Apply
}

func (e *WeakLungs) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Weak Lungs")
  return state
}

var _ domain.Effect = &WeakLungs{}
