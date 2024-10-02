package domain

type Season string

const (
	Spring Season = "Spring"
	Summer Season = "Summer"
	Fall   Season = "Fall"
	Winter Season = "Winter"
)

type DistinguishingMark string
type DistinguishingMarks []DistinguishingMark

type Tattoo struct {
	Description string
	Season      Season
}
type Tattoos []Tattoo

type Drawback struct {
	Name        string
	Description string
	Effect      string
}
type Drawbacks []Drawback
