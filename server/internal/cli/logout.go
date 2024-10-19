package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
)

const QuitMessage = "peace out"

type LogoutHandler struct {
	stateProvider StateProvider
	players       *storage.Players
}

func (l *LogoutHandler) Handle(ctx context.Context, args []string) (string, error) {
	state := l.stateProvider()
	player := state.Player()
	if !player.LoggedIn {
		return "you're not logged in", nil
	}
	room := player.Room()
	room.RemovePlayer(player)
	_, err := l.players.StorePlayer(ctx, player, player.Connection)
	if err != nil {
		log.Printf("error storing player: %s", err)
	}
	player.LoggedIn = false
	return QuitMessage + " " + player.Name, nil
}

func (l *LogoutHandler) Help(args []string) string {
	return "abandon your dawgs to the streets"
}

func (l *LogoutHandler) State() State {
	return l.stateProvider()
}

func NewLogoutHandler(stateProvider StateProvider, players *storage.Players) *LogoutHandler {
	return &LogoutHandler{
		stateProvider: stateProvider,
		players:       players,
	}
}

var _ Handler = &LogoutHandler{}
