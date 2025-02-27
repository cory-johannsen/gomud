package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
)

type PlannerGenerator struct {
	planners htn.Planners
}

func (p *PlannerGenerator) GetPlanner(name string) (*htn.Planner, error) {
	if planner, ok := p.planners[name]; ok {
		return planner, nil
	}
	return nil, nil
}

func (p *PlannerGenerator) AddPlanner(npc *domain.NPC, planner *htn.Planner) {
	p.planners[npc.Name] = planner
}

func (p *PlannerGenerator) DeletePlanner(npc *domain.NPC) {
	delete(p.planners, npc.Name)
}

var _ htn.PlannerResolver = &PlannerGenerator{}

func NewPlannerGenerator() *PlannerGenerator {
	return &PlannerGenerator{
		planners: make(htn.Planners),
	}
}
