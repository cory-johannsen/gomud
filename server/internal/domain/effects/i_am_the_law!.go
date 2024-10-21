package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type IAmTheLaw struct {
  Name string
  Description string
}

func (e *IAmTheLaw) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Intimidate Tests. When you succeed, it is always considered a Critical Success. Furthermore, you always influence a number of people with the Intimidate Skill equal to three times your [BB] – this includes use of Litany of Hatred during combat.
  log.Println("applying I Am The Law!")
  return state
}
