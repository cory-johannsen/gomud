package domain

type Team struct {
	Name string
	Jobs []*Job
}

func (t *Team) String() string {
	return t.Name
}

func (t *Team) Value() interface{} {
	return t
}

type Teams []*Team
