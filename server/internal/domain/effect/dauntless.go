package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Dauntless struct {
  name string
  description string
}

func NewDauntless() *Dauntless {
  return &Dauntless{
    name: "Dauntless",
    description: "Effect1",
  }
}

func (e *Dauntless) Name() string {
  return e.name
}

func (e *Dauntless) Description() string {
  return e.description
}

func (e *Dauntless) Applier() domain.Applier {
  return e.Apply
}

func (e *Dauntless) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Dauntless")
  return state
}

var _ domain.Effect = &Dauntless{}
