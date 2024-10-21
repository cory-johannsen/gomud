package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HolyRoller struct {
  name string
  description string
}

func NewHolyRoller() *HolyRoller {
  return &HolyRoller{
    name: "Holy Roller",
    description: "Adjust your Damage Threshold by +3, but only when you aren’t wearing armor.",
  }
}

func (e *HolyRoller) Name() string {
  return e.name
}

func (e *HolyRoller) Description() string {
  return e.description
}

func (e *HolyRoller) Applier() domain.Applier {
  return e.Apply
}

func (e *HolyRoller) Apply(state domain.State) domain.State {
  // - Adjust your Damage Threshold by +3, but only when you aren’t wearing armor.
  log.Println("applying Holy Roller")
  return state
}

var _ domain.Effect = &HolyRoller{}
