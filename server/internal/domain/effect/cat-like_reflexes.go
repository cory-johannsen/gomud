package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type CatlikeReflexes struct {
	name        string
	description string
}

func NewCatlikeReflexes() *CatlikeReflexes {
	return &CatlikeReflexes{
		name:        "Cat-like Reflexes",
		description: "Whenever you fall, you may spend one Fortune Point to avoid Damage. In this case, you land on your feet, apparently unharmed.,",
	}
}

func (e *CatlikeReflexes) Name() string {
	return e.name
}

func (e *CatlikeReflexes) Description() string {
	return e.description
}

func (e *CatlikeReflexes) Applier() domain.Applier {
	return e.Apply
}

func (e *CatlikeReflexes) Apply(state domain.GameState) domain.GameState {
	// - Whenever you fall, you may spend one Fortune Point to avoid Damage. In this case, you land on your feet, apparently unharmed.,
	log.Println("applying Cat-like Reflexes")
	return state
}

var _ domain.Effect = &CatlikeReflexes{}
