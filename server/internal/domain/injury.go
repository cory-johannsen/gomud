package domain

type InjurySpec struct {
	Name     string `yaml:"name"`
	Severity string `yaml:"severity"`
	Effect   string `yaml:"effect"`
}

type Severity string

type Injury struct {
	Name     string
	Severity Severity
	Effect   Effect
}

func (i Injury) String() string {
	return i.Name
}

var _ Property = &Injury{}

type Injuries []*Injury

func (i Injuries) String() string {
	var msg string
	for _, injury := range i {
		msg += injury.Name + "\n"
	}
	return msg
}

var _ Property = Injuries{}
