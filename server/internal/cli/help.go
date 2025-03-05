package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"sort"
	"strings"
)

type HelpHandler struct {
	stateProvider domain.StateProvider
	handlers      map[string]Handler
}

func (h *HelpHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		commands := make([]string, 0)
		for k := range h.handlers {
			commands = append(commands, k)
		}
		sort.Strings(commands)
		return "available commands: " + strings.Join(commands, ", "), nil
	}
	cmd := args[0]
	handler, ok := h.handlers[cmd]
	if !ok {
		return fmt.Sprintf("unrecognized command \"%s\"", cmd), nil
	}
	return handler.Help(args[1:]), nil
}

func (h *HelpHandler) Help(args []string) string {
	return "help [command]"
}

func (h *HelpHandler) State() domain.GameState {
	return h.stateProvider()
}

var _ Handler = &HelpHandler{}
