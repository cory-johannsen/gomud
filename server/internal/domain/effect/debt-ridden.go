package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
)

type DebtRidden struct {
	name        string
	description string
}

func NewDebtRidden() *DebtRidden {
	return &DebtRidden{
		name:        "Debt-Ridden",
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

func (e *DebtRidden) Apply(state domain.GameState) domain.GameState {
	// - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
	log.Println("applying Debt-Ridden")
	return state
}

var _ domain.Effect = &DebtRidden{}
