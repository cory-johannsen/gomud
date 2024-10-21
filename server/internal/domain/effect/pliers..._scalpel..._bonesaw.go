package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type PliersScalpelBonesaw struct {
  name string
  description string
}

func NewPliersScalpelBonesaw() *PliersScalpelBonesaw {
  return &PliersScalpelBonesaw{
    name: "Pliers... Scalpel... Bonesaw",
    description: "Whenever you successfully bind wounds to heal Damage, your patient moves one additional step up the Damage Condition Track positively. In addition, you always succeed at Heal Tests to prepare bandages.",
  }
}

func (e *PliersScalpelBonesaw) Name() string {
  return e.name
}

func (e *PliersScalpelBonesaw) Description() string {
  return e.description
}

func (e *PliersScalpelBonesaw) Applier() domain.Applier {
  return e.Apply
}

func (e *PliersScalpelBonesaw) Apply(state domain.State) domain.State {
  // - Whenever you successfully bind wounds to heal Damage, your patient moves one additional step up the Damage Condition Track positively. In addition, you always succeed at Heal Tests to prepare bandages.
  log.Println("applying Pliers... Scalpel... Bonesaw")
  return state
}

var _ domain.Effect = &PliersScalpelBonesaw{}
