package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountedDefense struct {
  name string
  description string
}

func NewMountedDefense() *MountedDefense {
  return &MountedDefense{
    name: "Mounted Defense",
    description: "When fighting on horseback or atop a vehicle like a cart, coach or wagon, you can substitute your Drive or Ride Skill in place of Coordination when you Dodge or Parry. In addition, you always succeed at Skill Tests to retain control when your mount or the animal that’s pulling your vehicle suffers Damage.",
  }
}

func (e *MountedDefense) Name() string {
  return e.name
}

func (e *MountedDefense) Description() string {
  return e.description
}

func (e *MountedDefense) Applier() domain.Applier {
  return e.Apply
}

func (e *MountedDefense) Apply(state domain.State) domain.State {
  // - When fighting on horseback or atop a vehicle like a cart, coach or wagon, you can substitute your Drive or Ride Skill in place of Coordination when you Dodge or Parry. In addition, you always succeed at Skill Tests to retain control when your mount or the animal that’s pulling your vehicle suffers Damage.
  log.Println("applying Mounted Defense")
  return state
}

var _ domain.Effect = &MountedDefense{}
