package domain

import (
	"fmt"
	"math/rand"
)

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

var _ Property = Season("")

type DistinguishingMark string
type DistinguishingMarks []DistinguishingMark

func (d DistinguishingMarks) String() string {
	msg := ""
	for i, mark := range d {
		if i > 0 {
			msg += ", "
		}
		msg += string(mark)
	}
	return msg
}

func (d DistinguishingMark) String() string {
	return string(d)
}

var _ Property = DistinguishingMark("")
var _ Property = DistinguishingMarks{}

func (d DistinguishingMarks) Random(age int) DistinguishingMarks {
	if len(d) == 0 {
		return []DistinguishingMark{}
	}
	if age < 25 {
		return []DistinguishingMark{}
	}
	if age < 35 {
		return []DistinguishingMark{d[rand.Intn(len(d))]}
	}
	if age < 50 {
		return []DistinguishingMark{d[rand.Intn(len(d))], d[rand.Intn(len(d))]}
	}
	return []DistinguishingMark{d[rand.Intn(len(d))], d[rand.Intn(len(d))], d[rand.Intn(len(d))]}
}

type TattooLocation string
type TattooLocations []TattooLocation

type Tattoo struct {
	Description string
	Location    TattooLocation
	Season      Season
}
type Tattoos []Tattoo
type SeasonalTattoos map[Season]Tattoos

func (t *Tattoo) String() string {
	return fmt.Sprintf("%s on the %s", t.Description, t.Location)
}

var _ Property = &Tattoo{}

func (t Tattoos) Random(locations TattooLocations) Tattoo {
	if len(t) == 0 {
		return Tattoo{}
	}
	tat := t[rand.Intn(len(t))]

	return Tattoo{
		Description: tat.Description,
		Location:    locations[rand.Intn(len(locations))],
		Season:      tat.Season,
	}
}

type DrawbackSpec struct {
	Name        string
	Description string
	Effect      string
}
type DrawbackSpecs []*DrawbackSpec

type Drawback struct {
	Name        string
	Description string
	Effect      Effect
}
type Drawbacks []*Drawback

func (d Drawback) String() string {
	return d.Name
}

func (d Drawbacks) Random() *Drawback {
	if len(d) == 0 {
		return nil
	}
	return d[rand.Intn(len(d))]

}
