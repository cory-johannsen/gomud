package domain

import (
	"fmt"
)

type Stats struct {
	Brutality int `json:"brutality"`
	Muscle    int `json:"muscle"`
	Quickness int `json:"quickness"`
	Savvy     int `json:"savvy"`
	Reasoning int `json:"reasoning"`
	Grit      int `json:"grit"`
	Flair     int `json:"flair"`
}

func (s *Stats) Value() interface{} {
	return s
}

func (s *Stats) String() string {
	return fmt.Sprintf("Brutality: %d\nMuscle: %d\nQuickness: %d\nSavvy: %d\nReasoning: %d\nGrit: %d\nFlair: %d\n", s.Brutality, s.Muscle, s.Quickness, s.Savvy, s.Reasoning, s.Grit, s.Flair)
}

func (s *Stats) StatValue(name string) int {
	switch name {
	case "Brutality":
		return s.Brutality
	case "Muscle":
		return s.Muscle
	case "Quickness":
		return s.Quickness
	case "Savvy":
		return s.Savvy
	case "Reasoning":
		return s.Reasoning
	case "Grit":
		return s.Grit
	case "Flair":
		return s.Flair
	}
	return -1
}

func (s *Stats) StatBonus(name string) int {
	val := s.StatValue(name)
	return val / 10
}

var _ Property = &Stats{}

func NewStats() *Stats {
	return &Stats{
		Brutality: ThreeD10() + 25,
		Muscle:    ThreeD10() + 25,
		Quickness: ThreeD10() + 25,
		Savvy:     ThreeD10() + 25,
		Reasoning: ThreeD10() + 25,
		Grit:      ThreeD10() + 25,
		Flair:     ThreeD10() + 25,
	}
}
