package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MeditativeHealing struct {
	name        string
	description string
}

func NewMeditativeHealing() *MeditativeHealing {
	return &MeditativeHealing{
		name:        "Meditative Healing",
		description: "Whenever you awaken from a good night’s sleep, you move one step up the Damage Condition Track positively. This means that if you were Grievously Wounded, you are now Seriously Wounded instead. Note that this does not ‘cure’ Injuries.,",
	}
}

func (e *MeditativeHealing) Name() string {
	return e.name
}

func (e *MeditativeHealing) Description() string {
	return e.description
}

func (e *MeditativeHealing) Applier() domain.Applier {
	return e.Apply
}

func (e *MeditativeHealing) Apply(state domain.GameState) domain.GameState {
	// - Whenever you awaken from a good night’s sleep, you move one step up the Damage Condition Track positively. This means that if you were Grievously Wounded, you are now Seriously Wounded instead. Note that this does not ‘cure’ Injuries.,
	log.Println("applying Meditative Healing")
	return state
}

var _ domain.Effect = &MeditativeHealing{}
