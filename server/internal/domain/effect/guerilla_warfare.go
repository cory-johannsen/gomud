package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GuerillaWarfare struct {
  name string
  description string
}

func NewGuerillaWarfare() *GuerillaWarfare {
  return &GuerillaWarfare{
    name: "Guerilla Warfare",
    description: "You never provoke Opportunity Attacks with Movement Actions or any other action you take. In addition, any Movement Action a foe takes while Engaged with you immediately provokes an Opportunity Attack from you.",
  }
}

func (e *GuerillaWarfare) Name() string {
  return e.name
}

func (e *GuerillaWarfare) Description() string {
  return e.description
}

func (e *GuerillaWarfare) Applier() domain.Applier {
  return e.Apply
}

func (e *GuerillaWarfare) Apply(state domain.State) domain.State {
  // - You never provoke Opportunity Attacks with Movement Actions or any other action you take. In addition, any Movement Action a foe takes while Engaged with you immediately provokes an Opportunity Attack from you.
  log.Println("applying Guerilla Warfare")
  return state
}

var _ domain.Effect = &GuerillaWarfare{}
