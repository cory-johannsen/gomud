package domain

type Tier string

const (
	BasicTier        Tier = "Basic"
	IntermediateTier Tier = "Intermediate"
	AdvancedTier     Tier = "Advanced"
)

type JobSpec struct {
	Name           string   `yaml:"name"`
	Description    string   `yaml:"description"`
	Archetype      string   `yaml:"archetype"`
	Tier           Tier     `yaml:"tier"`
	ExperienceCost int      `yaml:"experienceCost"`
	Traits         []string `yaml:"traits"`
}

func (s *JobSpec) String() string {
	return s.Name
}
func (s *JobSpec) Value() interface{} {
	return s
}

type Job struct {
	Name           string
	Description    string
	Archetype      *Archetype
	Tier           Tier
	ExperienceCost int
	Traits         Traits
}

type Jobs []*Job

func (j *Job) String() string {
	return j.Name
}

func (j *Job) Value() interface{} {
	return j
}

var _ Property = &Job{}
