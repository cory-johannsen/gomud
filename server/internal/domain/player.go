package domain

import (
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/io"
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
	ConditionProperty          = "condition"
	ConsumedAdvancesProperty   = "consumedAdvances"
	DisordersProperty          = "disorders"
	DistinguishingMarkProperty = "distinguishingMark"
	DrawbackProperty           = "drawback"
	ExperienceProperty         = "experience"
	FatePointsProperty         = "fatePoints"
	InjuriesProperty           = "injuries"
	JobProperty                = "job"
	PerilProperty              = "peril"
	ReputationPointsProperty   = "reputationPoints"
	RoomProperty               = "room"
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
type ConsumedAdvances map[string][]*ConsumedAdvance

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

type PerilCondition int

const (
	PerilConditionUnhindered PerilCondition = iota
	PerilConditionImperiled
	PerilConditionIgnore1SkillRank
	PerilConditionIgnore2SkillRanks
	PerilConditionIgnore3SkillRanks
	PerilConditionIncapacitated
)

func (p PerilCondition) String() string {
	switch p {
	case PerilConditionUnhindered:
		return "Unhindered"
	case PerilConditionImperiled:
		return "Imperiled"
	case PerilConditionIgnore1SkillRank:
		return "Ignore 1 Skill Rank"
	case PerilConditionIgnore2SkillRanks:
		return "Ignore 2 Skill Ranks"
	case PerilConditionIgnore3SkillRanks:
		return "Ignore 3 Skill Ranks"
	case PerilConditionIncapacitated:
		return "INCAPACITATED!"
	}
	return "undefined"
}

type Peril struct {
	Threshold int
	Condition PerilCondition
}

func (p Peril) Value() interface{} {
	return p
}

func (p Peril) String() string {
	return fmt.Sprintf("Peril Threshold: %d\nPeril Condition: %s", p.Threshold, p.Condition.String())
}

var _ Property = &Peril{}

type Condition string

func (c Condition) Value() interface{} {
	return c
}

func (c Condition) String() string {
	return string(c)
}

const (
	ConditionUnharmed          Condition = "Unharmed"
	ConditionLightlyWounded    Condition = "Lightly Wounded"
	ConditionModeratelyWounded Condition = "Moderately Wounded"
	ConditionSeriouslyWounded  Condition = "Seriously Wounded"
	ConditionGrievouslyWounded Condition = "Grievously Wounded"
	ConditionSlain             Condition = "Slain!"
)

var RankedConditions = []Condition{
	ConditionUnharmed,
	ConditionLightlyWounded,
	ConditionModeratelyWounded,
	ConditionSeriouslyWounded,
	ConditionGrievouslyWounded,
	ConditionSlain,
}

func (c Condition) Rank() int {
	for i, condition := range RankedConditions {
		if c == condition {
			return i
		}
	}
	return -1
}

func (c Condition) LessSeriousThan(other Condition) bool {
	if c.Rank() < other.Rank() {
		return true
	}
	return false
}

var _ Property = Condition("")

type Player struct {
	Character
	Id         *int
	Password   string
	Connection io.Connection
	LoggedIn   bool
}

type Players map[int]*Player

func NewPlayer(id *int, name string, password string, data map[string]Property, conn io.Connection) *Player {
	p := &Player{
		Character: Character{
			Name: name,
			Data: data,
		},
		Id:         id,
		Password:   password,
		Connection: conn,
		LoggedIn:   false,
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
		FatePointsProperty,
		ReputationPointsProperty,
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
		TeamProperty,
		ConditionProperty,
		DisordersProperty,
		PerilProperty,
		InjuriesProperty,
	}
	for _, k := range properties {
		v, ok := p.Data[k]
		if !ok {
			log.Printf("property %s not found when displaying player", k)
			continue
		}
		switch k {
		case AgeProperty:
			msg += fmt.Sprintf("  Age - %d\n", v.(*BaseProperty).Val.(int))
		case AlignmentProperty:
			if v == nil {
				log.Warnf("alignment property is nil")
			}
			msg += fmt.Sprintf("  Alignment - %s/%s\n\tOrder: %s (rank: %d)\n\tChaos: %s (rank: %d)\n\tCorruption: %d\n", v.(*Alignment).Order.Name, v.(*Alignment).Chaos.Name, v.(*Alignment).Order.Name, v.(*Alignment).Order.Rank, v.(*Alignment).Chaos.Name, v.(*Alignment).Chaos.Rank, v.(*Alignment).Corruption)
		case ArchetypeProperty:
			msg += fmt.Sprintf("  Archetype - %s\n", v.(*Archetype).Name)
			for _, trait := range v.(*Archetype).Traits {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n\t\tEffects:\n", trait.Name, trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t\t\t%s\n\t\t\t%s\n", effect.Name(), effect.Description())
				}
			}
		case BackgroundProperty:
			msg += fmt.Sprintf("  Background - \n\t%s\n\t%s\n", v.(*Background).Name, v.(*Background).Description)
		case BackgroundTraitProperty:
			msg += fmt.Sprintf("  Background Trait - \n\t%s\n\t%s\n", v.(*Trait).Name, v.(*Trait).Description)
			for _, effect := range v.(*Trait).Effects {
				msg += fmt.Sprintf("\t\t%s\n\t\t%s\n", effect.Name(), effect.Description())
			}
		case BirthSeasonProperty:
			msg += fmt.Sprintf("  Birth Season - %s\n", v.(Season))
		case ConditionProperty:
			msg += fmt.Sprintf("  Condition - %s\n", v.(Condition))
		case ConsumedAdvancesProperty:
			msg += "  Bonus Advances: \n"
			for job, advances := range v.(ConsumedAdvances) {
				msg += fmt.Sprintf("\t%s\n", job)
				for _, advance := range advances {
					msg += fmt.Sprintf("\t\t%s: %d\n", advance.Stat, advance.Amount)
				}
			}
		case DisordersProperty:
			msg += "  Disorders: \n"
			for _, disorder := range v.(Disorders) {
				msg += fmt.Sprintf("\t%s\n\t%s\n", disorder.Name, disorder.Description)
			}
		case DistinguishingMarkProperty:
			msg += "  Distinguishing Marks: \n"
			for _, mark := range v.(DistinguishingMarks) {
				msg += fmt.Sprintf("\t%s\n", mark)
			}
		case DrawbackProperty:
			msg += fmt.Sprintf("  Drawback - \n\t%s\n\tDescription: %s\n\tEffect: \n\t\t%s", v.(*Drawback).Name, v.(*Drawback).Description, v.(*Drawback).Effect.Description())
		case ExperienceProperty:
			msg += fmt.Sprintf("  Experience - %d\n", v.(*BaseProperty).Val.(int))
		case FatePointsProperty:
			msg += fmt.Sprintf("  Fate Points - %d\n", v.(*BaseProperty).Val.(int))
		case InjuriesProperty:
			msg += "  Injuries: \n"
			for _, injury := range v.(Injuries) {
				msg += fmt.Sprintf("\t%s\n", injury)
			}
		case JobProperty:
			msg += fmt.Sprintf("  Job - \n\t%s\n\tDescription: %s\n\tArchetype: %s\n\tTier: %s\n", v.(*Job).Name, v.(*Job).Description, v.(*Job).Archetype.Name, v.(*Job).Tier)
		case PerilProperty:
			msg += fmt.Sprintf("  Peril - \n\tThreshold: %d\n\tCondition: %s", v.(*Peril).Threshold, v.(*Peril).Condition.String())
		case StatsProperty:
			stats := v.(*Stats)
			bonuses := p.StatBonuses()
			msg += fmt.Sprintf("  Stats - \n\tFighting: %d [%d]\n\tMuscle: %d [%d]\n\tSpeed: %d [%d]\n\tSavvy: %d [%d]\n\tSmarts: %d [%d]\n\tGrit: %d [%d]\n\tFlair: %d [%d]\n",
				stats.Fighting, bonuses.Fighting, stats.Muscle, bonuses.Muscle,
				stats.Speed, bonuses.Speed, stats.Savvy, bonuses.Savvy,
				stats.Smarts, bonuses.Smarts, stats.Grit, bonuses.Grit,
				stats.Flair, bonuses.Flair)
		case SkillRanksProperty:
			msg += "  Skill Ranks: \n"
			for _, rank := range v.(SkillRanks) {
				msg += fmt.Sprintf("\t%s (from %s)\n", rank.Skill.Name, rank.Job.Name)
			}
		case TattooProperty:
			msg += fmt.Sprintf("  Tattoo - \n\t\"%s\" on your %s\n", v.(*Tattoo).Description, v.(*Tattoo).Location)
		case TalentsProperty:
			msg += "  Talents: \n"
			for _, talent := range v.(Talents) {
				msg += fmt.Sprintf("\t%s\n\t\t%s\n\t\t%s\n", talent.Name, talent.Description, talent.Effect.Description())
			}
		case TeamProperty:
			msg += fmt.Sprintf("  Team - %s\n", v.(*Team).Name)
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
	skillRanks := p.SkillRanks()
	skillRanks = append(skillRanks, &SkillRank{
		Job:   job,
		Skill: skill,
	})
	p.Data[SkillRanksProperty] = skillRanks
	p.DeductExperience(exp)
}

func (p *Player) SkillRanks() SkillRanks {
	if _, ok := p.Data[SkillRanksProperty]; !ok {
		p.Data[SkillRanksProperty] = make(SkillRanks, 0)
	}
	return p.Data[SkillRanksProperty].(SkillRanks)
}

func (p *Player) HasSkillRank(job *Job, skil *Skill) bool {
	skillRanks := p.SkillRanks()
	for _, rank := range skillRanks {
		if rank.Job == job && rank.Skill == skil {
			return true
		}
	}
	return false
}

func (p *Player) ConsumedBonusAdvances() ConsumedAdvances {
	if _, ok := p.Data[ConsumedAdvancesProperty]; !ok {
		p.Data[ConsumedAdvancesProperty] = make(ConsumedAdvances)
	}
	return p.Data[ConsumedAdvancesProperty].(ConsumedAdvances)
}

func (p *Player) ConsumeBonusAdvance(job string, stat string, exp int) {
	consumed := p.ConsumedBonusAdvances()
	advances, ok := consumed[job]
	if !ok {
		advances = make([]*ConsumedAdvance, 0)
	}
	var advance *ConsumedAdvance
	for _, a := range advances {
		if a.Stat == stat {
			advance = a
			break
		}
	}
	if advance == nil {
		advance = &ConsumedAdvance{
			Job:    job,
			Stat:   stat,
			Amount: 1,
		}
		advances = append(advances, advance)
	} else {
		advance.Amount++
	}
	consumed[job] = advances
	p.Data[ConsumedAdvancesProperty] = consumed
	p.DeductExperience(exp)
}

func (p *Player) ConsumeTalent(job *Job, talent *Talent, exp int) {
	talents := p.Talents()
	talents = append(talents, talent)
	p.Data[TalentsProperty] = talents
	p.DeductExperience(exp)
}

func (p *Player) Talents() Talents {
	if _, ok := p.Data[TalentsProperty]; !ok {
		p.Data[TalentsProperty] = make(Talents, 0)
	}
	return p.Data[TalentsProperty].(Talents)
}

func (p *Player) HasTalent(job *Job, talent *Talent) bool {
	talents := p.Talents()
	for _, t := range talents {
		if t == talent {
			return true
		}
	}
	return false
}

func (p *Player) Stats() *Stats {
	if _, ok := p.Data[StatsProperty]; !ok {
		p.Data[StatsProperty] = &Stats{}
	}
	return p.Data[StatsProperty].(*Stats)
}

func bonusFromStat(stat int) int {
	return stat / 10
}

func (p *Player) StatBonuses() *Stats {
	stats := p.Stats()
	bonuses := &Stats{
		Fighting: bonusFromStat(stats.Fighting),
		Muscle:   bonusFromStat(stats.Muscle),
		Speed:    bonusFromStat(stats.Speed),
		Savvy:    bonusFromStat(stats.Savvy),
		Smarts:   bonusFromStat(stats.Smarts),
		Grit:     bonusFromStat(stats.Grit),
		Flair:    bonusFromStat(stats.Flair),
	}
	advances := p.ConsumedBonusAdvances()
	for job, jobAdvances := range advances {
		for _, advance := range jobAdvances {
			switch advance.Stat {
			case "Fighting":
				bonuses.Fighting += advance.Amount
			case "Muscle":
				bonuses.Muscle += advance.Amount
			case "Speed":
				bonuses.Speed += advance.Amount
			case "Savvy":
				bonuses.Savvy += advance.Amount
			case "Smarts":
				bonuses.Smarts += advance.Amount
			case "Grit":
				bonuses.Grit += advance.Amount
			case "Flair":
				bonuses.Flair += advance.Amount
			default:
				log.Warnf("unknown stat %s for job %s", advance.Stat, job)
			}
		}
	}
	return bonuses
}

func (p *Player) Background() *Background {
	return p.Data[BackgroundProperty].(*Background)
}

func (p *Player) BackgroundTrait() *Trait {
	return p.Data[BackgroundTraitProperty].(*Trait)
}

func (p *Player) Room() *Room {
	r := p.Data[RoomProperty]
	if r == nil {
		return nil
	}
	return r.(*Room)
}

func (p *Player) SetRoom(r *Room) {
	currentRoom := p.Room()
	if r == currentRoom {
		return
	}
	p.Data[RoomProperty] = r
}

func (p *Player) RoomHandler(player *Player, action string) {
	if player == p {
		return
	}
	p.Connection.Writeln(fmt.Sprintf("\n%s the %s %ss the room", player.Name, player.Job().Name, action))
}

func (p *Player) Peril() *Peril {
	return p.Data[PerilProperty].(*Peril)
}

func (p *Player) SetPeril(peril *Peril) {
	p.Data[PerilProperty] = peril
}

func (p *Player) FatePoints() int {
	return p.Data[FatePointsProperty].(*BaseProperty).Val.(int)
}

func (p *Player) AddFatePoints(points int) {
	p.Data[FatePointsProperty] = &BaseProperty{Val: p.FatePoints() + points}
}

func (p *Player) SubtractFatePoints(points int) {
	result := p.FatePoints() - points
	if result < 0 {
		result = 0
	}
	p.Data[FatePointsProperty] = &BaseProperty{Val: result}
}

func (p *Player) ReputationPoints() int {
	return p.Data[ReputationPointsProperty].(*BaseProperty).Val.(int)
}

func (p *Player) AddReputationPoints(points int) {
	p.Data[ReputationPointsProperty] = &BaseProperty{Val: p.ReputationPoints() + points}
}

func (p *Player) SubtractReputationPoints(points int) {
	result := p.ReputationPoints() - points
	if result < 0 {
		result = 0
	}
	p.Data[ReputationPointsProperty] = &BaseProperty{Val: result}
}

func (p *Player) Alignment() *Alignment {
	return p.Data[AlignmentProperty].(*Alignment)
}

func (p *Player) AddCorruption(corruption int) {
	rank := p.Alignment().Chaos.Rank
	p.Alignment().AddCorruption(corruption)
	if p.Alignment().Chaos.Rank > rank {
		// add a disorder
		p.Alignment().ResetCorruption()
	}
}

func (p *Player) AddOrderRank(rank int) {
	p.Alignment().AddOrderRank(rank)
	if p.Alignment().Order.Rank >= 10 {
		p.AddFatePoints(1)
		p.Alignment().ResetOrderRank()
	}
}

func (p *Player) Condition() Condition {
	condition, ok := p.Data[ConditionProperty]
	if !ok {
		condition = ConditionUnharmed
		p.Data[ConditionProperty] = ConditionUnharmed
	}
	return condition.(Condition)
}

func (p *Player) Injuries() Injuries {
	if _, ok := p.Data[InjuriesProperty]; !ok {
		p.Data[InjuriesProperty] = make(Injuries, 0)
	}
	return p.Data[InjuriesProperty].(Injuries)
}
