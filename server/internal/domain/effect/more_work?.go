package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MoreWork struct {
  name string
  description string
}

func NewMoreWork() *MoreWork {
  return &MoreWork{
    name: "More Work?",
    description: "Adjust your Peril Threshold by +3. In addition, you always recover your Peril Condition Track to Unhindered after resting.",
  }
}

func (e *MoreWork) Name() string {
  return e.name
}

func (e *MoreWork) Description() string {
  return e.description
}

func (e *MoreWork) Applier() domain.Applier {
  return e.Apply
}

func (e *MoreWork) Apply(state domain.State) domain.State {
  // - Adjust your Peril Threshold by +3. In addition, you always recover your Peril Condition Track to Unhindered after resting.
  log.Println("applying More Work?")
  return state
}

var _ domain.Effect = &MoreWork{}
