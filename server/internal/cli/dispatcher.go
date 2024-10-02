package cli

import (
	"fmt"
	"github.com/openengineer/go-repl"
	"strings"
)

type Dispatcher struct {
	Repl     *repl.Repl
	handlers map[string]Handler
}

func NewDispatcher() *Dispatcher {
	dispatcher := &Dispatcher{
		handlers: map[string]Handler{
			"help": &HelpHandler{},
		},
	}
	dispatcher.Repl = repl.NewRepl(dispatcher)
	dispatcher.handlers["quit"] = &QuitHandler{R: dispatcher.Repl}
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

		handler, ok := d.handlers[cmd]
		if !ok {
			return fmt.Sprintf("unrecognized command \"%s\"", cmd)
		}

		result, err := handler.Handle(fields[1:])
		if err != nil {
			return fmt.Sprintf("error: %s", err)
		} else {
			return result
		}
	}
}
