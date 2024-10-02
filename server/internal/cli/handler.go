package cli

type Handler interface {
	Handle(args []string) (string, error)
	Help(args []string) string
}

const QuitMessage = "peace out"

type QuitHandler struct {
}

func (h *QuitHandler) Handle([]string) (string, error) {
	return QuitMessage, nil
}

func (h *QuitHandler) Help([]string) string {
	return "abandon your dawgs to the streets"
}
