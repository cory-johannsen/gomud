package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LightSleeper struct {
  name string
  description string
}

func NewLightSleeper() *LightSleeper {
  return &LightSleeper{
    name: "Light Sleeper",
    description: "When you are sleeping, you cannot be Surprised or left Helpless.",
  }
}

func (e *LightSleeper) Name() string {
  return e.name
}

func (e *LightSleeper) Description() string {
  return e.description
}

func (e *LightSleeper) Applier() domain.Applier {
  return e.Apply
}

func (e *LightSleeper) Apply(state domain.State) domain.State {
  // - When you are sleeping, you cannot be Surprised or left Helpless.
  log.Println("applying Light Sleeper")
  return state
}

var _ domain.Effect = &LightSleeper{}
