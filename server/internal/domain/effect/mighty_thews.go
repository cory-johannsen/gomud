package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type MightyThews struct {
	name        string
	description string
}

func NewMightyThews() *MightyThews {
	return &MightyThews{
		name:        "Mighty Thews",
		description: "You are able to wield any two-handed melee weapon with one hand. This also means you may freely take advantage of the Adaptable Quality for weapons using one hand. Whenever you gain the benefits of the Adaptable Quality, you may add +2 to Total Damage, instead of the normal +1.,",
	}
}

func (e *MightyThews) Name() string {
	return e.name
}

func (e *MightyThews) Description() string {
	return e.description
}

func (e *MightyThews) Applier() domain.Applier {
	return e.Apply
}

func (e *MightyThews) Apply(state domain.GameState) domain.GameState {
	// - You are able to wield any two-handed melee weapon with one hand. This also means you may freely take advantage of the Adaptable Quality for weapons using one hand. Whenever you gain the benefits of the Adaptable Quality, you may add +2 to Total Damage, instead of the normal +1.,
	log.Println("applying Mighty Thews")
	return state
}

var _ domain.Effect = &MightyThews{}
