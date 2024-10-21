package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GallowsHumor struct {
  name string
  description string
}

func NewGallowsHumor() *GallowsHumor {
  return &GallowsHumor{
    name: "Gallows Humor",
    description: "At your option, you can substitute the Guile Skill in place of Resolve.",
  }
}

func (e *GallowsHumor) Name() string {
  return e.name
}

func (e *GallowsHumor) Description() string {
  return e.description
}

func (e *GallowsHumor) Applier() domain.Applier {
  return e.Apply
}

func (e *GallowsHumor) Apply(state domain.State) domain.State {
  // - At your option, you can substitute the Guile Skill in place of Resolve.
  log.Println("applying Gallows Humor")
  return state
}

var _ domain.Effect = &GallowsHumor{}
