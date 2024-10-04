package domain

type Tier int

const (
	Tier0 Tier = iota
	Tier1
	Tier2
	Tier3
)

type Job struct {
	Name           string
	Description    string
	Archetype      *Archetype
	Tier           Tier
	ExperienceCost int
	Trait          *Trait
}
