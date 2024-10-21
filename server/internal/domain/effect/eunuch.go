package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Eunuch struct {
  name string
  description string
}

func NewEunuch() *Eunuch {
  return &Eunuch{
    name: "Eunuch",
    description: "You are immune to the charms and seduction by those who find you attractive, and unable to have children. However, being made victim to these same charms and seduction causes you to suffer 1 Corruption.",
  }
}

func (e *Eunuch) Name() string {
  return e.name
}

func (e *Eunuch) Description() string {
  return e.description
}

func (e *Eunuch) Applier() domain.Applier {
  return e.Apply
}

func (e *Eunuch) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Eunuch")
  return state
}

var _ domain.Effect = &Eunuch{}
