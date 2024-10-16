package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type LookHandler struct {
	stateProvider StateProvider
}

func Look(player *domain.Player) string {
	room := player.Room()
	msg := fmt.Sprintf("You are in %s\n\t%s\n\tExits:\n", room.Name, room.Description)
	for _, exit := range room.Exits() {
		msg += fmt.Sprintf("\t\t%s: %s\n", exit.Direction, exit.Description)
	}
	for _, other := range room.Players {
		if other == player {
			continue
		}
		msg += fmt.Sprintf("%s the %s is here\n", other.Name, other.Job().Name)
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

func (h *LookHandler) State() State {
	return h.stateProvider()
}

func NewLookHandler(stateProvider StateProvider) *LookHandler {
	return &LookHandler{stateProvider: stateProvider}
}

var _ Handler = &LookHandler{}
