package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

type LoginHandler struct {
	state            State
	stateConstructor StateConstructor
	players          *storage.Players
	generator        *generator.PlayerGenerator
	teams            *loader.TeamLoader
	rooms            *loader.RoomLoader
	conn             Connection
}

func NewLoginHandler(stateConstructor StateConstructor, players *storage.Players, generator *generator.PlayerGenerator,
	teams *loader.TeamLoader, rooms *loader.RoomLoader, conn Connection) *LoginHandler {
	return &LoginHandler{
		stateConstructor: stateConstructor,
		players:          players,
		generator:        generator,
		teams:            teams,
		rooms:            rooms,
		conn:             conn,
	}
}

func (h *LoginHandler) Handle(ctx context.Context, args []string) (string, error) {
	var name string
	if len(args) != 1 {
		read, err := h.readName()
		if err != nil {
			return "", err
		}
		name = read
	} else {
		name = args[0]
	}
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

	player.Room().AddPlayer(player)
	msg := fmt.Sprintf("Welcome %s!\n%s", name, Look(player))
	return msg, nil
}

func (h *LoginHandler) createPlayer(name string) (*domain.Player, error) {
	msg := fmt.Sprintf("creating new player %s", name)
	_ = h.conn.Writeln(msg)

	team, err := h.selectTeam()
	if err != nil {
		return nil, err
	}
	pw, err := h.enterPassword()

	takeDrawback := h.takeDrawback()

	player, err := h.generator.Generate(name, pw, team, takeDrawback)
	if err != nil {
		return nil, err
	}

	_ = h.conn.Writeln(fmt.Sprintf("Created player %s", player.String()))
	_ = h.conn.Writeln(fmt.Sprintf("You now get to select your skills, bonuses and talents.  You have %d experience to spend, each selection cosets 100 experience.", player.Experience()))
	for {
		if player.Experience() < 100 {
			break
		}
		err = h.selectSkillRanks(player)
		if err != nil {
			return nil, err
		}

		err = h.selectBonusAdvances(player)
		if err != nil {
			return nil, err
		}

		err = h.selectTalents(player)
		if err != nil {
			return nil, err
		}

		if player.Experience() >= 100 {
			_ = h.conn.Write("You have more experience to spend.  Do you wish to continue spending? (y/n): ")
			choice := h.conn.Read()
			if choice == "n" {
				break
			}
		}
	}

	room := h.rooms.GetRoom("Wayne Dawg's Trailer")
	player.SetRoom(room)
	room.AddPlayer(player)

	log.Printf("Created player %s", player.Name)

	return player, nil
}

