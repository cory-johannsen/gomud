package cli

import (
	"github.com/openengineer/go-repl"
)

var helpMessage = `help              display this message
quit              quit this program`

type Handler interface {
	Handle(args []string) (string, error)
}

type HelpHandler struct{}

func (h *HelpHandler) Handle([]string) (string, error) {
	// todo: implement per-command help
	return helpMessage, nil
}

type QuitHandler struct {
	R *repl.Repl
}

func (h *QuitHandler) Handle([]string) (string, error) {
	h.R.Quit()
	return "peace out", nil
}
