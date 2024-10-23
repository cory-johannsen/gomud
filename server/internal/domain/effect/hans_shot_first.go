package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type HansShotFirst struct {
  name string
  description string
}

func NewHansShotFirst() *HansShotFirst {
  return &HansShotFirst{
    name: "Hans Shot First",
    description: "When combat begins, you gain 3 APs that must be used immediately at the top of the Initiative Order – even if you were Surprised. Once spent, determine your place in the Initiative Ladder and take your Turns normally. If more than two Smugglers are present, the one with the highest [SB] goes first.",
  }
}

func (e *HansShotFirst) Name() string {
  return e.name
}

func (e *HansShotFirst) Description() string {
  return e.description
}

func (e *HansShotFirst) Applier() domain.Applier {
  return e.Apply
}

func (e *HansShotFirst) Apply(state domain.State) domain.State {
  // - When combat begins, you gain 3 APs that must be used immediately at the top of the Initiative Order – even if you were Surprised. Once spent, determine your place in the Initiative Ladder and take your Turns normally. If more than two Smugglers are present, the one with the highest [SB] goes first.
  log.Println("applying Hans Shot First")
  return state
}

var _ domain.Effect = &HansShotFirst{}
