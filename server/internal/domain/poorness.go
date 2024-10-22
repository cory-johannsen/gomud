package domain

import "math/rand"

type Poorness string

func (p Poorness) Value() interface{} {
	return p
}

func (p Poorness) String() string {
	return string(p)
}

const (
	PoornessDestitute   Poorness = "destitute"
	PoornessPoor        Poorness = "poor"
	PoornessMiddleClass Poorness = "middle class"
	PoornessWealthy     Poorness = "wealthy"
	PoornessRich        Poorness = "rich"
)

func (p Poorness) Rank() int {
	for i, item := range RankedPoorness {
		if item == p {
			return i
		}
	}
	return -1
}

func (p Poorness) PoorerThan(other Poorness) bool {
	return p.Rank() < other.Rank()
}

var RankedPoorness = []Poorness{
	PoornessDestitute,
	PoornessPoor,
	PoornessMiddleClass,
	PoornessWealthy,
	PoornessRich,
}

func RandomPoorness() Poorness {
	roll := rand.Intn(100)
	if roll < 50 {
		return PoornessDestitute
	}
	if roll < 75 {
		return PoornessPoor
	}
	if roll < 90 {
		return PoornessMiddleClass
	}
	if roll < 98 {
		return PoornessWealthy
	}
	return PoornessRich
}
