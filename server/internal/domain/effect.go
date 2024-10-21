package domain

type Applier func(state State) State

type Effect struct {
	Name        string
	Description string
	Applier     Applier
}
type Effects []*Effect
