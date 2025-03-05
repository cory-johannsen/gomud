package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type BlessinginDisguise struct {
	name        string
	description string
}

func NewBlessinginDisguise() *BlessinginDisguise {
	return &BlessinginDisguise{
		name:        "Blessing in Disguise",
		description: "Whenever the time calls for you to suffer a debilitating Injury or face death, you may alternatively gain 1 permanent Chaos Rank (instead of spending a Fate Point) to ignore it entirely.,",
	}
}

func (e *BlessinginDisguise) Name() string {
	return e.name
}

func (e *BlessinginDisguise) Description() string {
	return e.description
}

func (e *BlessinginDisguise) Applier() domain.Applier {
	return e.Apply
}

func (e *BlessinginDisguise) Apply(state domain.GameState) domain.GameState {
	// - Whenever the time calls for you to suffer a debilitating Injury or face death, you may alternatively gain 1 permanent Chaos Rank (instead of spending a Fate Point) to ignore it entirely.,
	log.Println("applying Blessing in Disguise")
	return state
}

var _ domain.Effect = &BlessinginDisguise{}
