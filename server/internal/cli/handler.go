package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type State interface {
	Player() *domain.Player
	Property(string) domain.Property
}

type StateConstructor func(player *domain.Player) State

type Handler interface {
	Handle(ctx context.Context, args []string) (string, error)
	Help(args []string) string
	State() State
}

type Alias struct {
	Alias   string
	Handler Handler
}

func (a *Alias) Handle(ctx context.Context, args []string) (string, error) {
	return a.Handler.Handle(ctx, args)
}

func (a *Alias) Help(args []string) string {
	return a.Handler.Help(args)
}

func (a *Alias) State() State {
	return a.Handler.State()
}

type Aliases []*Alias

func CreateAliases(handler Handler, aliases ...string) Aliases {
	var a Aliases
	for _, alias := range aliases {
		a = append(a, &Alias{Alias: alias, Handler: handler})
	}
	return a
}

var _ Handler = &Alias{}

const WelcomeMessage = "\n<-- 🔫 Gunchete 🔪 -->\n\nWelcome to Gunchete!  Type 'help' for a list of commands.\n"
