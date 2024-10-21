package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type SplitFace struct {
  name string
  description string
}

func NewSplitFace() *SplitFace {
  return &SplitFace{
    name: "Split Face",
    description: "You must flip the results to fail all Skill Tests which rely on smell and taste.",
  }
}

func (e *SplitFace) Name() string {
  return e.name
}

func (e *SplitFace) Description() string {
  return e.description
}

func (e *SplitFace) Applier() domain.Applier {
  return e.Apply
}

func (e *SplitFace) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Split Face")
  return state
}

var _ domain.Effect = &SplitFace{}
