package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DealwiththeDevil struct {
  name string
  description string
}

func NewDealwiththeDevil() *DealwiththeDevil {
  return &DealwiththeDevil{
    name: "Deal with the Devil",
    description: "You begin gameplay with one permanent Chaos Rank.",
  }
}

func (e *DealwiththeDevil) Name() string {
  return e.name
}

func (e *DealwiththeDevil) Description() string {
  return e.description
}

func (e *DealwiththeDevil) Applier() domain.Applier {
  return e.Apply
}

func (e *DealwiththeDevil) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Deal with the Devil")
  return state
}

var _ domain.Effect = &DealwiththeDevil{}
