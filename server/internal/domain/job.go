package domain

type Tier string

const (
	BasicTier        Tier = "Basic"
	IntermediateTier Tier = "Intermediate"
	AdvancedTier     Tier = "Advanced"
)

type JobSpec struct {
	Name           string         `yaml:"name"`
	Description    string         `yaml:"description"`
	Archetype      string         `yaml:"archetype"`
	Tier           Tier           `yaml:"tier"`
	ExperienceCost int            `yaml:"experienceCost"`
	Traits         []string       `yaml:"traits"`
	SkillRanks     []string       `yaml:"skillRanks"`
	BonusAdvances  map[string]int `yaml:"bonusAdvances"`
	Talents        []string       `yaml:"talents"`
}

func (s *JobSpec) String() string {
	return s.Name
}
func (s *JobSpec) Value() interface{} {
	return s
}

type BonusAdvances struct {
	Fighting int
	Muscle   int
	Speed    int
	Savvy    int
	Smarts   int
	Grit     int
	Flair    int
}

type Job struct {
	Name           string
	Description    string
	Archetype      *Archetype
	Tier           Tier
	ExperienceCost int
	Traits         Traits
	SkillRanks     Skills
	BonusAdvances  BonusAdvances
	Talents        Talents
}

func (j *Job) Value() interface{} {
	return j
}

func (j *Job) String() string {
	return j.Name
}

type Jobs []*Job

var _ Property = &Job{}
