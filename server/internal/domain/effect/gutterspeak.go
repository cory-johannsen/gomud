package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Gutterspeak struct {
	name        string
	description string
}

func NewGutterspeak() *Gutterspeak {
	return &Gutterspeak{
		name:        "Gutterspeak",
		description: "You may use any Fellowship-based Skill to directly communicate with Small Animals (specifically moles, rats, rabbits, rodents, snakes, weasels and dogs). This interaction is not always beneficial - many creatures are skittish, ignorant, confused, cunningly deceptive or even openly aggressive. This empathy acts as an all but supernatural type of communication; the complexity determined by the GM. In addition, you are immune to specific Diseases such as Filth Fever and Grey Plague.",
	}
}

func (e *Gutterspeak) Name() string {
	return e.name
}

func (e *Gutterspeak) Description() string {
	return e.description
}

func (e *Gutterspeak) Applier() domain.Applier {
	return e.Apply
}

func (e *Gutterspeak) Apply(state domain.GameState) domain.GameState {
	// - You may use any Fellowship-based Skill to directly communicate with Small Animals (specifically moles, rats, rabbits, rodents, snakes, weasels and dogs). This interaction is not always beneficial - many creatures are skittish, ignorant, confused, cunningly deceptive or even openly aggressive. This empathy acts as an all but supernatural type of communication; the complexity determined by the GM. In addition, you are immune to specific Diseases such as Filth Fever and Grey Plague.
	log.Println("applying Gutterspeak")
	return state
}

var _ domain.Effect = &Gutterspeak{}
