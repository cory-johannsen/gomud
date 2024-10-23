package domain

import (
	"fmt"
	"github.com/fatih/color"
)

type SkillType string

const (
	CommonSkill  SkillType = "Common"
	SpecialSkill SkillType = "Special"
)

type Skill struct {
	Name        string
	Description string
	Stat        string
	Type        SkillType
	Focuses     []string
}

func (s *Skill) Value() interface{} {
	return s
}

func (s *Skill) String() string {
	return fmt.Sprintf("Name: %s, Description: %s, Stat: %s, Type: %s, Focuses: %v", s.Name, s.Description, s.Stat, s.Type, s.Focuses)
}

type Skills []*Skill

var _ Property = &Skill{}

type RankedSkill struct {
	Skill             *Skill
	Rank              int
	SuccessPercentage int
}
type RankedSkills []*RankedSkill

func RankSkillsString(skills RankedSkills) string {
	cyan := color.New(color.FgCyan)
	cyanFunc := cyan.SprintFunc()
	green := color.New(color.FgGreen)
	msg := fmt.Sprintf("%s\n", cyanFunc("Skills"))

	stats := []string{
		Brutality,
		Muscle,
		Quickness,
		Savvy,
		Reasoning,
		Grit,
		Flair,
	}
	for _, stat := range stats {
		msg += fmt.Sprintf("    %s\n", cyanFunc(stat))
		for _, skill := range skills {
			if skill.Skill.Stat != stat {
				continue
			}
			pct := cyan.Sprintf("%d%%", skill.SuccessPercentage)
			rank := cyan.Sprintf("%d", skill.Rank)
			if skill.Rank > 0 {
				pct = green.Sprintf("%d%%", skill.SuccessPercentage)
				rank = green.Sprintf("%d", skill.Rank)
			}
			padding := 16 - len(skill.Skill.Name)
			name := skill.Skill.Name
			for i := 0; i < padding; i++ {
				name += " "
			}
			msg += fmt.Sprintf("\t%s[%s]\t[%s ranks]\n", name, pct, rank)
		}
	}
	return msg
}
