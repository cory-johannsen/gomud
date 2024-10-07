package domain

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type Property interface {
	Value() interface{}
	String() string
}

const PropertyNotFound = "property not found"

const (
	AgeProperty                = "age"
	AlignmentProperty          = "alignment"
	ArchetypeProperty          = "archetype"
	BackgroundProperty         = "background"
	BackgroundTraitProperty    = "backgroundTrait"
	BirthSeasonProperty        = "birthSeason"
	ConsumedAdvancesProperty   = "consumedAdvances"
	DistinguishingMarkProperty = "distinguishingMark"
	DrawbackProperty           = "drawback"
	ExperienceProperty         = "experience"
	JobProperty                = "job"
	TeamProperty               = "team"
	TattooProperty             = "tattoo"
	SkillRanksProperty         = "skillRanks"
	StatsProperty              = "stats"
	TalentsProperty            = "talents"
)

type BaseProperty struct {
	Val interface{}
}

func (p *BaseProperty) Value() interface{} {
	return p.Val
}
func (p *BaseProperty) String() string {
	return fmt.Sprintf("%v", p.Val)
}

var _ Property = &BaseProperty{}

type Character struct {
	Name string
	Data map[string]Property
}

type ConsumedAdvance struct {
	Job    string
	Stat   string
	Amount int
}
type ConsumedAdvances map[string][]ConsumedAdvance

func (c ConsumedAdvances) String() string {
	msg := ""
	for job, advances := range c {
		msg += fmt.Sprintf("%s\n", job)
		for _, advance := range advances {
			msg += fmt.Sprintf("\t%s: %d\n", advance.Stat, advance.Amount)
		}
	}
	return msg
}
func (c ConsumedAdvances) Value() interface{} {
	return c
}

func (c ConsumedAdvances) ConsumedAdvance(job string, stat string) int {
	advances, ok := c[job]
	if !ok {
		return 0
	}
	for _, a := range advances {
		if a.Stat == stat {
			return a.Amount
		}
	}
	return 0
}

var _ Property = &ConsumedAdvances{}

type SkillRank struct {
	Job   *Job
	Skill *Skill
}
type SkillRanks []*SkillRank

func (s SkillRanks) String() string {
	msg := ""
	for _, rank := range s {
		msg += fmt.Sprintf("%s\n%s\n", rank.Skill.Name, rank.Skill.Description)
	}
	return msg
}
func (s SkillRanks) Value() interface{} {
	return s
}

var _ Property = &SkillRanks{}

type Player struct {
	Character
	Id       *int
	Password string
}

func NewPlayer(id *int, name string, password string, data map[string]Property) *Player {
	p := &Player{
		Character: Character{
			Name: name,
			Data: data,
		},
		Id:       id,
		Password: password,
	}
	if p.Data == nil {
		p.Data = make(map[string]Property)
	}
	return p
}

func (p *Player) ValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (p *Player) GetProperty(key string) (Property, error) {
	result, ok := p.Data[key]
	if !ok {
		return nil, errors.New(PropertyNotFound)
	}
	return result, nil
}

