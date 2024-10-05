package domain

import "math/rand"

type Season string

const (
	Spring Season = "Spring"
	Summer Season = "Summer"
	Fall   Season = "Fall"
	Winter Season = "Winter"
)

func RandomSeason() Season {
	r := rand.Intn(4)
	switch r {
	case 0:
		return Spring
	case 1:
		return Summer
	case 2:
		return Fall
	case 3:
		return Winter
	}
	return Spring
}

func (s Season) String() string {
	return string(s)
}

func (s Season) Value() interface{} {
	return s
}

var _ Property = Season("")

type DistinguishingMark string
type DistinguishingMarks []DistinguishingMark

func (d DistinguishingMark) String() string {
	return string(d)
}

func (d DistinguishingMark) Value() interface{} {
	return d
}

var _ Property = DistinguishingMark("")

func (d DistinguishingMarks) Random() DistinguishingMark {
	if len(d) == 0 {
		return ""
	}
	return d[rand.Intn(len(d))]
}

type Tattoo struct {
	Description string
	Season      Season
}
type Tattoos []Tattoo
type SeasonalTattoos map[Season]Tattoos

func (t *Tattoo) String() string {
	return t.Description
}

func (t *Tattoo) Value() interface{} {
	return t
}

var _ Property = &Tattoo{}

func (t Tattoos) Random() Tattoo {
	if len(t) == 0 {
		return Tattoo{}
	}
	return t[rand.Intn(len(t))]
}

type Drawback struct {
	Name        string
	Description string
	Effect      string
}
type Drawbacks []*Drawback

func (d Drawback) String() string {
	return d.Name
}

func (d Drawback) Value() interface{} {
	return d
}

func (d Drawbacks) Random() *Drawback {
	if len(d) == 0 {
		return nil
	}
	return d[rand.Intn(len(d))]

}
