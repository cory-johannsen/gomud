package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type RoadtoEnlightenment struct {
	name        string
	description string
}

func NewRoadtoEnlightenment() *RoadtoEnlightenment {
	return &RoadtoEnlightenment{
		name:        "Road to Enlightenment",
		description: "When you suffer 1 to 3 Corruption, make a (Routine +10%) Resolve Test. If you suffer 4 to 6 Corruption, make a (Standard +/-0%) Resolve Test. If you suffer 7 to 9 Corruption, make a (Challenging -10%) Resolve Test. If successful, you suffer no Corruption.",
	}
}

func (e *RoadtoEnlightenment) Name() string {
	return e.name
}

func (e *RoadtoEnlightenment) Description() string {
	return e.description
}

func (e *RoadtoEnlightenment) Applier() domain.Applier {
	return e.Apply
}

func (e *RoadtoEnlightenment) Apply(state domain.GameState) domain.GameState {
	// - When you suffer 1 to 3 Corruption, make a (Routine +10%) Resolve Test. If you suffer 4 to 6 Corruption, make a (Standard +/-0%) Resolve Test. If you suffer 7 to 9 Corruption, make a (Challenging -10%) Resolve Test. If successful, you suffer no Corruption.
	log.Println("applying Road to Enlightenment")
	return state
}

var _ domain.Effect = &RoadtoEnlightenment{}
