package domain

import (
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/io"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type Property interface {
	String() string
}

const PropertyNotFound = "property not found"

const (
	AgeProperty                 = "age"
	AlignmentProperty           = "alignment"
	ArchetypeProperty           = "archetype"
	BackgroundProperty          = "background"
	BackgroundTraitProperty     = "backgroundTrait"
	BirthSeasonProperty         = "birthSeason"
	ConditionProperty           = "condition"
	ConsumedAdvancesProperty    = "consumedAdvances"
	DisordersProperty           = "disorders"
	DistinguishingMarksProperty = "distinguishingMarks"
	DrawbackProperty            = "drawback"
	ExperienceProperty          = "experience"
	EngagedProperty             = "engaged"
	FatePointsProperty          = "fatePoints"
	InjuriesProperty            = "injuries"
	InventoryProperty           = "inventory"
	JobProperty                 = "job"
	PerilProperty               = "peril"
	PoornessProperty            = "poorness"
	ReputationPointsProperty    = "reputationPoints"
	RoomProperty                = "room"
	TeamProperty                = "team"
	TattooProperty              = "tattoo"
	SkillRanksProperty          = "skillRanks"
	StatsProperty               = "stats"
	SleepingProperty            = "sleeping"
	TalentsProperty             = "talents"
	TagsProperty                = "tags"
	UpbringingProperty          = "upbringing"
)

type BaseProperty struct {
	Val interface{}
}

func (p *BaseProperty) String() string {
	return fmt.Sprintf("%v", p.Val)
}

var _ Property = &BaseProperty{}

