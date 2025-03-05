package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BeggarsBowl struct {
	name        string
	description string
}

func NewBeggarsBowl() *BeggarsBowl {
	return &BeggarsBowl{
		name:        "Beggars Bowl",
		description: "Whenever you fail a Guile or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.",
	}
}

func (e *BeggarsBowl) Name() string {
	return e.name
}

func (e *BeggarsBowl) Description() string {
	return e.description
}

func (e *BeggarsBowl) Applier() domain.Applier {
	return e.Apply
}

func (e *BeggarsBowl) Apply(state domain.GameState) domain.GameState {
	// - Whenever you fail a Guile or Intimidate Test, you may re-roll to generate a better result, but must accept the outcome.
	log.Println("applying Beggars Bowl")
	return state
}

var _ domain.Effect = &BeggarsBowl{}
