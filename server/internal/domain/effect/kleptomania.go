package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Kleptomania struct {
  name string
  description string
}

func NewKleptomania() *Kleptomania {
  return &Kleptomania{
    name: "Kleptomania",
    description: "Whenever you are in a pinch and need one object that could realistically fit into your pocket, spend one Fortune Point to automatically find it within. For instance, you could use this to produce a single silver shilling (ss) from your pockets, but not several silver shillings (ss) without spending an equal number of Fortune Points.,",
  }
}

func (e *Kleptomania) Name() string {
  return e.name
}

func (e *Kleptomania) Description() string {
  return e.description
}

func (e *Kleptomania) Applier() domain.Applier {
  return e.Apply
}

func (e *Kleptomania) Apply(state domain.State) domain.State {
  // - Whenever you are in a pinch and need one object that could realistically fit into your pocket, spend one Fortune Point to automatically find it within. For instance, you could use this to produce a single silver shilling (ss) from your pockets, but not several silver shillings (ss) without spending an equal number of Fortune Points.,
  log.Println("applying Kleptomania")
  return state
}

var _ domain.Effect = &Kleptomania{}
