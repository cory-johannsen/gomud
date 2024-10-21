package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PliersScalpelBonesaw struct {
  Name string
  Description string
}

func (e *PliersScalpelBonesaw) Apply(state domain.State) domain.State {
  // - Whenever you successfully bind wounds to heal Damage, your patient moves one additional step up the Damage Condition Track positively. In addition, you always succeed at Heal Tests to prepare bandages.
  log.Println("applying Pliers... Scalpel... Bonesaw")
  return state
}
