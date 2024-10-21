package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DebtRidden struct {
  name string
  description string
}

func NewDebtRidden() *DebtRidden {
  return &DebtRidden{
    name: "Debt-Ridden",
    description: "You must flip the results to fail all Skill Tests that rely on your ability to barter, bargain or strike monetary deals in your favor.",
  }
}

func (e *DebtRidden) Name() string {
  return e.name
}

func (e *DebtRidden) Description() string {
  return e.description
}

func (e *DebtRidden) Applier() domain.Applier {
  return e.Apply
}

func (e *DebtRidden) Apply(state domain.State) domain.State {
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Debt-Ridden")
  return state
}

var _ domain.Effect = &DebtRidden{}
