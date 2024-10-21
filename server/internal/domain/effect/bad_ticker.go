package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BadTicker struct {
  name string
  description string
}

func NewBadTicker() *BadTicker {
  return &BadTicker{
    name: "Bad Ticker",
    description: "Whenever you fail to Resist against Stress, Fear or Terror, you gain 3 additional Corruption.",
  }
}

func (e *BadTicker) Name() string {
  return e.name
}

func (e *BadTicker) Description() string {
  return e.description
}

func (e *BadTicker) Applier() domain.Applier {
  return e.Apply
}

func (e *BadTicker) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Bad Ticker")
  return state
}

var _ domain.Effect = &BadTicker{}
