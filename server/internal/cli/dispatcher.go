package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"strings"
)

type Dispatcher struct {
	handlers map[string]Handler
	ctx      context.Context
}

func NewDispatcher(stateConstructor StateConstructor, players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, conn Connection) *Dispatcher {
	dispatcher := &Dispatcher{
		handlers: make(map[string]Handler),
		ctx:      context.Background(),
	}
	quit := &QuitHandler{}
	quitAliases := CreateAliases(dispatcher.handlers["quit"], "exit", "q")
	for _, alias := range quitAliases {
		dispatcher.Register(alias.Alias, alias)
	}
	dispatcher.handlers["quit"] = quit
	helpAliases := CreateAliases(dispatcher.handlers["help"], "?", "h")
	for _, alias := range helpAliases {
		dispatcher.Register(alias.Alias, alias)
	}
	dispatcher.handlers["login"] = NewLoginHandler(stateConstructor, players, generator, teams, conn)
	return dispatcher
}

func (d *Dispatcher) Register(name string, handler Handler) {
	d.handlers[name] = handler
}

func (d *Dispatcher) Prompt() string {
	return "> "
}

func (d *Dispatcher) Tab(buffer string) string {
	return "" // do nothing
}

func (d *Dispatcher) Eval(buffer string) string {
	fields := strings.Fields(buffer)

	if len(fields) == 0 {
		return ""
	} else {
		cmd := fields[0]

		if cmd == "help" {
			if len(fields) == 1 {
				commands := make([]string, 0)
				for k := range d.handlers {
					commands = append(commands, k)
				}
				return "available commands: " + strings.Join(commands, ", ")
			}
			cmd = fields[1]
			handler, ok := d.handlers[cmd]
			if !ok {
				return fmt.Sprintf("unrecognized command \"%s\"", cmd)
			}
			return handler.Help(fields[2:])
		}

		handler, ok := d.handlers[cmd]
		if !ok {
			return fmt.Sprintf("unrecognized command \"%s\"", cmd)
		}

		result, err := handler.Handle(d.ctx, fields[1:])
		if err != nil {
			return fmt.Sprintf("error: %s", err)
		} else {
			return result
		}
	}
}
