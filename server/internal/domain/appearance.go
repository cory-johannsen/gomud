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

type Dooming struct {
	Description string
	Season      Season
}
type Doomings []Dooming

type Drawback struct {
	Name        string
	Description string
	Effect      string
}
type Drawbacks []Drawback
