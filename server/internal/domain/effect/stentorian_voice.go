package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type StentorianVoice struct {
	name        string
	description string
}

func NewStentorianVoice() *StentorianVoice {
	return &StentorianVoice{
		name:        "Stentorian Voice",
		description: "Whenever you use Fellowship-based Skills, refer to either your Brawn or Fellowship Primary Attribute (whichever is highest).,",
	}
}

func (e *StentorianVoice) Name() string {
	return e.name
}

func (e *StentorianVoice) Description() string {
	return e.description
}

func (e *StentorianVoice) Applier() domain.Applier {
	return e.Apply
}

func (e *StentorianVoice) Apply(state domain.GameState) domain.GameState {
	// - Whenever you use Fellowship-based Skills, refer to either your Brawn or Fellowship Primary Attribute (whichever is highest).,
	log.Println("applying Stentorian Voice")
	return state
}

var _ domain.Effect = &StentorianVoice{}