func (p *Player) String() string {
	msg := fmt.Sprintf("Name: %s\n", p.Name)
	// Enforce the ordering of the character properties
	properties := []string{
		StatsProperty,
		ConsumedAdvancesProperty,
		ExperienceProperty,
		AlignmentProperty,
		BirthSeasonProperty,
		AgeProperty,
		BackgroundProperty,
		DistinguishingMarkProperty,
		TattooProperty,
		DrawbackProperty,
		ArchetypeProperty,
		JobProperty,
		SkillRanksProperty,
		TalentsProperty,
		BackgroundTraitProperty,
		TeamProperty}
	for _, k := range properties {
		v, ok := p.Data[k]
		if !ok {
			log.Printf("property %s not found when displaying player", k)
			continue
		}
		switch k {
		case AgeProperty:
			msg += fmt.Sprintf("  Age - %d\n", v.(*BaseProperty).Val.(int))
			continue
		case AlignmentProperty:
			msg += fmt.Sprintf("  Alignment - %s/%s\n\tOrder: %s (rank: %d)\n\tChaos: %s (rank: %d)\n\tCorruption: %d\n", v.(*Alignment).Order.Name, v.(*Alignment).Chaos.Name, v.(*Alignment).Order.Name, v.(*Alignment).Order.Rank, v.(*Alignment).Chaos.Name, v.(*Alignment).Chaos.Rank, v.(*Alignment).Corruption)
			continue
		case ArchetypeProperty:
			msg += fmt.Sprintf("  Archetype - %s\n", v.(*Archetype).Name)
			for _, trait := range v.(*Archetype).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name, effect.Description)
				}
			}
			continue
		case BackgroundProperty:
			msg += fmt.Sprintf("  Background - \n\t%s\n\t%s\n", v.(*Background).Name, v.(*Background).Description)
			for _, trait := range v.(*Background).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name, effect.Description)
				}
			}
			continue
		case BackgroundTraitProperty:
			msg += fmt.Sprintf("  Background Trait - \n\t%s\n\t%s\n", v.(*Trait).Name, v.(*Trait).Description)
			for _, effect := range v.(*Trait).Effects {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n", effect.Name, effect.Description)
			}
			continue
		case BirthSeasonProperty:
			msg += fmt.Sprintf("  Birth Season - %s\n", v.(Season))
			continue
		case ConsumedAdvancesProperty:
			msg += "  Consumed Advances: \n"
			for job, advances := range v.(ConsumedAdvances) {
				msg += fmt.Sprintf("\t%s\n", job)
				for _, advance := range advances {
					msg += fmt.Sprintf("\t\t%s: %d\n", advance.Stat, advance.Amount)
				}
			}
			continue
		case DistinguishingMarkProperty:
			msg += "  Distinguishing Marks: \n"
			for _, mark := range v.(DistinguishingMarks) {
				msg += fmt.Sprintf("\t%s\n", mark)
			}
			continue
		case DrawbackProperty:
			msg += fmt.Sprintf("  Drawback - \n\t%s\n\tDescription: %s\n\tEffect: \n\t\t%s\n\t\tdesc\n\t\teffect\n", v.(*Drawback).Name, v.(*Drawback).Description, v.(*Drawback).Effect)

			continue
		case JobProperty:
			msg += fmt.Sprintf("  Job - \n\t%s\n\tDescription: %s\n\tArchetype: %s\n\tTier: %s\n", v.(*Job).Name, v.(*Job).Description, v.(*Job).Archetype.Name, v.(*Job).Tier)
			continue
		case StatsProperty:
			msg += fmt.Sprintf("  Stats - \n\tFighting: %d\n\tMuscle: %d\n\tSpeed: %d\n\tSavvy: %d\n\tSmarts: %d\n\tGrit: %d\n\tFlair: %d\n", v.(*Stats).Fighting, v.(*Stats).Muscle, v.(*Stats).Speed, v.(*Stats).Savvy, v.(*Stats).Smarts, v.(*Stats).Grit, v.(*Stats).Flair)
			continue
		case SkillRanksProperty:
			msg += "  Skill Ranks: \n"
			for _, rank := range v.(SkillRanks) {
				msg += fmt.Sprintf("\t%s (from %s)\n\t%s\n", rank.Skill.Name, rank.Job.Name, rank.Skill.Description)
			}
			continue
		case TattooProperty:
			msg += fmt.Sprintf("  Tattoo - \n\t\"%s\" on your %s\n", v.(*Tattoo).Description, v.(*Tattoo).Location)
			continue
		case TalentsProperty:
			msg += "  Talents: \n"
			for _, talent := range v.(Talents) {
				msg += fmt.Sprintf("\t%s\n\t%s\n\t%s\n", talent.Name, talent.Description, talent.Effect.Name)
			}
		case TeamProperty:
			msg += fmt.Sprintf("  Team - %s\n", v.(*Team).Name)
			continue
		default:
			msg += fmt.Sprintf("  %s: %s\n", k, v)
		}
	}
	return msg
}

func (p *Player) Experience() int {
	return p.Data[ExperienceProperty].(*BaseProperty).Val.(int)
}
func (p *Player) AddExperience(exp int) {
	p.Data[ExperienceProperty] = &BaseProperty{Val: p.Experience() + exp}
}
func (p *Player) DeductExperience(exp int) {
	p.Data[ExperienceProperty] = &BaseProperty{Val: p.Experience() - exp}
}

func (p *Player) Job() *Job {
	return p.Data[JobProperty].(*Job)
}

func (p *Player) PurchaseSkillRank(job *Job, skill *Skill, exp int) {
	skillRanks := p.Data[SkillRanksProperty].(SkillRanks)
	skillRanks = append(skillRanks, &SkillRank{
		Job:   job,
		Skill: skill,
	})
	p.DeductExperience(exp)
}

func (p *Player) ConsumedBonusAdvances() ConsumedAdvances {
	return p.Data[ConsumedAdvancesProperty].(ConsumedAdvances)
}
func (p *Player) ConsumeBonusAdvance(job string, stat string, exp int) {
	consumed := p.ConsumedBonusAdvances()
	advances, ok := consumed[job]
	if !ok {
		advances = make([]ConsumedAdvance, 0)
	}
	var advance *ConsumedAdvance
	for _, a := range advances {
		if a.Stat == stat {
			advance = &a
			break
		}
	}
	if advance == nil {
		advance = &ConsumedAdvance{
			Job:    job,
			Stat:   stat,
			Amount: 1,
		}
		advances = append(advances, *advance)
	} else {
		advance.Amount++
	}

	p.Data[ConsumedAdvancesProperty] = consumed
	p.DeductExperience(exp)
}

func (p *Player) ConsumeTalent(job *Job, talent *Talent, exp int) {
	var talents Talents
	if _, ok := p.Data[TalentsProperty]; !ok {
		talents = make(Talents, 0)
		p.Data[TalentsProperty] = talents
	}
	talents = append(talents, talent)
	p.DeductExperience(exp)
}

func (p *Player) SkillRanks() SkillRanks {
	return p.Data[SkillRanksProperty].(SkillRanks)
}

func (p *Player) Talents() Talents {
	return p.Data[TalentsProperty].(Talents)
}

func (p *Player) Stats() *Stats {
	return p.Data[StatsProperty].(*Stats)
}

func (p *Player) Background() *Background {
	return p.Data[BackgroundProperty].(*Background)
}

func (p *Player) BackgroundTrait() *Trait {
	return p.Data[BackgroundTraitProperty].(*Trait)
}
