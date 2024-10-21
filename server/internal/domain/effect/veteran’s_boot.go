package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VeteransBoot struct {
  name string
  description string
}

func NewVeteransBoot() *VeteransBoot {
  return &VeteransBoot{
    name: "Veteran’s Boot",
    description: "You cannot Charge, Maneuver or Run with Movement Actions without spending an additional Action Point.",
  }
}

func (e *VeteransBoot) Name() string {
  return e.name
}

func (e *VeteransBoot) Description() string {
  return e.description
}

func (e *VeteransBoot) Applier() domain.Applier {
  return e.Apply
}

func (e *VeteransBoot) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Veteran’s Boot")
  return state
}

var _ domain.Effect = &VeteransBoot{}
