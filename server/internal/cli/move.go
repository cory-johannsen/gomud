package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
)

type MoveHandler struct {
	stateProvider StateProvider
	players       *storage.Players
	rooms         *loader.RoomLoader
}

func NewMoveHandler(stateProvider StateProvider, players *storage.Players, rooms *loader.RoomLoader) *MoveHandler {
	return &MoveHandler{
		stateProvider: stateProvider,
		players:       players,
		rooms:         rooms,
	}
}

func (m *MoveHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "move where?", nil
	}
	player := m.stateProvider().Player()
	room := player.Room()
	direction := args[0]
	exits := room.Exits()
	exit, ok := exits[direction]
	if !ok {
		return "no exit in that direction", nil
	}
	destination := exit.Target
	player.SetRoom(destination)
	_, err := m.players.StorePlayer(ctx, player, player.Connection)
	if err != nil {
		log.Printf("error storing player: %s", err)
		return "", err
	}
	room.RemovePlayer(player)
	destination.AddPlayer(player)
	return Look(player), nil
}

func (m *MoveHandler) Help(args []string) string {
	return "move to another room.  Usage: move <exit>"
}

func (m *MoveHandler) State() State {
	return m.stateProvider()
}

var _ Handler = &MoveHandler{}

type DirectionalMoveHandler struct {
	MoveHandler
	Direction string
}

func NewDirectionalMoveHandler(stateProvider StateProvider, players *storage.Players, rooms *loader.RoomLoader, direction string) *DirectionalMoveHandler {
	return &DirectionalMoveHandler{
		MoveHandler: MoveHandler{
			stateProvider: stateProvider,
			players:       players,
			rooms:         rooms,
		},
		Direction: direction,
	}
}

func (d *DirectionalMoveHandler) Handle(ctx context.Context, args []string) (string, error) {
	return d.MoveHandler.Handle(ctx, []string{d.Direction})
}

func (d *DirectionalMoveHandler) Help(args []string) string {
	return fmt.Sprintf("move %s", d.Direction)
}

var _ Handler = &DirectionalMoveHandler{}
