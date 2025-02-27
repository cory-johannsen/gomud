package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
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

func (s *StateGenerator) AddState(npc *domain.NPC, state *htn.State) {
	s.states[npc.Name] = state
}

func (s *StateGenerator) DeleteState(npc *domain.NPC) {
	delete(s.states, npc.Name)
}

var _ htn.StateResolver = &StateGenerator{}

func NewStateGenerator() *StateGenerator {
	return &StateGenerator{
		states: make(htn.States),
	}
}
