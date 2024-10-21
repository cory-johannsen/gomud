package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type VeteransEye struct {
  name string
  description string
}

func NewVeteransEye() *VeteransEye {
  return &VeteransEye{
    name: "Veteran’s Eye",
    description: "Reduce the Distance for all ranged weapons you use by -3 (to a minimum of 1).",
  }
}

func (e *VeteransEye) Name() string {
  return e.name
}

func (e *VeteransEye) Description() string {
  return e.description
}

func (e *VeteransEye) Applier() domain.Applier {
  return e.Apply
}

func (e *VeteransEye) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Veteran’s Eye")
  return state
}

var _ domain.Effect = &VeteransEye{}
