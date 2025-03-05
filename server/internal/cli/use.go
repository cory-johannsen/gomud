package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/fatih/color"
	"strings"
)

const (
	UseItem  string = "item"
	UseSkill string = "skill"

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

func (u *UseHandler) useSkill(skillName string) (string, error) {
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

func (u *UseHandler) useItem(itemName string) (string, error) {
	return "using item " + itemName, nil
}

func (u *UseHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) == 0 {
		return "use what? (skill/item)", nil
	}
	useType := strings.ToLower(args[0])
	if len(args) < 2 {
		if useType == strings.ToLower(UseSkill) {
			return "use which skill?", nil
		}
		if useType == strings.ToLower(UseItem) {
			return "use which item?", nil
		}
	}
	obj := strings.Join(args[1:], " ")
	if useType == strings.ToLower(UseSkill) {
		return u.useSkill(obj)
	}
	return u.useItem(obj)
}

func (u *UseHandler) Help(args []string) string {
	return "use a skill or item.  Usage: use [skill|item] [name]"
}

func (u *UseHandler) State() domain.GameState {
	return u.stateProvider()
}

var _ Handler = &UseHandler{}
