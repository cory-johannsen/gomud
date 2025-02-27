package domain

import (
	"fmt"
	"github.com/fatih/color"
)

const (
	Brutality = "Brutality"
	Muscle    = "Muscle"
	Quickness = "Quickness"
	Savvy     = "Savvy"
	Reasoning = "Reasoning"
	Grit      = "Grit"
	Flair     = "Flair"
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

func (s *Stats) String() string {
	return fmt.Sprintf("Brutality: %d\nMuscle: %d\nQuickness: %d\nSavvy: %d\nReasoning: %d\nGrit: %d\nFlair: %d\n", s.Brutality, s.Muscle, s.Quickness, s.Savvy, s.Reasoning, s.Grit, s.Flair)
}

func (s *Stats) StatValue(name string) int {
	switch name {
	case Brutality:
		return s.Brutality
	case Muscle:
		return s.Muscle
	case Quickness:
		return s.Quickness
	case Savvy:
		return s.Savvy
	case Reasoning:
		return s.Reasoning
	case Grit:
		return s.Grit
	case Flair:
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

func statBonusAsString(bonuses *Stats, advances ConsumedAdvances, stat string) string {
	bonus := bonuses.StatValue(stat)
	advance := 0
	for _, jobAdvances := range advances {
		for _, a := range jobAdvances {
			if a.Stat == stat {
				advance += a.Amount
			}
		}
	}
	if advance > 0 {
		return color.New(color.FgGreen).Sprintf("%d", bonus+advance)
	}
	return color.New(color.FgCyan).Sprintf("%d", bonus+advance)
}

func StatsString(stats *Stats, bonuses *Stats, advances ConsumedAdvances) string {
	bb := statBonusAsString(bonuses, advances, Brutality)
	bm := statBonusAsString(bonuses, advances, Muscle)
	bq := statBonusAsString(bonuses, advances, Quickness)
	bs := statBonusAsString(bonuses, advances, Savvy)
	br := statBonusAsString(bonuses, advances, Reasoning)
	bg := statBonusAsString(bonuses, advances, Grit)
	bf := statBonusAsString(bonuses, advances, Flair)

	cyan := color.New(color.FgCyan).SprintFunc()

	return fmt.Sprintf("%s\n\tBrutality: %s\tBB[%s]\n\tMuscle:    %s\tMB[%s]\n\tQuickness: %s\tQB[%s]\n\tSavvy:     %s\tSB[%s]\n\tReasoning: %s\tRB[%s]\n\tGrit:      %s\tGB[%s]\n\tFlair:     %s\tRB[%s]\n",
		cyan("Stats"),
		cyan(stats.Brutality), bb,
		cyan(stats.Muscle), bm,
		cyan(stats.Quickness), bq,
		cyan(stats.Savvy), bs,
		cyan(stats.Reasoning), br,
		cyan(stats.Grit), bg,
		cyan(stats.Flair), bf)
}
