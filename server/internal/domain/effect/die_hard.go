package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DieHard struct {
  name string
  description string
}

func NewDieHard() *DieHard {
  return &DieHard{
    name: "Die Hard",
    description: "You reduce the time to Recuperate from your own Injuries by three days, to a minimum of one day. In addition, you never Bleed.",
  }
}

func (e *DieHard) Name() string {
  return e.name
}

func (e *DieHard) Description() string {
  return e.description
}

func (e *DieHard) Applier() domain.Applier {
  return e.Apply
}

func (e *DieHard) Apply(state domain.State) domain.State {
  // - You reduce the time to Recuperate from your own Injuries by three days, to a minimum of one day. In addition, you never Bleed.
  log.Println("applying Die Hard")
  return state
}

var _ domain.Effect = &DieHard{}
