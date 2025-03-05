package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type DungeonsDeep struct {
	name        string
	description string
}

func NewDungeonsDeep() *DungeonsDeep {
	return &DungeonsDeep{
		name:        "Dungeons Deep",
		description: "You leave no trace of your passing in underground areas or caves whatsoever, unless discovered by Magick or at a Critically Successful Awareness Test.,",
	}
}

func (e *DungeonsDeep) Name() string {
	return e.name
}

func (e *DungeonsDeep) Description() string {
	return e.description
}

func (e *DungeonsDeep) Applier() domain.Applier {
	return e.Apply
}

func (e *DungeonsDeep) Apply(state domain.GameState) domain.GameState {
	// - You leave no trace of your passing in underground areas or caves whatsoever, unless discovered by Magick or at a Critically Successful Awareness Test.,
	log.Println("applying Dungeons Deep")
	return state
}

var _ domain.Effect = &DungeonsDeep{}
