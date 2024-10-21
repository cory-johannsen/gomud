package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SwaggerWagon struct {
  name string
  description string
}

func NewSwaggerWagon() *SwaggerWagon {
  return &SwaggerWagon{
    name: "Swagger Wagon",
    description: "Add 3 to your Encumbrance Limit. In addition, you may optionally substitute your Brawn in place of your Willpower when it comes to determining your Peril Threshold.",
  }
}

func (e *SwaggerWagon) Name() string {
  return e.name
}

func (e *SwaggerWagon) Description() string {
  return e.description
}

func (e *SwaggerWagon) Applier() domain.Applier {
  return e.Apply
}

func (e *SwaggerWagon) Apply(state domain.State) domain.State {
  // - Add 3 to your Encumbrance Limit. In addition, you may optionally substitute your Brawn in place of your Willpower when it comes to determining your Peril Threshold.
  log.Println("applying Swagger Wagon")
  return state
}

var _ domain.Effect = &SwaggerWagon{}
