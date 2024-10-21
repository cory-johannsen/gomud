package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type IAmTheLaw struct {
  name string
  description string
}

func NewIAmTheLaw() *IAmTheLaw {
  return &IAmTheLaw{
    name: "I Am The Law!",
    description: "You may flip the results to succeed at Intimidate Tests. When you succeed, it is always considered a Critical Success. Furthermore, you always influence a number of people with the Intimidate Skill equal to three times your [BB] – this includes use of Litany of Hatred during combat.",
  }
}

func (e *IAmTheLaw) Name() string {
  return e.name
}

func (e *IAmTheLaw) Description() string {
  return e.description
}

func (e *IAmTheLaw) Applier() domain.Applier {
  return e.Apply
}

func (e *IAmTheLaw) Apply(state domain.State) domain.State {
  // - You may flip the results to succeed at Intimidate Tests. When you succeed, it is always considered a Critical Success. Furthermore, you always influence a number of people with the Intimidate Skill equal to three times your [BB] – this includes use of Litany of Hatred during combat.
  log.Println("applying I Am The Law!")
  return state
}

var _ domain.Effect = &IAmTheLaw{}
