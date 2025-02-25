package generator

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type StateGenerator struct {
	states htn.States
}

func (s StateGenerator) Get(name string) (*htn.State, error) {
	// TODO
	return &htn.State{}, nil
}

var _ htn.StateResolver = &StateGenerator{}

func NewStateGenerator() htn.StateResolver {
	return &StateGenerator{
		states: make(htn.States),
	}
}
