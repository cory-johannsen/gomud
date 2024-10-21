package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MeditativeHealing struct {
  name string
  description string
}

func NewMeditativeHealing() *MeditativeHealing {
  return &MeditativeHealing{
    name: "Meditative Healing",
    description: "Effect1",
  }
}

func (e *MeditativeHealing) Name() string {
  return e.name
}

func (e *MeditativeHealing) Description() string {
  return e.description
}

func (e *MeditativeHealing) Applier() domain.Applier {
  return e.Apply
}

func (e *MeditativeHealing) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Meditative Healing")
  return state
}

var _ domain.Effect = &MeditativeHealing{}
