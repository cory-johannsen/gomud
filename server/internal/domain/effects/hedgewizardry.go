package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Hedgewizardry struct {
  Name string
  Description string
}

func (e *Hedgewizardry) Apply(state domain.State) domain.State {
  // - When casting Generalist Magick – at your option – you may automatically succeed at the Incantation Test, but move one step down the Peril Condition Track negatively. You also understand how to use the Ritual of Magick Circle.
  log.Println("applying Hedgewizardry")
  return state
}
