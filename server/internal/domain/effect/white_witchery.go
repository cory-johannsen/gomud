package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type WhiteWitchery struct {
  name string
  description string
}

func NewWhiteWitchery() *WhiteWitchery {
  return &WhiteWitchery{
    name: "White Witchery",
    description: "Select any one Generalist Magick spell (providing that it does not cause harm to another). When you attempt to cast this Magick, you do not need to roll percentile dice to make a Skill Test. Instead, you Critically Succeed at the Incantation Test you intended to make. Furthermore, you may select one additional Generalist Magick spell at both Intermediate and Advanced Tier with the aforementioned limitations.",
  }
}

func (e *WhiteWitchery) Name() string {
  return e.name
}

func (e *WhiteWitchery) Description() string {
  return e.description
}

func (e *WhiteWitchery) Applier() domain.Applier {
  return e.Apply
}

func (e *WhiteWitchery) Apply(state domain.State) domain.State {
  // - Select any one Generalist Magick spell (providing that it does not cause harm to another). When you attempt to cast this Magick, you do not need to roll percentile dice to make a Skill Test. Instead, you Critically Succeed at the Incantation Test you intended to make. Furthermore, you may select one additional Generalist Magick spell at both Intermediate and Advanced Tier with the aforementioned limitations.
  log.Println("applying White Witchery")
  return state
}

var _ domain.Effect = &WhiteWitchery{}
