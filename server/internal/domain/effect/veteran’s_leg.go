package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VeteransLeg struct {
  name string
  description string
}

func NewVeteransLeg() *VeteransLeg {
  return &VeteransLeg{
    name: "Veteran’s Leg",
    description: "You must reduce your Movement by 3 and cannot Run.",
  }
}

func (e *VeteransLeg) Name() string {
  return e.name
}

func (e *VeteransLeg) Description() string {
  return e.description
}

func (e *VeteransLeg) Applier() domain.Applier {
  return e.Apply
}

func (e *VeteransLeg) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Veteran’s Leg")
  return state
}

var _ domain.Effect = &VeteransLeg{}
