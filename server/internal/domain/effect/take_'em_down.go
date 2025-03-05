package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type TakeEmDown struct {
	name        string
	description string
}

func NewTakeEmDown() *TakeEmDown {
	return &TakeEmDown{
		name:        "Take 'Em Down",
		description: "At Distance, you can substitute either your Simple Ranged or Martial Ranged Skills when you attempt to use Disarm, Stunning Blow or Takedown. Note that you do not inflict Damage when you attempt such maneuvers. You must be armed with a loaded ranged weapon in order to Take ‘Em Down.",
	}
}

func (e *TakeEmDown) Name() string {
	return e.name
}

func (e *TakeEmDown) Description() string {
	return e.description
}

func (e *TakeEmDown) Applier() domain.Applier {
	return e.Apply
}

func (e *TakeEmDown) Apply(state domain.GameState) domain.GameState {
	// - At Distance, you can substitute either your Simple Ranged or Martial Ranged Skills when you attempt to use Disarm, Stunning Blow or Takedown. Note that you do not inflict Damage when you attempt such maneuvers. You must be armed with a loaded ranged weapon in order to Take ‘Em Down.
	log.Println("applying Take 'Em Down")
	return state
}

var _ domain.Effect = &TakeEmDown{}
