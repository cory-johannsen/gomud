package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"sort"
	"strconv"
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
	result, err := object.Interact(state, &player.Character, target)
	if err != nil {
		return "", err
	}
	return result, nil
}

func selectUseType(ctx context.Context, state domain.StateProvider) string {
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	msg := fmt.Sprintf("%s: %s\n%s: %s\n%s: %s\n", green("1. skill"), "Use a skill", cyan("2. item"), "Use an item", yellow("3. object"), "Use an object")
	player := state().Player()
	player.Connection.Writeln(msg)
	response := state().Player().Connection.Read()
	if len(response) == 0 {
		return selectUseType(ctx, state)
	}
	switch response {
	case "1", "skill":
		response = UseSkill
	case "2", "item":
		response = UseItem
	case "3", "object":
		response = UseObject
	default:
		player.Connection.Writeln("invalid selection")
		return selectUseType(ctx, state)
	}
	return response
}

func selectSkill(skills domain.Skills, state domain.StateProvider) string {
	green := color.New(color.FgGreen).SprintFunc()
	player := state().Player()
	player.Connection.Writeln("use which skill?")
	playerSkills := player.Skills(skills)
	sortedSkills := make([]string, 0)
	for _, s := range playerSkills {
		if len(s.Skill.Name) > 0 {
			sortedSkills = append(sortedSkills, s.Skill.Name)
		}
	}
	sort.Strings(sortedSkills)
	for index, skill := range sortedSkills {
		msg := fmt.Sprintf("%s: %s", green(index+1), green(skill))
		player.Connection.Writeln(msg)
	}
	response := state().Player().Connection.Read()
	if len(response) == 0 {
		return selectSkill(skills, state)
	}
	index, err := strconv.Atoi(response)
	if err != nil {
		return response
	}
	return sortedSkills[index-1]
}

func selectItem(state domain.StateProvider) string {
	cyan := color.New(color.FgCyan).SprintFunc()
	player := state().Player()
	player.Connection.Writeln("use which item?")
	inventory := player.Inventory()
	items := make([]string, 0)
	if inventory.MainHand() != nil {
		items = append(items, inventory.MainHand().Name())
	}
	if inventory.OffHand() != nil {
		items = append(items, inventory.OffHand().Name())
	}
	for _, item := range inventory.Pack().Items() {
		items = append(items, item.Name())
	}
	for index, item := range items {
		player.Connection.Writeln(fmt.Sprintf("%s: %s", cyan(index+1), cyan(item)))
	}
	response := state().Player().Connection.Read()
	if len(response) == 0 {
		return selectItem(state)
	}
	index, err := strconv.Atoi(response)
	if err != nil {
		return response
	}
	return items[index-1]
}

func selectObject(state domain.StateProvider) string {
	yellow := color.New(color.FgYellow).SprintFunc()
	player := state().Player()
	player.Connection.Writeln("use which object?")
	objects := make([]string, 0)
	for _, o := range player.Room().Objects {
		objects = append(objects, o.Name())
	}
	for index, obj := range objects {
		player.Connection.Writeln(fmt.Sprintf("%s: %s", yellow(index+1), yellow(obj)))
	}
	response := state().Player().Connection.Read()
	index, err := strconv.Atoi(response)
	if err != nil {
		return response
	}
	return objects[index-1]
}

func selectTarget(obj string, state domain.StateProvider) *string {
	player := state().Player()
	player.Connection.Writeln(fmt.Sprintf("use %s on what? (self, target name)", obj))
	response := player.Connection.Read()
	var target *string
	if len(response) > 0 && response != UseTargetSelf {
		target = &response
	}
	return target
}

func (u *UseHandler) Handle(ctx context.Context, args []string) (string, error) {
	var useType string
	var obj string
	if len(args) == 0 {
		u.stateProvider().Player().Connection.Writeln("use what? (skill/item/object)")
		useType = selectUseType(ctx, u.stateProvider)
	} else {
		useType = strings.ToLower(args[0])
	}
	if len(args) < 2 {
		if useType == strings.ToLower(UseSkill) {
			skills, err := u.skills.LoadSkills()
			if err != nil {
				return "", err
			}
			obj = selectSkill(skills, u.stateProvider)
			return "use which skill?", nil
		}
		if useType == strings.ToLower(UseItem) {
			obj = selectItem(u.stateProvider)
		}
		if useType == strings.ToLower(UseObject) {
			obj = selectObject(u.stateProvider)
		}
	} else {
		obj = strings.Join(args[1:], " ")
	}

	target := selectTarget(obj, u.stateProvider)

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
