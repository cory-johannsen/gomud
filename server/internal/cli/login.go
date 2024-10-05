package cli

import (
	"context"
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type LoginHandler struct {
	state            State
	stateConstructor StateConstructor
	players          *storage.Players
	generator        *generator.PlayerGenerator
	teams            *loader.TeamLoader
	conn             Connection
}

func NewLoginHandler(stateConstructor StateConstructor, players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, conn Connection) *LoginHandler {
	return &LoginHandler{
		stateConstructor: stateConstructor,
		players:          players,
		generator:        generator,
		teams:            teams,
		conn:             conn,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) != 1 {
		return "", errors.New("usage: login <username>")
	}
	name := args[0]
	exists, err := h.players.Exists(ctx, name)
	if err != nil {
		return "", err
	}
	var player *domain.Player
	if !exists {
		player, err = h.createPlayer(name)
		if err != nil {
			return "failed to create player", err
		}
		log.Printf("Storing player %s", name)
		player, err = h.players.StorePlayer(ctx, player)
		if err != nil {
			return "failed to store player", err
		}
		_ = h.conn.Writeln(fmt.Sprintf("Created player \n%s", player.String()))
	} else {
		player = h.validatePassword(name)
	}
	h.state = h.stateConstructor(player)
	return fmt.Sprintf("Welcome %s", name), nil
}

func (h *LoginHandler) createPlayer(name string) (*domain.Player, error) {
	msg := fmt.Sprintf("creating new player %s", name)
	_ = h.conn.Writeln(msg)

	msg = fmt.Sprintf("Select Team:\n")
	teams, err := h.teams.LoadTeams()
	if err != nil {
		return nil, err
	}
	for _, team := range teams {
		msg += fmt.Sprintf("%s\n", team.Name)
	}
	_ = h.conn.Write(msg + "> ")
	team, err := h.selectTeam(teams)
	if err != nil {
		return nil, err
	}
	pw, err := h.enterPassword()

	player, err := h.generator.Generate(name, pw, team)
	if err != nil {
		return nil, err
	}

	log.Printf("Created player %s", name)

	return player, nil
}

func (h *LoginHandler) enterPassword() (string, error) {
	_ = h.conn.Write("Enter Password: ")
	pw := h.conn.Read()
	_ = h.conn.Write("Confirm Password: ")

	confirm := h.conn.Read()
	if pw != confirm {
		_ = h.conn.Writeln("Passwords do not match.  Please try again")
		return h.enterPassword()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (h *LoginHandler) validatePassword(name string) *domain.Player {
	retries := 3
	player, err := h.players.FetchPlayerByName(context.Background(), name)
	if err != nil {
		_ = h.conn.Writeln(err.Error())
		return nil
	}
	for {
		_ = h.conn.Write("Enter Password: ")
		pw := h.conn.Read()
		if player.ValidPassword(pw) {
			break
		}
		_ = h.conn.Writeln("Invalid password")
		retries--
		if retries == 0 {
			_ = h.conn.Writeln("Too many retries.  Exiting")
			return nil
		}
	}
	return player
}

func (h *LoginHandler) selectTeam(teams domain.Teams) (*domain.Team, error) {
	var t *domain.Team
	for {
		team := h.conn.Read()
		for _, tm := range teams {
			if tm.Name == team {
				t = tm
				break
			}
		}
		if t != nil {
			break
		}
		_ = h.conn.Write("Invalid team.  Please select a valid team\n> ")
	}
	return t, nil
}

func (h *LoginHandler) Help(args []string) string {
	return "login to the system.  Usage: login <username>"
}

func (h *LoginHandler) State() State {
	return h.state
}

var _ Handler = &LoginHandler{}
