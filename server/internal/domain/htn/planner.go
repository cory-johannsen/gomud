package htn

import (
	log "github.com/sirupsen/logrus"
	"time"
)

type TaskNodeSpec struct {
	Task     string          `yaml:"task,omitempty"`
	Children []*TaskNodeSpec `yaml:"children,omitempty"`
}

type TaskGraphSpec struct {
	Name string        `yaml:"name"`
	Root *TaskNodeSpec `yaml:"root"`
}

type TaskNode struct {
	TaskResolver TaskResolver
	Children     []*TaskNode
}

func (t *TaskNode) Clone() *TaskNode {
	children := make([]*TaskNode, 0)
	if t.Children == nil {
		log.Errorf("clone task node with nil children")
	}
	for _, child := range t.Children {
		children = append(children, child.Clone())
	}
	return &TaskNode{
		TaskResolver: t.TaskResolver,
		Children:     children,
	}
}

type TaskGraph struct {
	Name string
	Root *TaskNode
}
type TaskGraphs map[string]*TaskGraph

func (t *TaskGraph) Clone() *TaskGraph {
	return &TaskGraph{
		Name: t.Name,
		Root: t.Root.Clone(),
	}
}

type PrioritizedTasks []Task

type Plan struct {
	Name  string
	Tasks PrioritizedTasks
}

type Planner struct {
	Name  string
	Tasks *TaskGraph
}

type Planners map[string]*Planner

type PlannerResolver interface {
	GetPlanner(name string) (*Planner, error)
}

func evaluateNode(node *TaskNode, domain *Domain) []Task {
	task, err := node.TaskResolver()
	if err != nil {
		panic(err)
	}
	log.Debugf("evaluating task node {%s}", task.Name())
	tasks := make([]Task, 0)
	if !task.IsComplete() {
		log.Debugf("task node {%s} is not complete, adding to tasks", task.Name())
		tasks = append(tasks, task)
	} else {
		log.Debugf("task node {%s} is complete, omitting", task.Name())
	}
	log.Debugf("evaluating task node {%s} children", task.Name())
	for _, child := range node.Children {
		childTasks := evaluateNode(child, domain)
		for _, childTask := range childTasks {
			tasks = append([]Task{childTask}, tasks...)
		}
	}
	return tasks
}

func (p *Planner) Plan(domain *Domain) (*Plan, error) {
	log.Debugf("building plan %s", p.Name)
	plan := &Plan{
		Name:  p.Name,
		Tasks: make(PrioritizedTasks, 0),
	}
	// clone the task graph and create new instances of the tasks
	// this is necessary because the tasks may have domain that is modified during execution.
	// we want to start with a clean slate each time we plan
	taskGraph := p.Tasks.Clone()

	// walk the Task graph, starting at the root, and find the executable plan
	node := taskGraph.Root
	if node != nil {
		tasks := evaluateNode(node, domain)
		for _, task := range tasks {
			plan.Tasks = append(plan.Tasks, task)
		}
	}
	log.Debugf("plan %s contains %d tasks", plan.Name, len(plan.Tasks))
	return plan, nil
}

func Execute(plan *Plan, domain *Domain, cutoff time.Duration) (*Domain, error) {
	log.Debugf("executing plan %s with %d tasks", plan.Name, len(plan.Tasks))
	startTime := time.Now()
	for _, task := range plan.Tasks {
		log.Debugf("executing task %s", task.Name())
		_, err := task.Execute(domain)
		if err != nil {
			return nil, err
		}
		if time.Now().Sub(startTime) > cutoff {
			break
		}
	}

	return domain, nil
}
