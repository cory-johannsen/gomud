package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HueCry struct {
  name string
  description string
}

func NewHueCry() *HueCry {
  return &HueCry{
    name: "Hue & Cry",
    description: "When combat begins, roll 2D10, instead of 1D10, to determine your place in the Initiative Order.",
  }
}

func (e *HueCry) Name() string {
  return e.name
}

func (e *HueCry) Description() string {
  return e.description
}

func (e *HueCry) Applier() domain.Applier {
  return e.Apply
}

func (e *HueCry) Apply(state domain.State) domain.State {
  // - When combat begins, roll 2D10, instead of 1D10, to determine your place in the Initiative Order.
  log.Println("applying Hue & Cry")
  return state
}

var _ domain.Effect = &HueCry{}
