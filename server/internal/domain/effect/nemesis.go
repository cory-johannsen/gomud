package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Nemesis struct {
  name string
  description string
}

func NewNemesis() *Nemesis {
  return &Nemesis{
    name: "Nemesis",
    description: "When confronted by your Nemesis, you cannot sacrifice Fate or Fortune Points, as their presence confounds you. Your Nemesis is determined by the GM, who will likely take your ideas about your Character’s history or past rivals into consideration.",
  }
}

func (e *Nemesis) Name() string {
  return e.name
}

func (e *Nemesis) Description() string {
  return e.description
}

func (e *Nemesis) Applier() domain.Applier {
  return e.Apply
}

func (e *Nemesis) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Nemesis")
  return state
}

var _ domain.Effect = &Nemesis{}
