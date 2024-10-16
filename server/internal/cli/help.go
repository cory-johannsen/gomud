package cli

import (
	"context"
	"fmt"
	"sort"
	"strings"
)

type HelpHandler struct {
	stateProvider StateProvider
	dispatcher    *Dispatcher
}

func (h *HelpHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		commands := make([]string, 0)
		for k := range h.dispatcher.handlers {
			commands = append(commands, k)
		}
		sort.Strings(commands)
		return "available commands: " + strings.Join(commands, ", "), nil
	}
	cmd := args[0]
	handler, ok := h.dispatcher.handlers[cmd]
	if !ok {
		return fmt.Sprintf("unrecognized command \"%s\"", cmd), nil
	}
	return handler.Help(args[1:]), nil
}

func (h *HelpHandler) Help(args []string) string {
	return "help [command]"
}

func (h *HelpHandler) State() State {
	return h.stateProvider()
}

var _ Handler = &HelpHandler{}
