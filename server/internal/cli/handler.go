package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type Property interface {
	Value() interface{}
	String() string
}

type State interface {
	Player() *domain.Player
	Property(string) Property
}

type StateConstructor func(player *domain.Player) State

type Connection interface {
	Read() string
	Write(string) int
	Writeln(string) int
}

type Handler interface {
	Handle(ctx context.Context, args []string) (string, error)
	Help(args []string) string
	State() State
}

const WelcomeMessage = "-- Gunchete -->\n\nWelcome to Gunchete!  Type 'help' for a list of commands.\n"
const QuitMessage = "peace out"

type QuitHandler struct {
}

func (h *QuitHandler) Handle(ctx context.Context, args []string) (string, error) {
	return QuitMessage, nil
}

func (h *QuitHandler) Help([]string) string {
	return "abandon your dawgs to the streets"
}

func (h *QuitHandler) State() State {
	return nil
}
