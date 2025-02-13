package domain

type TeamSpec struct {
	Name string `yaml:"name"`
}

func SpecFromTeam(t *Team) *TeamSpec {
	return &TeamSpec{
		Name: t.Name,
	}
}

type Team struct {
	Name string
	Jobs []*Job
}

func (t *Team) String() string {
	return t.Name
}

type Teams []*Team
