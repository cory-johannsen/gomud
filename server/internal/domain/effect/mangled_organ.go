package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MangledOrgan struct {
  name string
  description string
}

func NewMangledOrgan() *MangledOrgan {
  return &MangledOrgan{
    name: "Mangled Organ",
    description: "Until fully Recuperated, you remain Incapacitated!. You must undergo a successful surgery or suffer the consequences. Once a Mangled Organ has undergone a failed surgery, you permanently gain the Eunuch Drawback. If you already have this Drawback, you permanently lose 9% Fellowship.",
  }
}

func (e *MangledOrgan) Name() string {
  return e.name
}

func (e *MangledOrgan) Description() string {
  return e.description
}

func (e *MangledOrgan) Applier() domain.Applier {
  return e.Apply
}

func (e *MangledOrgan) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you remain Incapacitated!. You must undergo a successful surgery or suffer the consequences. Once a Mangled Organ has undergone a failed surgery, you permanently gain the Eunuch Drawback. If you already have this Drawback, you permanently lose 9% Fellowship.
  log.Println("applying Mangled Organ")
  return state
}

var _ domain.Effect = &MangledOrgan{}
