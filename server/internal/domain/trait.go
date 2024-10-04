package domain

type Applier func()

type Effect struct {
	Name    string
	Applier Applier
}
type Effects []*Effect

type TraitSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effects"`
}

type Trait struct {
	Name        string
	Description string
	Effects     Effects
}
type Traits []*Trait
