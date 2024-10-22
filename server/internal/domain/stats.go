package domain

import (
	"fmt"
)

type Stats struct {
	Fighting int `json:"fighting"`
	Muscle   int `json:"muscle"`
	Speed    int `json:"speed"`
	Savvy    int `json:"savvy"`
	Smarts   int `json:"smarts"`
	Grit     int `json:"grit"`
	Flair    int `json:"flair"`
}

func (s *Stats) Value() interface{} {
	return s
}

func (s *Stats) String() string {
	return fmt.Sprintf("Fighting: %d\nMuscle: %d\nSpeed: %d\nSavvy: %d\nSmarts: %d\nGrit: %d\nFlair: %d\n", s.Fighting, s.Muscle, s.Speed, s.Savvy, s.Smarts, s.Grit, s.Flair)
}

func (s *Stats) StatValue(name string) int {
	switch name {
	case "Fighting":
		return s.Fighting
	case "Muscle":
		return s.Muscle
	case "Speed":
		return s.Speed
	case "Savvy":
		return s.Savvy
	case "Smarts":
		return s.Smarts
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
		Fighting: ThreeD10() + 25,
		Muscle:   ThreeD10() + 25,
		Speed:    ThreeD10() + 25,
		Savvy:    ThreeD10() + 25,
		Smarts:   ThreeD10() + 25,
		Grit:     ThreeD10() + 25,
		Flair:    ThreeD10() + 25,
	}
}
