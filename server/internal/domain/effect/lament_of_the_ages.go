package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LamentoftheAges struct {
  name string
  description string
}

func NewLamentoftheAges() *LamentoftheAges {
  return &LamentoftheAges{
    name: "Lament of the Ages",
    description: "Whenever combat begins, select one foe. If the foe can clearly see and hear you, they must reduce all Damage they do to you by your [FB].,",
  }
}

func (e *LamentoftheAges) Name() string {
  return e.name
}

func (e *LamentoftheAges) Description() string {
  return e.description
}

func (e *LamentoftheAges) Applier() domain.Applier {
  return e.Apply
}

func (e *LamentoftheAges) Apply(state domain.State) domain.State {
  // - Whenever combat begins, select one foe. If the foe can clearly see and hear you, they must reduce all Damage they do to you by your [FB].,
  log.Println("applying Lament of the Ages")
  return state
}

var _ domain.Effect = &LamentoftheAges{}
