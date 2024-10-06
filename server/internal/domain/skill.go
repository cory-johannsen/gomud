package domain

import "fmt"

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
