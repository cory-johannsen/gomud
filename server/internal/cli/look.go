package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type LookHandler struct {
	stateProvider domain.StateProvider
}

func Look(player *domain.Player) string {
	room := player.Room()
	msg := fmt.Sprintf("You are in %s\n\t%s\n\tExits:\n", room.Name, room.Description)
	for _, exit := range room.Exits() {
		msg += fmt.Sprintf("\t\t%s: %s\n", exit.Direction, exit.Description)
	}
	msg += "\n"
	for _, other := range room.Players {
		if other == player {
			continue
		}
		perilCondition := other.Peril().Condition
		var perilDescription string
		switch perilCondition {
		default:
			fallthrough
		case domain.PerilConditionUnhindered:
			perilDescription = "cool as a cucumber."
		case domain.PerilConditionImperiled:
			perilDescription = "kinda concerned."
		case domain.PerilConditionIgnore1SkillRank:
			perilDescription = "pretty freaked out."
		case domain.PerilConditionIgnore2SkillRanks:
			perilDescription = "super freaked out."
		case domain.PerilConditionIgnore3SkillRanks:
			perilDescription = "freaked out of their damn mind!"
		case domain.PerilConditionIncapacitated:
			perilDescription = "totally incapacitated by all this shit!"
		}
		msg += fmt.Sprintf("%s the %s is here.  They look %s\n", other.Name, other.Job().Name, perilDescription)
	}
	return msg
}

func (h *LookHandler) Handle(ctx context.Context, args []string) (string, error) {
	msg := Look(h.stateProvider().Player())
	return msg, nil
}

func (h *LookHandler) Help(args []string) string {
	return "look around"
}

func (h *LookHandler) State() domain.State {
	return h.stateProvider()
}

func NewLookHandler(stateProvider domain.StateProvider) *LookHandler {
	return &LookHandler{stateProvider: stateProvider}
}

var _ Handler = &LookHandler{}
