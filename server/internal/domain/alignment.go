package domain

import (
	"fmt"
	"math/rand"
)

type Order struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effects"`
	Rank        int
}

type Chaos struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effect      []string `yaml:"effects"`
	Rank        int
}

type AlignmentSpec struct {
	Order string `json:"order"`
	Chaos string `json:"chaos"`
	Rank  int    `json:"rank"`
}

type Alignment struct {
	Order      *Order
	Chaos      *Chaos
	Corruption int
}

func (a Alignment) Value() interface{} {
	return a
}

func (a Alignment) String() string {
	return fmt.Sprintf("%s - %s", a.Order.Name, a.Chaos.Name)
}

type Alignments []Alignment

func (a Alignments) Random() Alignment {
	if len(a) == 0 {
		return Alignment{}
	}
	return a[rand.Intn(len(a))]
}

func SpecFromAlignment(a *Alignment) *AlignmentSpec {
	return &AlignmentSpec{
		Order: a.Order.Name,
		Chaos: a.Chaos.Name,
		Rank:  a.Corruption,
	}
}
