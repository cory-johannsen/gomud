package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Chatty struct {
  name string
  description string
}

func NewChatty() *Chatty {
  return &Chatty{
    name: "Chatty",
    description: "'When using Flair to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'",
  }
}

func (e *Chatty) Name() string {
  return e.name
}

func (e *Chatty) Description() string {
  return e.description
}

func (e *Chatty) Applier() domain.Applier {
  return e.Apply
}

func (e *Chatty) Apply(state domain.State) domain.State {
  // - 'When using Flair to solve a problem, flip the roll to Success.  Otherwise flip the roll to Failure.'
  log.Println("applying Chatty")
  return state
}

var _ domain.Effect = &Chatty{}
