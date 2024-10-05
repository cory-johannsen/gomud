package domain

import (
	"fmt"
	"math/rand"
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

var _ Property = &Stats{}

func NewStats() *Stats {
	return &Stats{
		Fighting: threeD10() + 25,
		Muscle:   threeD10() + 25,
		Speed:    threeD10() + 25,
		Savvy:    threeD10() + 25,
		Smarts:   threeD10() + 25,
		Grit:     threeD10() + 25,
		Flair:    threeD10() + 25,
	}
}

func threeD10() int {
	return rand.Intn(10) + rand.Intn(10) + rand.Intn(10)
}
