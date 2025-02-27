package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain/htn"
)

type StateGenerator struct {
	states htn.States
}

func (s *StateGenerator) GetState(name string) (*htn.State, error) {
	if state, ok := s.states[name]; ok {
		return state, nil
	}
	return nil, nil
}

func (s *StateGenerator) AddState(name string, state *htn.State) {
	s.states[name] = state
}

func (s *StateGenerator) DeleteState(name string) {
	delete(s.states, name)
}

var _ htn.StateResolver = &StateGenerator{}

func NewStateGenerator() *StateGenerator {
	return &StateGenerator{
		states: make(htn.States),
	}
}
