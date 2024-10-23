package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Bewitching struct {
  name string
  description string
}

func NewBewitching() *Bewitching {
  return &Bewitching{
    name: "Bewitching",
    description: "Whenever you cast Magick to control minds or use the Charm Skill on others, foes must flip their results to fail to Resist its effects. However, this only works if your foe can both see and hear you.,",
  }
}

func (e *Bewitching) Name() string {
  return e.name
}

func (e *Bewitching) Description() string {
  return e.description
}

func (e *Bewitching) Applier() domain.Applier {
  return e.Apply
}

func (e *Bewitching) Apply(state domain.State) domain.State {
  // - Whenever you cast Magick to control minds or use the Charm Skill on others, foes must flip their results to fail to Resist its effects. However, this only works if your foe can both see and hear you.,
  log.Println("applying Bewitching")
  return state
}

var _ domain.Effect = &Bewitching{}
