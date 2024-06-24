package cli

import (
	"fmt"
	"github.com/openengineer/go-repl"
	"strings"
)

var helpMessage = `help              display this message
quit              quit this program`

type CommandHandler struct {
	R *repl.Repl
}

func (h *CommandHandler) Prompt() string {
	return "> "
}

func (h *CommandHandler) Tab(buffer string) string {
	return "" // do nothing
}

func (h *CommandHandler) Eval(line string) string {
	fields := strings.Fields(line)

	if len(fields) == 0 {
		return ""
	} else {
		cmd := fields[0]

		switch cmd {
		case "help":
			return helpMessage
		case "quit":
			h.R.Quit()
			return ""
		default:
			return fmt.Sprintf("unrecognized command \"%s\"", cmd)
		}
	}
}
