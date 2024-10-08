package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"strings"
)

type StateProvider func() State

type Dispatcher struct {
	handlers map[string]Handler
	ctx      context.Context
	state    State
}

func (d *Dispatcher) State() State {
	return d.state
}

func NewDispatcher(stateConstructor StateConstructor, players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, conn Connection) *Dispatcher {
	dispatcher := &Dispatcher{
		handlers: make(map[string]Handler),
		ctx:      context.Background(),
	}
	quit := &QuitHandler{}
	quitAliases := CreateAliases(dispatcher.handlers["quit"], "exit", "q")
	dispatcher.Register("quit", quit)
	for _, alias := range quitAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	helpAliases := CreateAliases(dispatcher.handlers["help"], "?", "h")
	for _, alias := range helpAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	dispatcher.Register("login", NewLoginHandler(stateConstructor, players, generator, teams, conn))
	
	characterHandler := CharacterHandler{stateProvider: dispatcher.State}
	dispatcher.Register("character", &characterHandler)
	characterAliases := CreateAliases(&characterHandler, "c", "char", "me", "self")
	for _, alias := range characterAliases {
		dispatcher.Register(alias.Alias, alias)
	}
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
		}
		state := handler.State()
		if state != nil {
			d.state = state
		}
		return result
	}
}
