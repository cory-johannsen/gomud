package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/fatih/color"
	"strings"
)

type LookHandler struct {
	stateProvider domain.StateProvider
}

type LookTarget struct {
	Target string
	Player *domain.Player
	NPC    *domain.Character
	Room   *domain.Room
}

func Look(player *domain.Player, target *LookTarget) string {
	if target == nil {
		return LookRoom(player, player.Room())
	}
	if target.Room != nil {
		return LookRoom(player, target.Room)
	}
	if target.Player != nil {
		return LookCharacter(&target.Player.Character, false)
	}
	if target.NPC != nil {
		return LookCharacter(target.NPC, true)
	}
	return LookRoom(player, player.Room())
}

func LookRoom(player *domain.Player, target *domain.Room) string {
	msg := fmt.Sprintf("%s\n\t%s\n\tExits:\n", target.Name, target.Description)
	for _, exit := range target.Exits() {
		msg += fmt.Sprintf("\t\t%s: %s\n", exit.Direction, exit.Description)
	}
	msg += "\n"
	for _, other := range target.Players {
		if other == player {
			continue
		}
		msg += LookCharacter(&other.Character, false)
	}
	for _, other := range target.NPCs {
		msg += LookCharacter(other, true)
	}
	return msg
}

func LookCharacter(target *domain.Character, npc bool) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	var armor string
	if target.Inventory().Armor() != nil {
		armor = target.Inventory().Armor().Name()
	} else {
		armor = "nothin'"
	}
	var wields string
	if target.Inventory().MainHand() != nil {
		wields = fmt.Sprintf("%s in their main hand", target.Inventory().MainHand().Name())
		if target.Inventory().OffHand() != nil {
			wields = fmt.Sprintf("%s and %s in their off hand", wields, target.Inventory().OffHand().Name())
		}
	} else {
		wields = "their bare fists"
	}
	if target.Inventory().Shield() != nil {
		wields = fmt.Sprintf("%s. They are wearing a %s", target.Inventory().Shield().Name(), wields)
	}
	perilCondition := target.Peril().Condition
	var perilDescription string
	switch perilCondition {
	default:
		fallthrough
	case domain.PerilConditionUnhindered:
		perilDescription = green("cool as a cucumber.")
	case domain.PerilConditionImperiled:
		perilDescription = yellow("kinda concerned.")
	case domain.PerilConditionIgnore1SkillRank:
		perilDescription = yellow("pretty freaked out.")
	case domain.PerilConditionIgnore2SkillRanks:
		perilDescription = red("super freaked out.")
	case domain.PerilConditionIgnore3SkillRanks:
		perilDescription = red("freaked out of their damn mind!")
	case domain.PerilConditionIncapacitated:
		perilDescription = red("totally incapacitated by all this shit!")
	}
	name := cyan(target.Name)
	if npc {
		name = fmt.Sprintf("%s %s", cyan(target.Name), green("[NPC]"))
	}
	msg := fmt.Sprintf("%s the %s looks %s  They are wearing %s and wield %s.  \n  They are %s", name, green(target.Job().Name),
		perilDescription, magenta(armor), magenta(wields), yellow(target.Condition()))
	if len(target.Injuries()) == 0 {
		msg += " and unscathed."
	} else {
		msg += " and injured.  They have:\n"
		for _, injury := range target.Injuries() {
			msg += fmt.Sprintf("\t%s\n", red(injury))
		}
	}
	msg += "\n"
	return msg
}

func (h *LookHandler) Handle(ctx context.Context, args []string) (string, error) {
	var target *LookTarget
	if len(args) > 0 {
		target = &LookTarget{
			Target: strings.Join(args, " "),
			Player: nil,
			NPC:    nil,
			Room:   nil,
		}

		room := h.stateProvider().Player().Room()
		if room.Name == target.Target {
			target.Room = room
		} else {
			found := false
			for name, r := range room.Exits() {
				if name == target.Target {
					target.Room = r.Target
					found = true
					break
				}
			}
			if !found {
				for _, player := range room.Players {
					if player.Name == target.Target {
						target.Player = player
						break
					}
				}
			}
			if !found {
				for _, npc := range room.NPCs {
					if npc.Name == target.Target {
						target.NPC = npc
						break
					}
				}
			}
		}
	}

	msg := Look(h.stateProvider().Player(), target)
	return msg, nil
}

func (h *LookHandler) Help(args []string) string {
	return "look around.  Usage: look <target>"
}

func (h *LookHandler) State() domain.State {
	return h.stateProvider()
}

func NewLookHandler(stateProvider domain.StateProvider) *LookHandler {
	return &LookHandler{stateProvider: stateProvider}
}

var _ Handler = &LookHandler{}
