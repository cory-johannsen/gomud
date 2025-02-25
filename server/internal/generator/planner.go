package generator

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type PlannerGenerator struct {
}

func (p PlannerGenerator) Get(name string) (*htn.Planner, error) {
	// TODO
	return &htn.Planner{
		Tasks: &htn.TaskGraph{
			Root: nil, // TODO
		},
	}, nil
}

var _ htn.PlannerResolver = &PlannerGenerator{}

func NewPlannerGenerator() htn.PlannerResolver {
	return &PlannerGenerator{}
}
