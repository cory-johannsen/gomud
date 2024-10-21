package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CruisinforaBruisin struct {
  name string
  description string
}

func NewCruisinforaBruisin() *CruisinforaBruisin {
  return &CruisinforaBruisin{
    name: "Cruisin' for a Bruisin'",
    description: "Effect1",
  }
}

func (e *CruisinforaBruisin) Name() string {
  return e.name
}

func (e *CruisinforaBruisin) Description() string {
  return e.description
}

func (e *CruisinforaBruisin) Applier() domain.Applier {
  return e.Apply
}

func (e *CruisinforaBruisin) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Cruisin' for a Bruisin'")
  return state
}

var _ domain.Effect = &CruisinforaBruisin{}
