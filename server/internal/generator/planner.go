package generator

import "github.com/cory-johannsen/gomud/internal/domain/htn"

type PlannerGenerator struct {
	planners htn.Planners
}

func (p *PlannerGenerator) GetPlanner(name string) (*htn.Planner, error) {
	if planner, ok := p.planners[name]; ok {
		return planner, nil
	}
	return nil, nil
}

func (p *PlannerGenerator) Register(planner *htn.Planner) {
	p.planners[planner.Name] = planner
}

func (p *PlannerGenerator) Deregister(planner *htn.Planner) error {
	delete(p.planners, planner.Name)
	return nil
}

var _ htn.PlannerResolver = &PlannerGenerator{}

func NewPlannerGenerator() htn.PlannerResolver {
	return &PlannerGenerator{
		planners: make(htn.Planners),
	}
}
