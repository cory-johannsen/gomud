package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MettleSteel struct {
  name string
  description string
}

func NewMettleSteel() *MettleSteel {
  return &MettleSteel{
    name: "Mettle & Steel",
    description: "At your option, you may substitute the Warfare Skill in place of Resolve. In addition, you may substitute Warfare in place of any Skill required to Resist Perilous Stunts.",
  }
}

func (e *MettleSteel) Name() string {
  return e.name
}

func (e *MettleSteel) Description() string {
  return e.description
}

func (e *MettleSteel) Applier() domain.Applier {
  return e.Apply
}

func (e *MettleSteel) Apply(state domain.State) domain.State {
  // - At your option, you may substitute the Warfare Skill in place of Resolve. In addition, you may substitute Warfare in place of any Skill required to Resist Perilous Stunts.
  log.Println("applying Mettle & Steel")
  return state
}

var _ domain.Effect = &MettleSteel{}