func (h *LoginHandler) readName() (string, error) {
	_ = h.conn.Write("Who dis?: ")
	var name string
	for {
		name = h.conn.Read()
		if name != "" {
			break
		}
		_ = h.conn.Write("That ain't a name.  I said who dis?: ")
	}
	return name, nil
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

func (h *LoginHandler) selectTeam() (*domain.Team, error) {
	var t *domain.Team
	for {
		msg := fmt.Sprintf("Select Team:\n")
		teams, err := h.teams.LoadTeams()
		if err != nil {
			return nil, err
		}
		for i, team := range teams {
			msg += fmt.Sprintf("%d) %s\n", i, team.Name)
		}
		_ = h.conn.Write(msg + "> ")

		team := h.conn.Read()

		index, err := strconv.Atoi(team)
		if err != nil {
			_ = h.conn.Write("Invalid team.  Please select a valid team\n> ")
			continue
		}
		if index < 0 || index >= len(teams) {
			_ = h.conn.Write("Invalid team.  Please select a valid team\n> ")
			continue
		}
		t = teams[index]
		if t != nil {
			break
		}
		_ = h.conn.Write("Invalid team.  Please select a valid team\n> ")
	}
	return t, nil
}

func (h *LoginHandler) takeDrawback() bool {
	_ = h.conn.Write("Do you want to take a drawback? (y/n): ")
	var takeDrawback bool
	for {
		drawback := h.conn.Read()
		if drawback == "y" {
			takeDrawback = true
			break
		}
		if drawback == "n" {
			takeDrawback = false
			break
		}
		_ = h.conn.Write("Invalid response.  Please enter 'y' or 'n'\n> ")
	}
	return takeDrawback
}

func (h *LoginHandler) selectSkillRanks(player *domain.Player) error {
	for {
		_ = h.conn.Write(fmt.Sprintf("You have %d experience to spend. Purchase a Skill Rank (100 exp each):\n", player.Experience()))
		purchased := make(map[int]bool)
		for i, skillRank := range player.Job().SkillRanks {
			if !player.HasSkillRank(player.Job(), skillRank) {
				_ = h.conn.Write(fmt.Sprintf("%d) %s\n", i, skillRank.Name))
			} else {
				purchased[i] = true
				_ = h.conn.Write(fmt.Sprintf("%d) %s (purchased)\n", i, skillRank.Name))
			}
		}
		_ = h.conn.Write("Q) Quit purchasing Skill Ranks\n> ")
		choice := h.conn.Read()
		if strings.ToUpper(choice) == "Q" {
			break
		}
		index, err := strconv.Atoi(choice)
		if err != nil {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		if index < 0 || index >= len(player.Job().SkillRanks) {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		if _, found := purchased[index]; found {
			_ = h.conn.Writeln("Skill rank already purchased")
			continue
		}
		skill := player.Job().SkillRanks[index]
		player.PurchaseSkillRank(player.Job(), skill, 100)
		if player.Experience() < 100 {
			break
		}
	}

	return nil
}

func (h *LoginHandler) selectBonusAdvances(player *domain.Player) error {
	for {
		if player.Experience() < 100 {
			break
		}
		_ = h.conn.Write(fmt.Sprintf("You have %d experience to spend. Purchase an Advance:\n", player.Experience()))
		bonusAdvances := player.Job().BonusAdvances
		consumedAdvances := player.ConsumedBonusAdvances()
		i := 0
		choices := make([]string, 0)
		if bonusAdvances.Fighting > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Fighting") < bonusAdvances.Fighting {
				advancesLeft := bonusAdvances.Fighting - consumedAdvances.ConsumedAdvance(player.Job().Name, "Fighting")
				_ = h.conn.Write(fmt.Sprintf("%d) Fighting (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Fighting")
			}
		}
		if bonusAdvances.Muscle > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Muscle") < bonusAdvances.Muscle {
				advancesLeft := bonusAdvances.Muscle - consumedAdvances.ConsumedAdvance(player.Job().Name, "Muscle")
				_ = h.conn.Write(fmt.Sprintf("%d) Muscle (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Muscle")
			}
		}
		if bonusAdvances.Speed > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Speed") < bonusAdvances.Speed {
				advancesLeft := bonusAdvances.Speed - consumedAdvances.ConsumedAdvance(player.Job().Name, "Speed")
				_ = h.conn.Write(fmt.Sprintf("%d) Speed (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Speed")
			}
		}
		if bonusAdvances.Savvy > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Savvy") < bonusAdvances.Savvy {
				advancesLeft := bonusAdvances.Savvy - consumedAdvances.ConsumedAdvance(player.Job().Name, "Savvy")
				_ = h.conn.Write(fmt.Sprintf("%d) Savvy (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Savvy")
			}
		}
		if bonusAdvances.Smarts > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Smarts") < bonusAdvances.Smarts {
				advancesLeft := bonusAdvances.Smarts - consumedAdvances.ConsumedAdvance(player.Job().Name, "Smarts")
				_ = h.conn.Write(fmt.Sprintf("%d) Smarts (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Smarts")
			}
		}
		if bonusAdvances.Grit > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Grit") < bonusAdvances.Grit {
				advancesLeft := bonusAdvances.Grit - consumedAdvances.ConsumedAdvance(player.Job().Name, "Grit")
				_ = h.conn.Write(fmt.Sprintf("%d) Grit (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Grit")
			}
		}
		if bonusAdvances.Flair > 0 {
			if consumedAdvances.ConsumedAdvance(player.Job().Name, "Flair") < bonusAdvances.Flair {
				advancesLeft := bonusAdvances.Flair - consumedAdvances.ConsumedAdvance(player.Job().Name, "Flair")
				_ = h.conn.Write(fmt.Sprintf("%d) Flair (%d remaining)\n", i, advancesLeft))
				i++
				choices = append(choices, "Flair")
			}
		}
		_ = h.conn.Write("Q) Quit purchasing Advances\n> ")
		choice := h.conn.Read()
		if strings.ToUpper(choice) == "Q" {
			break
		}
		index, err := strconv.Atoi(choice)
		if err != nil {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		if index < 0 || index >= len(choices) {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		advance := choices[index]
		player.ConsumeBonusAdvance(player.Job().Name, advance, 100)
	}

	return nil
}

func (h *LoginHandler) selectTalents(player *domain.Player) error {
	for {
		if player.Experience() < 100 {
			break
		}
		_ = h.conn.Write(fmt.Sprintf("You have %d experience to spend. Purchase a Talent:\n", player.Experience()))
		purchased := make(map[int]bool)
		for i, talent := range player.Job().Talents {
			if !player.HasTalent(player.Job(), talent) {
				_ = h.conn.Write(fmt.Sprintf("%d) %s\n", i, talent.Name))
			} else {
				purchased[i] = true
				_ = h.conn.Write(fmt.Sprintf("%d) %s (purchased)\n", i, talent.Name))
			}
		}
		_ = h.conn.Write("Q) Quit purchasing Talents\n> ")
		choice := h.conn.Read()
		if strings.ToUpper(choice) == "Q" {
			break
		}
		index, err := strconv.Atoi(choice)
		if err != nil {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		if index < 0 || index >= len(player.Job().Talents) {
			_ = h.conn.Writeln("Invalid choice")
			continue
		}
		if _, found := purchased[index]; found {
			_ = h.conn.Writeln("Talent already purchased")
			continue
		}
		talent := player.Job().Talents[index]
		player.ConsumeTalent(player.Job(), talent, 100)
	}
	return nil
}

func (h *LoginHandler) Help(args []string) string {
	return "login to the system.  Usage: login <username>"
}

func (h *LoginHandler) State() State {
	return h.state
}

var _ Handler = &LoginHandler{}
