package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MountainAmongstMen struct {
  name string
  description string
}

func NewMountainAmongstMen() *MountainAmongstMen {
  return &MountainAmongstMen{
    name: "Mountain Amongst Men",
    description: "You are able to wield any two-handed melee weapon with one hand. This also means you may freely take advantage of the Adaptable Quality for weapons using only one hand. Finally, you will reference “91 to 100” on the Height table and will have a Husky build on the Build table.,",
  }
}

func (e *MountainAmongstMen) Name() string {
  return e.name
}

func (e *MountainAmongstMen) Description() string {
  return e.description
}

func (e *MountainAmongstMen) Applier() domain.Applier {
  return e.Apply
}

func (e *MountainAmongstMen) Apply(state domain.State) domain.State {
  // - You are able to wield any two-handed melee weapon with one hand. This also means you may freely take advantage of the Adaptable Quality for weapons using only one hand. Finally, you will reference “91 to 100” on the Height table and will have a Husky build on the Build table.,
  log.Println("applying Mountain Amongst Men")
  return state
}

var _ domain.Effect = &MountainAmongstMen{}
