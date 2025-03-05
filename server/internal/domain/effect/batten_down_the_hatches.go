package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BattenDowntheHatches struct {
	name        string
	description string
}

func NewBattenDowntheHatches() *BattenDowntheHatches {
	return &BattenDowntheHatches{
		name:        "Batten Down the Hatches",
		description: "You may flip the results to succeed at Pilot Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Swim, you do not have to add the additional 1 Action Point cost.",
	}
}

func (e *BattenDowntheHatches) Name() string {
	return e.name
}

func (e *BattenDowntheHatches) Description() string {
	return e.description
}

func (e *BattenDowntheHatches) Applier() domain.Applier {
	return e.Apply
}

func (e *BattenDowntheHatches) Apply(state domain.GameState) domain.GameState {
	// - You may flip the results to succeed at Pilot Tests. When you succeed, it is always considered a Critical Success. In addition, when using the Movement subtype of Swim, you do not have to add the additional 1 Action Point cost.
	log.Println("applying Batten Down the Hatches")
	return state
}

var _ domain.Effect = &BattenDowntheHatches{}
