package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CultofPersonality struct {
  Name string
  Description string
}

func (e *CultofPersonality) Apply(state domain.State) domain.State {
  // - When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Blessed Sacrament.
  log.Println("applying Cult of Personality")
  return state
}
