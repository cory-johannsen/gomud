package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type NostrumRemedium struct {
	name        string
	description string
}

func NewNostrumRemedium() *NostrumRemedium {
	return &NostrumRemedium{
		name:        "Nostrum Remedium",
		description: "You may flip the results to succeed at Alchemy Tests. When you succeed, it is always considered a Critical Success. In addition, you never suffer Peril as a result of failed or Critically Failed Alchemy Tests.",
	}
}

func (e *NostrumRemedium) Name() string {
	return e.name
}

func (e *NostrumRemedium) Description() string {
	return e.description
}

func (e *NostrumRemedium) Applier() domain.Applier {
	return e.Apply
}

func (e *NostrumRemedium) Apply(state domain.GameState) domain.GameState {
	// - You may flip the results to succeed at Alchemy Tests. When you succeed, it is always considered a Critical Success. In addition, you never suffer Peril as a result of failed or Critically Failed Alchemy Tests.
	log.Println("applying Nostrum Remedium")
	return state
}

var _ domain.Effect = &NostrumRemedium{}
