package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type AlfredPfennigworth struct {
  name string
  description string
}

func NewAlfredPfennigworth() *AlfredPfennigworth {
  return &AlfredPfennigworth{
    name: "Alfred Pfennigworth",
    description: "When you Assist others with Skill Tests, they gain two Assist Dice instead, choosing the better of the two to use. In addition, when you Assist them, they may ignore any Critical Failures they roll, treating them as regular failures instead.",
  }
}

func (e *AlfredPfennigworth) Name() string {
  return e.name
}

func (e *AlfredPfennigworth) Description() string {
  return e.description
}

func (e *AlfredPfennigworth) Applier() domain.Applier {
  return e.Apply
}

func (e *AlfredPfennigworth) Apply(state domain.State) domain.State {
  // - When you Assist others with Skill Tests, they gain two Assist Dice instead, choosing the better of the two to use. In addition, when you Assist them, they may ignore any Critical Failures they roll, treating them as regular failures instead.
  log.Println("applying Alfred Pfennigworth")
  return state
}

var _ domain.Effect = &AlfredPfennigworth{}
