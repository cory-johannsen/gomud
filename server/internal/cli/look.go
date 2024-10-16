package cli

import (
	"context"
	"fmt"
)

type LookHandler struct {
	stateProvider StateProvider
}

func (h *LookHandler) Handle(ctx context.Context, args []string) (string, error) {
	player := h.stateProvider().Player()
	room := player.Room()
	msg := fmt.Sprintf("You are in %s\n\t%s\n\tExits:\n", room.Name, room.Description)
	for _, exit := range room.Exits() {
		msg += fmt.Sprintf("\t\t%s: %s\n", exit.Direction, exit.Description)
	}
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
