package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CultofPersonality struct {
  name string
  description string
}

func NewCultofPersonality() *CultofPersonality {
  return &CultofPersonality{
    name: "Cult of Personality",
    description: "When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Blessed Sacrament.",
  }
}

func (e *CultofPersonality) Name() string {
  return e.name
}

func (e *CultofPersonality) Description() string {
  return e.description
}

func (e *CultofPersonality) Applier() domain.Applier {
  return e.Apply
}

func (e *CultofPersonality) Apply(state domain.State) domain.State {
  // - When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Blessed Sacrament.
  log.Println("applying Cult of Personality")
  return state
}

var _ domain.Effect = &CultofPersonality{}
