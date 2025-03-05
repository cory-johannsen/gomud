package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ThievingWays struct {
	name        string
	description string
}

func NewThievingWays() *ThievingWays {
	return &ThievingWays{
		name:        "Thieving Ways",
		description: "Foes must flip their results to fail all Attack Actions or Perilous Stunts made with ranged weapons against you. However, you cannot wield two-handed weapons, carry a shield and must reduce your Movement by 3. Finally, you will reference “1 to 10” on the Height table and will have a Frail build on the Build table.,",
	}
}

func (e *ThievingWays) Name() string {
	return e.name
}

func (e *ThievingWays) Description() string {
	return e.description
}

func (e *ThievingWays) Applier() domain.Applier {
	return e.Apply
}

func (e *ThievingWays) Apply(state domain.GameState) domain.GameState {
	// - Foes must flip their results to fail all Attack Actions or Perilous Stunts made with ranged weapons against you. However, you cannot wield two-handed weapons, carry a shield and must reduce your Movement by 3. Finally, you will reference “1 to 10” on the Height table and will have a Frail build on the Build table.,
	log.Println("applying Thieving Ways")
	return state
}

var _ domain.Effect = &ThievingWays{}
