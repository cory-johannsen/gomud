package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"sync"
)

type PlannerGenerator struct {
	mutex    sync.Mutex
	planners htn.Planners
}

func (p *PlannerGenerator) GetPlanner(name string) (*htn.Planner, error) {
	if planner, ok := p.planners[name]; ok {
		return planner, nil
	}
	return nil, nil
}

func (p *PlannerGenerator) AddPlanner(name string, planner *htn.Planner) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.planners[name] = planner
}

func (p *PlannerGenerator) DeletePlanner(name string) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	delete(p.planners, name)
}

var _ htn.PlannerResolver = &PlannerGenerator{}

func NewPlannerGenerator() *PlannerGenerator {
	return &PlannerGenerator{
		planners: make(htn.Planners),
	}
}
