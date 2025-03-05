package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/fatih/color"
)

type Handler interface {
	Handle(ctx context.Context, args []string) (string, error)
	Help(args []string) string
	State() domain.GameState
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

func (a *Alias) State() domain.GameState {
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

const WelcomeMessage = "Welcome to Gunchete!  Type 'help' for a list of commands."

func Welcome() string {
	red := color.New(color.FgRed).SprintFunc()
	orange := color.New(color.FgHiRed).Add(color.FgHiYellow).Add(color.Underline).SprintFunc()
	return fmt.Sprintf("\n%s 🔫 %s 🔪 %s\n\n%s\n", red("<--"), orange("Gunchete"), red("-->"), WelcomeMessage)
}
