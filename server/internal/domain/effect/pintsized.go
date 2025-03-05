package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type Pintsized struct {
	name        string
	description string
}

func NewPintsized() *Pintsized {
	return &Pintsized{
		name:        "Pintsized",
		description: "Foes must flip their results to fail all Attack Actions or Perilous Stunts made with ranged weapons against you. However, you cannot wield two-handed weapons, carry a shield and must reduce your Movement by 3. Finally, you will reference “1 to 10” on the Height table and will have a Frail build on the Build table.,",
	}
}

func (e *Pintsized) Name() string {
	return e.name
}

func (e *Pintsized) Description() string {
	return e.description
}

func (e *Pintsized) Applier() domain.Applier {
	return e.Apply
}

func (e *Pintsized) Apply(state domain.GameState) domain.GameState {
	// - Foes must flip their results to fail all Attack Actions or Perilous Stunts made with ranged weapons against you. However, you cannot wield two-handed weapons, carry a shield and must reduce your Movement by 3. Finally, you will reference “1 to 10” on the Height table and will have a Frail build on the Build table.,
	log.Println("applying Pintsized")
	return state
}

var _ domain.Effect = &Pintsized{}
