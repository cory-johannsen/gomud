package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SwaggerWagon struct {
  Name string
  Description string
}

func (e *SwaggerWagon) Apply(state domain.State) domain.State {
  // - Add 3 to your Encumbrance Limit. In addition, you may optionally substitute your Brawn in place of your Willpower when it comes to determining your Peril Threshold.
  log.Println("applying Swagger Wagon")
  return state
}
