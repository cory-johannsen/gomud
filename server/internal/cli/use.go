package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	UseItem   string = "item"
	UseSkill  string = "skill"
	UseObject string = "object"

	UseTargetSelf string = "self"
)

type UseHandler struct {
	stateProvider domain.StateProvider
	skills        *loader.SkillLoader
}

func NewUseHandler(stateProvider domain.StateProvider, skills *loader.SkillLoader) *UseHandler {
	return &UseHandler{
		stateProvider: stateProvider,
		skills:        skills,
	}
}

func (u *UseHandler) useSkill(skillName string, target *string) (string, error) {
	allSkills, err := u.skills.LoadSkills()
	if err != nil {
		return "", err
	}
	skills := u.stateProvider().Player().Skills(allSkills)
	skill := skills.Find(skillName)
	if skill == nil {
		return "skill not found", nil
	}
	roll := domain.D100()
	// todo: Add in difficulty scaling
	successTarget := skill.SuccessPercentage
	crit := roll%11 == 0
	cyan := color.New(color.FgCyan)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)
	msg := ""
	result := ""
	if roll >= successTarget {
		msg = cyan.Sprintf("%s", "success")
		result = cyan.Sprintf("%d", roll)
		if crit {
			msg = green.Sprintf("%s", "critical success!")
			result = green.Sprintf("%d", roll)
		}
		msg = "[" + result + "] " + msg + " (needed " + cyan.Sprintf("%d", successTarget) + ")"
	} else {
		msg = yellow.Sprintf("%s", "failure")
		result = yellow.Sprintf("%d", roll)
		if crit {
			msg = red.Sprintf("%s", "critical failure!")
			result = red.Sprintf("%d", roll)
		}
		msg = "[" + result + "] " + msg + " (needed " + yellow.Sprintf("%d", successTarget) + ")"
	}
	return fmt.Sprintf("%s: %s", skillName, msg), nil
}

func (u *UseHandler) useItem(itemName string, target *string) (string, error) {
	state := u.stateProvider()
	player := state.Player()
	if itemName == player.Inventory().MainHand().Name() {
		if target == nil {
			return "You can't use your main hand weapon on your self.", nil
		}
		return fmt.Sprintf("using main hand weapon %s on %s"+itemName, *target), nil
	}
	if itemName == player.Inventory().OffHand().Name() {
		if target == nil {
			return "You can't use your off-hand weapon on your self.", nil
		}
		return fmt.Sprintf("using off0hand weapon %s on %s"+itemName, *target), nil
	}
	for _, item := range player.Inventory().Pack().Items() {
		if itemName == item.Name() {
			if target == nil {
				return "using pack item %s on your self.", nil
			}
			return fmt.Sprintf("using pack item %s on %s"+itemName, *target), nil
		}
	}
	return fmt.Sprintf("You don't have a %s.", itemName), nil
}

func (u *UseHandler) useObject(objectName string, target *string) (string, error) {
	log.Printf("using object %s", objectName)
	state := u.stateProvider()
	player := state.Player()
	room := player.Room()
	object, ok := room.Objects[objectName]
	if !ok {
		return fmt.Sprintf("I don't see a %s here.", objectName), nil
	}
	result, err := object.Interact(&state, &player.Character, target)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (u *UseHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		return "use what? (skill/item/object)", nil
	}
	useType := strings.ToLower(args[0])
	if len(args) < 2 {
		if useType == strings.ToLower(UseSkill) {
			return "use which skill?", nil
		}
		if useType == strings.ToLower(UseItem) {
			return "use which item?", nil
		}
		if useType == strings.ToLower(UseItem) {
			return "use which object?", nil
		}
	}
	obj := strings.Join(args[1:], " ")

	u.stateProvider().Player().Connection.Writeln(fmt.Sprintf("use %s on what? (self, target name)", obj))
	response := u.stateProvider().Player().Connection.Read()
	var target *string
	if len(response) > 0 && response != UseTargetSelf {
		target = &response
	}

	switch useType {
	case UseSkill:
		return u.useSkill(obj, target)
	case UseItem:
		return u.useItem(obj, target)
	case UseObject:
		return u.useObject(obj, target)
	}
	return "unknown use type", nil
}

func (u *UseHandler) Help(args []string) string {
	return "use a skill, item, or object in the room.  Usage: use [skill|item|object] [name]"
}

func (u *UseHandler) State() domain.GameState {
	return u.stateProvider()
}

var _ Handler = &UseHandler{}
