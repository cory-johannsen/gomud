package domain

type TalentSpec struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Effect      string `yaml:"effect"`
}

type Talent struct {
	Name        string
	Description string
	Effect      Effect
}
type Talents []*Talent

func (t *Talent) String() string {
	return t.Name
}
func (t *Talent) Value() interface{} {
	return t
}

var _ Property = &Talent{}

func (t Talents) String() string {
	msg := ""
	for _, talent := range t {
		msg += talent.Name + "\n"
	}
	return msg
}
func (t Talents) Value() interface{} {
	return t
}

var _ Property = &Talents{}
