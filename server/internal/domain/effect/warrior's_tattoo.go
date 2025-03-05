package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type WarriorsTattoo struct {
	name        string
	description string
}

func NewWarriorsTattoo() *WarriorsTattoo {
	return &WarriorsTattoo{
		name:        "Warrior's Tattoo",
		description: "Add the Punishing Quality to any Martial Melee or Martial Ranged weapon you wield, and +1 to Total Damage with these same weapons.,",
	}
}

func (e *WarriorsTattoo) Name() string {
	return e.name
}

func (e *WarriorsTattoo) Description() string {
	return e.description
}

func (e *WarriorsTattoo) Applier() domain.Applier {
	return e.Apply
}

func (e *WarriorsTattoo) Apply(state domain.GameState) domain.GameState {
	// - Add the Punishing Quality to any Martial Melee or Martial Ranged weapon you wield, and +1 to Total Damage with these same weapons.,
	log.Println("applying Warrior's Tattoo")
	return state
}

var _ domain.Effect = &WarriorsTattoo{}