type ConsumedAdvance struct {
	Job    string `yaml:"job"`
	Stat   string `yaml:"stat"`
	Amount int    `yaml:"amount"`
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

func PerilConditionFromString(condition string) PerilCondition {
	switch condition {
	case "Unhindered":
		return PerilConditionUnhindered
	case "Imperiled":
		return PerilConditionImperiled
	case "Ignore 1 Skill Rank":
		return PerilConditionIgnore1SkillRank
	case "Ignore 2 Skill Ranks":
		return PerilConditionIgnore2SkillRanks
	case "Ignore 3 Skill Ranks":
		return PerilConditionIgnore3SkillRanks
	case "INCAPACITATED!":
		return PerilConditionIncapacitated
	}
	return PerilConditionUnhindered
}

type Peril struct {
	Threshold int
	Condition PerilCondition
}

func (p Peril) String() string {
	return fmt.Sprintf("Peril Threshold: %d\nPeril Condition: %s", p.Threshold, p.Condition.String())
}

var _ Property = &Peril{}

type Condition string

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

type Tags map[string]string

func (t Tags) String() string {
	msg := ""
	for k, v := range t {
		msg += fmt.Sprintf("%s: %s\n", k, v)
	}
	return msg
}

var _ Property = Tags{}

type CharacterSpec struct {
	Id   int
	Name string
	Data map[string]interface{}
}

type CharacterSpecs []*CharacterSpecs

// Character a character in the game that contains all the core properties
type Character struct {
	Id   *int
	Name string
	Data map[string]Property
}

type Characters []*Character

func NewCharacter(id *int, name string, data map[string]Property) *Character {
	return &Character{
		Id:   id,
		Name: name,
		Data: data,
	}
}

func (c *Character) Alignment() *Alignment {
	return c.Data[AlignmentProperty].(*Alignment)
}

func (c *Character) AddCorruption(corruption int) {
	rank := c.Alignment().Chaos.Rank
	c.Alignment().AddCorruption(corruption)
	if c.Alignment().Chaos.Rank > rank {
		// TODO: add a disorder
		c.Alignment().ResetCorruption()
	}
}

func (c *Character) AddOrderRank(rank int) {
	c.Alignment().AddOrderRank(rank)
	if c.Alignment().Order.Rank >= 10 {
		c.AddFatePoints(1)
		c.Alignment().ResetOrderRank()
	}
}

func (c *Character) Condition() Condition {
	condition, ok := c.Data[ConditionProperty]
	if !ok {
		condition = ConditionUnharmed
		c.Data[ConditionProperty] = ConditionUnharmed
	}
	return condition.(Condition)
}

func (c *Character) Engaged() bool {
	engaged, ok := c.Data[EngagedProperty]
	if !ok {
		return false
	}
	return engaged.(*BaseProperty).Val.(bool)
}

func (c *Character) Injuries() Injuries {
	if _, ok := c.Data[InjuriesProperty]; !ok {
		c.Data[InjuriesProperty] = make(Injuries, 0)
	}
	return c.Data[InjuriesProperty].(Injuries)
}

func (c *Character) Inventory() *Inventory {
	return c.Data[InventoryProperty].(*Inventory)
}

func (c *Character) Poorness() Poorness {
	return c.Data[PoornessProperty].(Poorness)
}

func (c *Character) Upbringing() *Upbringing {
	return c.Data[UpbringingProperty].(*Upbringing)
}

func (c *Character) Drawback() *Drawback {
	return c.Data[DrawbackProperty].(*Drawback)
}

func (c *Character) DistinguishingMarks() DistinguishingMarks {
	return c.Data[DistinguishingMarksProperty].(DistinguishingMarks)
}

func (c *Character) PrimaryStat() string {
	return c.Upbringing().Stat
}

func (c *Character) Skills(allSkills Skills) RankedSkills {
	byStat := make(map[string]RankedSkills)
	stats := c.Stats()
	for _, skill := range allSkills {
		if _, ok := byStat[skill.Stat]; !ok {
			byStat[skill.Stat] = make(RankedSkills, 0)
		}
		statValue := stats.StatValue(skill.Stat)
		byStat[skill.Stat] = append(byStat[skill.Stat], &RankedSkill{
			Skill:             skill,
			Rank:              0,
			SuccessPercentage: statValue,
		})
	}
	for _, skillRank := range c.SkillRanks() {
		rankedSkills, ok := byStat[skillRank.Skill.Stat]
		if !ok {
			log.Printf("skill %s not found in stats", skillRank.Skill.Name)
			continue
		}
		var skill *RankedSkill
		for _, s := range rankedSkills {
			if s.Skill == skillRank.Skill {
				skill = s
				break
			}
		}
		if skill == nil {
			log.Printf("skill %s not found in stats", skillRank.Skill.Name)
			continue
		}
		skill.Rank++
		skill.SuccessPercentage += 10
	}
	rankedSkills := make(RankedSkills, 0)
	for _, skills := range byStat {
		rankedSkills = append(rankedSkills, skills...)
	}
	return rankedSkills
}

func (c *Character) Job() *Job {
	return c.Data[JobProperty].(*Job)
}

func (c *Character) SkillRanks() SkillRanks {
	if _, ok := c.Data[SkillRanksProperty]; !ok {
		c.Data[SkillRanksProperty] = make(SkillRanks, 0)
	}
	return c.Data[SkillRanksProperty].(SkillRanks)
}

func (c *Character) HasSkillRank(job *Job, skil *Skill) bool {
	skillRanks := c.SkillRanks()
	for _, rank := range skillRanks {
		if rank.Job == job && rank.Skill == skil {
			return true
		}
	}
	return false
}

func (c *Character) Talents() Talents {
	if _, ok := c.Data[TalentsProperty]; !ok {
		c.Data[TalentsProperty] = make(Talents, 0)
	}
	return c.Data[TalentsProperty].(Talents)
}

func (c *Character) HasTalent(job *Job, talent *Talent) bool {
	talents := c.Talents()
	for _, t := range talents {
		if t == talent {
			return true
		}
	}
	return false
}

func (c *Character) Stats() *Stats {
	if _, ok := c.Data[StatsProperty]; !ok {
		c.Data[StatsProperty] = &Stats{}
	}
	return c.Data[StatsProperty].(*Stats)
}

func (c *Character) Tags() Tags {
	tags, ok := c.Data[TagsProperty]
	if !ok {
		t := make(Tags)
		c.Data[TagsProperty] = t
		return t
	}
	return tags.(Tags)
}

func (c *Character) AddTag(tag string, value string) {
	tags := c.Tags()
	tags[tag] = value
	c.Data[TagsProperty] = tags
}

func (c *Character) RemoveTag(tag string) {
	tags := c.Tags()
	delete(tags, tag)
	c.Data[TagsProperty] = tags
}

func (c *Character) HasTag(tag string) bool {
	tags := c.Tags()
	_, ok := tags[tag]
	return ok
}

func (c *Character) Tag(tag string) *string {
	tags := c.Tags()
	t, ok := tags[tag]
	if !ok {
		return nil
	}
	return &t
}

func bonusFromStat(stat int) int {
	return stat / 10
}

func (c *Character) FatePoints() int {
	return c.Data[FatePointsProperty].(*BaseProperty).Val.(int)
}

func (c *Character) AddFatePoints(points int) {
	c.Data[FatePointsProperty] = &BaseProperty{Val: c.FatePoints() + points}
}

func (c *Character) SubtractFatePoints(points int) {
	result := c.FatePoints() - points
	if result < 0 {
		result = 0
	}
	c.Data[FatePointsProperty] = &BaseProperty{Val: result}
}

func (c *Character) ConsumedBonusAdvances() ConsumedAdvances {
	if _, ok := c.Data[ConsumedAdvancesProperty]; !ok {
		c.Data[ConsumedAdvancesProperty] = make(ConsumedAdvances)
	}
	return c.Data[ConsumedAdvancesProperty].(ConsumedAdvances)
}

func (c *Character) StatBonuses() *Stats {
	stats := c.Stats()
	bonuses := &Stats{
		Brutality: bonusFromStat(stats.Brutality),
		Muscle:    bonusFromStat(stats.Muscle),
		Quickness: bonusFromStat(stats.Quickness),
		Savvy:     bonusFromStat(stats.Savvy),
		Reasoning: bonusFromStat(stats.Reasoning),
		Grit:      bonusFromStat(stats.Grit),
		Flair:     bonusFromStat(stats.Flair),
	}
	advances := c.ConsumedBonusAdvances()
	for job, jobAdvances := range advances {
		for _, advance := range jobAdvances {
			switch advance.Stat {
			case "Brutality":
				bonuses.Brutality += advance.Amount
			case "Muscle":
				bonuses.Muscle += advance.Amount
			case "Quickness":
				bonuses.Quickness += advance.Amount
			case "Savvy":
				bonuses.Savvy += advance.Amount
			case "Reasoning":
				bonuses.Reasoning += advance.Amount
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

func (c *Character) Background() *Background {
	return c.Data[BackgroundProperty].(*Background)
}

func (c *Character) BackgroundTrait() *Trait {
	return c.Data[BackgroundTraitProperty].(*Trait)
}

func (c *Character) Room() *Room {
	r := c.Data[RoomProperty]
	if r == nil {
		return nil
	}
	return r.(*Room)
}

func (c *Character) SetRoom(r *Room) {
	currentRoom := c.Room()
	if r == currentRoom {
		return
	}
	c.Data[RoomProperty] = r
}

func (c *Character) Peril() *Peril {
	return c.Data[PerilProperty].(*Peril)
}

func (c *Character) SetPeril(peril *Peril) {
	c.Data[PerilProperty] = peril
}

func (c *Character) Sleeping() bool {
	sleeping, ok := c.Data[SleepingProperty]
	if !ok {
		return false
	}
	return sleeping.(*BaseProperty).Val.(bool)
}

func (c *Character) SetSleeping(sleeping bool) {
	c.Data[SleepingProperty] = &BaseProperty{Val: sleeping}
}

// Player a Character that is controlled by a user
type Player struct {
	Character
	Password   string
	Connection io.Connection
	LoggedIn   bool
}

type Players map[int]*Player

func NewPlayer(id *int, name string, password string, data map[string]Property, conn io.Connection) *Player {
	p := &Player{
		Character: Character{
			Id:   id,
			Name: name,
			Data: data,
		},
		Password:   password,
		Connection: conn,
		LoggedIn:   false,
	}
	if p.Data == nil {
		p.Data = make(map[string]Property)
	}
	if _, ok := p.Data[InventoryProperty]; !ok {
		p.Data[InventoryProperty] = NewInventory()
	}
	err := conn.EventBus().SubscribeAsync(event.RoomChannel, func(r *RoomEvent) {
		if r == nil || r.Room == nil || r.Room != p.Room() || r.Character == nil || r.Character == &p.Character {
			return
		}
		log.Printf("room %s received action %s from player %s", r.Room.Name, r.Action, r.Character.Name)
		switch r.Action {
		case event.RoomEventEnter:
			p.Connection.Writeln(fmt.Sprintf("%s the %s enters the room\n%s", r.Character.Name, r.Character.Job().Name, p.Prompt()))
		case event.RoomEventExit:
			p.Connection.Writeln(fmt.Sprintf("%s the %s leaves the room\n%s", r.Character.Name, r.Character.Job().Name, p.Prompt()))
		case event.RoomEventSay:
			args := make([]string, 0)
			for _, arg := range r.Args {
				args = append(args, arg.(string))
			}
			msg := strings.Join(args, " ")
			timeOfDaySensor := p.Connection.Sensors()["TimeOfDay"].(*htn.TimeOfDaySensor)
			timeOfDay, err := timeOfDaySensor.Get()
			if err != nil {
				log.Errorf("error getting hour of day: %s", err)
			}
			p.Connection.Writeln(fmt.Sprintf("\n[%02d:%02d] %s says \"%s\"\n%s", timeOfDay.Hour, timeOfDay.Minute, r.Character.Name, msg, p.Prompt()))
		}
	}, false)
	if err != nil {
		log.Warnf("error subscribing to room channel: %s", err)
		return nil
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

func (p Player) String() string {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	msg := fmt.Sprintf("%s: %s\n", cyan("Name"), p.Name)
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
		PoornessProperty,
		DistinguishingMarksProperty,
		TattooProperty,
		DrawbackProperty,
		UpbringingProperty,
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
		InventoryProperty,
	}
	for _, k := range properties {
		v, ok := p.Data[k]
		if !ok {
			log.Printf("property %s not found when displaying player", k)
			continue
		}
		switch k {
		case AgeProperty:
			msg += fmt.Sprintf("  %s: %d\n", cyan("Age"), v.(*BaseProperty).Val.(int))
		case AlignmentProperty:
			if v == nil {
				log.Warnf("alignment property is nil")
			}
			corruption := cyan(v.(*Alignment).Corruption)
			if v.(*Alignment).Corruption > 4 {
				if v.(*Alignment).Corruption > 7 {
					corruption = red(v.(*Alignment).Corruption)
				} else {
					corruption = yellow(v.(*Alignment).Corruption)
				}
			}
			orderRank := cyan(v.(*Alignment).Order.Rank)
			if v.(*Alignment).Order.Rank > 0 {
				orderRank = green(v.(*Alignment).Order.Rank)
			}
			chaosRank := cyan(v.(*Alignment).Chaos.Rank)
			if v.(*Alignment).Chaos.Rank > 0 {
				if v.(*Alignment).Chaos.Rank > 4 {
					if v.(*Alignment).Chaos.Rank > 7 {
						chaosRank = red(v.(*Alignment).Chaos.Rank)
					} else {
						chaosRank = yellow(v.(*Alignment).Chaos.Rank)
					}
				}
				chaosRank = green(v.(*Alignment).Chaos.Rank)
			}
			msg += fmt.Sprintf("  %s: %s/%s\n\tOrder: %s (rank: %s)\n\tChaos: %s (rank: %s)\n\tCorruption: %s\n",
				cyan("Alignment"), v.(*Alignment).Order.Name, v.(*Alignment).Chaos.Name,
				v.(*Alignment).Order.Name, orderRank,
				v.(*Alignment).Chaos.Name, chaosRank,
				corruption)
		case ArchetypeProperty:
			msg += fmt.Sprintf("  %s: %s\n", cyan("Archetype"), v.(*Archetype).Name)
			for _, trait := range v.(*Archetype).Traits {
				msg += fmt.Sprintf("\t%s\n\t  %s\n", magenta(trait.Name), trait.Description)
				for _, effect := range trait.Effects {
					msg += fmt.Sprintf("\t  %s\n", yellow(effect.Description()))
				}
			}
		case BackgroundProperty:
			msg += fmt.Sprintf("  %s\n\t%s\n\t%s\n", cyan("Background"), v.(*Background).Name, v.(*Background).Description)
		case BackgroundTraitProperty:
			msg += fmt.Sprintf("  %s\n\t%s\n\t%s\n", cyan("Background Trait"), v.(*Trait).Name, v.(*Trait).Description)
			for _, effect := range v.(*Trait).Effects {
				msg += fmt.Sprintf("\t\t%s\n", yellow(effect.Description()))
			}
		case BirthSeasonProperty:
			msg += fmt.Sprintf("  %s: %s\n", cyan("Birth Season"), v.(Season))
		case ConditionProperty:
			msg += fmt.Sprintf("  %s: %s\n", cyan("Condition"), v.(Condition))
		case ConsumedAdvancesProperty:
			msg += fmt.Sprintf("  %s\n", cyan("Bonus Advances"))
			if len(v.(ConsumedAdvances)) == 0 {
				msg += "\tNone\n"
			}
			for job, advances := range v.(ConsumedAdvances) {
				msg += fmt.Sprintf("\t%s\n", job)
				for _, advance := range advances {
					msg += fmt.Sprintf("\t\t%s: %s\n", advance.Stat, green(advance.Amount))
				}
			}
		case DisordersProperty:
			msg += fmt.Sprintf("  %s\n", cyan("Disorders"))
			if len(v.(Disorders)) == 0 {
				msg += "\tNone\n"
			}
			for _, disorder := range v.(Disorders) {
				msg += fmt.Sprintf("\t%s\n\t%s\n", disorder.Name, disorder.Description)
			}
		case DistinguishingMarksProperty:
			msg += fmt.Sprintf("  %s\n", cyan("Distinguishing Marks"))
			if len(v.(DistinguishingMarks)) == 0 {
				msg += "\tNone\n"
			}
			for _, mark := range v.(DistinguishingMarks) {
				msg += fmt.Sprintf("\t%s\n", mark)
			}
		case DrawbackProperty:
			msg += fmt.Sprintf("  %s\n\t%s\n\t  %s\n\t  %s\n", cyan("Drawback"), magenta(v.(*Drawback).Name), v.(*Drawback).Description, yellow(v.(*Drawback).Effect.Description()))
		case ExperienceProperty:
			val := cyan(v.(*BaseProperty).Val.(int))
			if v.(*BaseProperty).Val.(int) > 0 {
				val = green(v.(*BaseProperty).Val.(int))
			}
			msg += fmt.Sprintf("  %s: %s\n", cyan("Experience"), val)
		case InventoryProperty:
			inventory := v.(*Inventory)
			msg += fmt.Sprintf("  %s\n\tMain Hand: ", cyan("Inventory"))
			if inventory.MainHand() == nil {
				msg += "empty"
			} else {
				msg += inventory.MainHand().Name()
			}
			msg += fmt.Sprintf("\n\tOff Hand: ")
			if inventory.OffHand() == nil {
				msg += "empty"
			} else {
				msg += inventory.OffHand().Name()
			}
			msg += fmt.Sprintf("\n\tArmor: ")
			if inventory.Armor() == nil {
				msg += "empty"
			} else {
				msg += inventory.Armor().Name()
			}
			msg += fmt.Sprintf("\n\tCash: %d\n", inventory.Cash())
		case FatePointsProperty:
			val := cyan(v.(*BaseProperty).Val.(int))
			if v.(*BaseProperty).Val.(int) > 0 {
				val = green(v.(*BaseProperty).Val.(int))
			}
			msg += fmt.Sprintf("  %s: %s\n", cyan("Fate Points"), val)
		case InjuriesProperty:
			msg += fmt.Sprintf("  %s\n", cyan("Injuries"))
			if len(v.(Injuries)) == 0 {
				msg += "\tNone\n"
			}
			for _, injury := range v.(Injuries) {
				msg += fmt.Sprintf("\t%s\n", injury)
			}
		case JobProperty:
			msg += fmt.Sprintf("  %s\n\t%s\n\tDescription: %s\n\tArchetype: %s\n\tTier: %s\n",
				cyan("Job"), v.(*Job).Name, v.(*Job).Description, v.(*Job).Archetype.Name, v.(*Job).Tier)
		case PerilProperty:
			msg += fmt.Sprintf("  %s\n\tThreshold: %d\n\tCondition: %s\n", cyan("Peril"), v.(*Peril).Threshold, v.(*Peril).Condition.String())
		case PoornessProperty:
			msg += fmt.Sprintf("  %s: %s\n", cyan("Poorness"), v.(Poorness))
		case ReputationPointsProperty:
			val := cyan(v.(*BaseProperty).Val.(int))
			if v.(*BaseProperty).Val.(int) > 0 {
				val = green(v.(*BaseProperty).Val.(int))
			}
			msg += fmt.Sprintf("  %s: %s\n", cyan("Reputation Points"), val)
		case StatsProperty:
			stats := v.(*Stats)
			bonuses := p.StatBonuses()
			advances := p.ConsumedBonusAdvances()
			msg += fmt.Sprintf("  %s", StatsString(stats, bonuses, advances))
		case SkillRanksProperty:
			msg += fmt.Sprintf("  %s: \n", cyan("Skill Ranks"))
			if len(v.(SkillRanks)) == 0 {
				msg += "\tNone\n"
			}
			for _, rank := range v.(SkillRanks) {
				msg += fmt.Sprintf("\t%s (from %s)\n", rank.Skill.Name, rank.Job.Name)
			}
		case TattooProperty:
			msg += fmt.Sprintf("  %s\n\t\"%s\" on your %s\n", cyan("Tattoo"), v.(*Tattoo).Description, v.(*Tattoo).Location)
		case TalentsProperty:
			msg += fmt.Sprintf("  %s: \n", cyan("Talents"))
			if len(v.(Talents)) == 0 {
				msg += "\tNone\n"
			}
			for _, talent := range v.(Talents) {
				desc := strings.ReplaceAll(talent.Description, "\n", " ")
				msg += fmt.Sprintf("\t%s\n\t  %s\n\t  %s\n\n", magenta(talent.Name), desc, yellow(talent.Effect.Description()))
			}
		case TeamProperty:
			msg += fmt.Sprintf("  %s: %s\n", cyan("Team"), v.(*Team).Name)
		case UpbringingProperty:
			msg += fmt.Sprintf("  %s\n\t%s\n\tPrimary Stat: %s\n", cyan("Upbringing"), v.(*Upbringing).Name, v.(*Upbringing).Stat)
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

func (p *Player) PurchaseSkillRank(job *Job, skill *Skill, exp int) {
	skillRanks := p.SkillRanks()
	skillRanks = append(skillRanks, &SkillRank{
		Job:   job,
		Skill: skill,
	})
	p.Data[SkillRanksProperty] = skillRanks
	p.DeductExperience(exp)
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

func (p *Player) Prompt() string {
	cyan := color.New(color.FgCyan).SprintFunc()
	if p == nil || !p.LoggedIn {
		return fmt.Sprintf("%s ", cyan(">"))
	}
	player := p
	green := color.New(color.FgGreen).SprintFunc()
	timeOfDay, err := p.Connection.Sensors()["TimeOfDay"].(*htn.TimeOfDaySensor).Get()
	if err != nil {
		log.Errorf("error getting hour of day: %s", err)
	}
	timestamp := fmt.Sprintf("[%02d:%02d]", timeOfDay.Hour, timeOfDay.Minute)
	return fmt.Sprintf("%s %s [%s, %s]%s ", cyan(timestamp), cyan(player.Name), green(player.Condition()), green(player.Peril().Condition.String()), cyan(">"))
}
