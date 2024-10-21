package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MightyThews struct {
  name string
  description string
}

func NewMightyThews() *MightyThews {
  return &MightyThews{
    name: "Mighty Thews",
    description: "Effect1",
  }
}

func (e *MightyThews) Name() string {
  return e.name
}

func (e *MightyThews) Description() string {
  return e.description
}

func (e *MightyThews) Applier() domain.Applier {
  return e.Apply
}

func (e *MightyThews) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Mighty Thews")
  return state
}

var _ domain.Effect = &MightyThews{}
