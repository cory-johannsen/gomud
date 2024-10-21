package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type TurntheOtherCheek struct {
  name string
  description string
}

func NewTurntheOtherCheek() *TurntheOtherCheek {
  return &TurntheOtherCheek{
    name: "Turn the Other Cheek",
    description: "You can read, write and speak the language of any creature that is classified as a Humanoid (including player Ancestries). In addition, whenever you suffer Corruption, decrease the number you gain by 1 (to a minimum of one). This means that if you suffer 3 Corruption, you gain 2 instead.",
  }
}

func (e *TurntheOtherCheek) Name() string {
  return e.name
}

func (e *TurntheOtherCheek) Description() string {
  return e.description
}

func (e *TurntheOtherCheek) Applier() domain.Applier {
  return e.Apply
}

func (e *TurntheOtherCheek) Apply(state domain.State) domain.State {
  // - You can read, write and speak the language of any creature that is classified as a Humanoid (including player Ancestries). In addition, whenever you suffer Corruption, decrease the number you gain by 1 (to a minimum of one). This means that if you suffer 3 Corruption, you gain 2 instead.
  log.Println("applying Turn the Other Cheek")
  return state
}

var _ domain.Effect = &TurntheOtherCheek{}
