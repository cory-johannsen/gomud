package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SecondSkin struct {
  name string
  description string
}

func NewSecondSkin() *SecondSkin {
  return &SecondSkin{
    name: "Second Skin",
    description: "You can Dodge when wearing armor with the Heavy Quality.",
  }
}

func (e *SecondSkin) Name() string {
  return e.name
}

func (e *SecondSkin) Description() string {
  return e.description
}

func (e *SecondSkin) Applier() domain.Applier {
  return e.Apply
}

func (e *SecondSkin) Apply(state domain.State) domain.State {
  // - You can Dodge when wearing armor with the Heavy Quality.
  log.Println("applying Second Skin")
  return state
}

var _ domain.Effect = &SecondSkin{}
