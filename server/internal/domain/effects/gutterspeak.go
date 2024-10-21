package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type Gutterspeak struct {
  Name string
  Description string
}

func (e *Gutterspeak) Apply(state domain.State) domain.State {
  // - You may use any Fellowship-based Skill to directly communicate with Small Animals (specifically moles, rats, rabbits, rodents, snakes, weasels and dogs). This interaction is not always beneficial - many creatures are skittish, ignorant, confused, cunningly deceptive or even openly aggressive. This empathy acts as an all but supernatural type of communication; the complexity determined by the GM. In addition, you are immune to specific Diseases such as Filth Fever and Grey Plague.
  log.Println("applying Gutterspeak")
  return state
}
