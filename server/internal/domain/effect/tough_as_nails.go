package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ToughAsNails struct {
  name string
  description string
}

func NewToughAsNails() *ToughAsNails {
  return &ToughAsNails{
    name: "Tough As Nails",
    description: "You no longer suffer Moderate Injuries, remaining uninjured as a result.",
  }
}

func (e *ToughAsNails) Name() string {
  return e.name
}

func (e *ToughAsNails) Description() string {
  return e.description
}

func (e *ToughAsNails) Applier() domain.Applier {
  return e.Apply
}

func (e *ToughAsNails) Apply(state domain.State) domain.State {
  // - You no longer suffer Moderate Injuries, remaining uninjured as a result.
  log.Println("applying Tough As Nails")
  return state
}

var _ domain.Effect = &ToughAsNails{}
