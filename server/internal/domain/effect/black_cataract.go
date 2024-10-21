package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BlackCataract struct {
  name string
  description string
}

func NewBlackCataract() *BlackCataract {
  return &BlackCataract{
    name: "Black Cataract",
    description: "Whenever you miss with Attack Actions using a ranged weapon, you must re-roll the result with the same Difficulty Rating. If it is a success, you strike a random ally who is Engaged with your target.",
  }
}

func (e *BlackCataract) Name() string {
  return e.name
}

func (e *BlackCataract) Description() string {
  return e.description
}

func (e *BlackCataract) Applier() domain.Applier {
  return e.Apply
}

func (e *BlackCataract) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Black Cataract")
  return state
}

var _ domain.Effect = &BlackCataract{}
