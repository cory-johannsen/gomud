package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MilitaryFormation struct {
  name string
  description string
}

func NewMilitaryFormation() *MilitaryFormation {
  return &MilitaryFormation{
    name: "Military Formation",
    description: "When you successfully use Inspiring Words in combat, both you and the allies you inspire raise their Initiative by 3.",
  }
}

func (e *MilitaryFormation) Name() string {
  return e.name
}

func (e *MilitaryFormation) Description() string {
  return e.description
}

func (e *MilitaryFormation) Applier() domain.Applier {
  return e.Apply
}

func (e *MilitaryFormation) Apply(state domain.State) domain.State {
  // - When you successfully use Inspiring Words in combat, both you and the allies you inspire raise their Initiative by 3.
  log.Println("applying Military Formation")
  return state
}

var _ domain.Effect = &MilitaryFormation{}
