package domain

import (
	"fmt"
	"math/rand"
)

type Order struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effect"`
	Rank        int
}

type Chaos struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effect      []string `yaml:"effect"`
	Rank        int
}

type AlignmentSpec struct {
	Order     string `json:"order"`
	OrderRank int    `json:"orderRank"`
	Chaos     string `json:"chaos"`
	ChaosRank int    `json:"chaosRank"`
	Rank      int    `json:"rank"`
}

type Alignment struct {
	Order      *Order
	Chaos      *Chaos
	Corruption int
}

func (a *Alignment) AddOrderRank(rank int) {
	a.Order.Rank += rank
}

func (a *Alignment) ResetOrderRank() {
	a.Order.Rank = 0
}

func (a *Alignment) AddChaosRank(rank int) {
	a.Chaos.Rank += rank
}

func (a *Alignment) ResetChaosRank() {
	a.Chaos.Rank = 0
}

func (a *Alignment) AddCorruption(corruption int) {
	a.Corruption += corruption
	if a.Corruption >= 10 {
		a.Chaos.Rank++
		a.Corruption = 0
	}
}

func (a *Alignment) ResetCorruption() {
	a.Corruption = 0
}

func (a *Alignment) Value() interface{} {
	return a
}

func (a *Alignment) String() string {
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
		Order:     a.Order.Name,
		OrderRank: a.Order.Rank,
		Chaos:     a.Chaos.Name,
		ChaosRank: a.Chaos.Rank,
		Rank:      a.Corruption,
	}
}
