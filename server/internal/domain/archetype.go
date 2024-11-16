package domain

import "math/rand"

type StartingEquipmentSpec struct {
	OneEach []string `yaml:"oneEach"`
	OneOf   []string `yaml:"oneOf"`
}

type StartingEquipment struct {
	OneEach Items
	OneOf   Items
}

type ArchetypeSpec struct {
	Name              string                `yaml:"name"`
	Description       string                `yaml:"description"`
	StartingEquipment StartingEquipmentSpec `yaml:"startingEquipment"`
	Traits            []string              `yaml:"traits"`
}

type Archetype struct {
	Name              string
	Description       string
	StartingEquipment StartingEquipment
	Traits            Traits
}

type Archetypes []*Archetype

func (a *Archetype) String() string {
	return a.Name
}

func (a *Archetype) Value() interface{} {
	return a
}

var _ Property = &Archetype{}

func (a Archetypes) Random() *Archetype {
	if len(a) == 0 {
		return nil
	}
	return a[rand.Intn(len(a))]
}
