package domain

type Applier func()

type Effect struct {
	Name    string
	Applier Applier
}
type Effects []*Effect

type Trait struct {
	Name        string
	Description string
	Effects     Effects
}
type Traits []*Trait
