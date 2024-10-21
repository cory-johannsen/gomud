package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Determination struct {
  name string
  description string
}

func NewDetermination() *Determination {
  return &Determination{
    name: "Determination",
    description: "When you attempt an Extended Test to take your time, you gain an additional +10 Base Chance to your Skill Test.",
  }
}

func (e *Determination) Name() string {
  return e.name
}

func (e *Determination) Description() string {
  return e.description
}

func (e *Determination) Applier() domain.Applier {
  return e.Apply
}

func (e *Determination) Apply(state domain.State) domain.State {
  // - When you attempt an Extended Test to take your time, you gain an additional +10 Base Chance to your Skill Test.
  log.Println("applying Determination")
  return state
}

var _ domain.Effect = &Determination{}
