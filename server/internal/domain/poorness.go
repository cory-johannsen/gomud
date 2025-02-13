package domain

import "math/rand"

const CIGS_PER_PACK = 20
const PACKS_PER_CASE = 24

type Poorness string

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

func (p Poorness) StartingCash() int {
	switch p {
	case PoornessDestitute:
		return ThreeD10() + 3
	case PoornessPoor:
		return (TwoD10() + 2) * CIGS_PER_PACK
	case PoornessMiddleClass:
		return (D10() + 1) * CIGS_PER_PACK * PACKS_PER_CASE
	case PoornessWealthy:
		return (TwoD10() + 2) * CIGS_PER_PACK * PACKS_PER_CASE
	case PoornessRich:
		return (ThreeD10() + 3) * CIGS_PER_PACK * PACKS_PER_CASE
	}
	return 0
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
