package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Sprint struct {
  name string
  description string
}

func NewSprint() *Sprint {
  return &Sprint{
    name: "Sprint",
    description: "When you Charge or Run, you cannot be struck by attacks made with ranged weapons until your next Turn.",
  }
}

func (e *Sprint) Name() string {
  return e.name
}

func (e *Sprint) Description() string {
  return e.description
}

func (e *Sprint) Applier() domain.Applier {
  return e.Apply
}

func (e *Sprint) Apply(state domain.State) domain.State {
  // - When you Charge or Run, you cannot be struck by attacks made with ranged weapons until your next Turn.
  log.Println("applying Sprint")
  return state
}

var _ domain.Effect = &Sprint{}
