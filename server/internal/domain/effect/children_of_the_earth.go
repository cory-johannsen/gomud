package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type ChildrenoftheEarth struct {
	name        string
	description string
}

func NewChildrenoftheEarth() *ChildrenoftheEarth {
	return &ChildrenoftheEarth{
		name:        "Children of the Earth",
		description: "You can never be forced off your feet or knocked Prone onto the ground by the elements, your enemies or even Magick. Finally, you will have a Corpulent build on the Build table.,",
	}
}

func (e *ChildrenoftheEarth) Name() string {
	return e.name
}

func (e *ChildrenoftheEarth) Description() string {
	return e.description
}

func (e *ChildrenoftheEarth) Applier() domain.Applier {
	return e.Apply
}

func (e *ChildrenoftheEarth) Apply(state domain.GameState) domain.GameState {
	// - You can never be forced off your feet or knocked Prone onto the ground by the elements, your enemies or even Magick. Finally, you will have a Corpulent build on the Build table.,
	log.Println("applying Children of the Earth")
	return state
}

var _ domain.Effect = &ChildrenoftheEarth{}
