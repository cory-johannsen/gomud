package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WretchedPrankster struct {
  name string
  description string
}

func NewWretchedPrankster() *WretchedPrankster {
  return &WretchedPrankster{
    name: "Wretched Prankster",
    description: "Effect1",
  }
}

func (e *WretchedPrankster) Name() string {
  return e.name
}

func (e *WretchedPrankster) Description() string {
  return e.description
}

func (e *WretchedPrankster) Applier() domain.Applier {
  return e.Apply
}

func (e *WretchedPrankster) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Wretched Prankster")
  return state
}

var _ domain.Effect = &WretchedPrankster{}
