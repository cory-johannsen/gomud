package htn

import (
	log "github.com/sirupsen/logrus"
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

type TaskGraph struct {
	Name string
	Root *TaskNode
}
type TaskGraphs map[string]*TaskGraph

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

func evaluateNode(node *TaskNode, state *State) []Task {
	task, err := node.TaskResolver()
	if err != nil {
		panic(err)
	}
	log.Debugf("evaluating task node {%s}", task.Name())
	tasks := make([]Task, 0)
	if !task.IsComplete() {
		log.Debugf("task node {%s} is not complete, adding to tasks", task.Name())
		tasks = append(tasks, task)
	}
	log.Debugf("evaluating task node {%s} children", task.Name())
	for _, child := range node.Children {
		childTasks := evaluateNode(child, state)
		for _, childTask := range childTasks {
			tasks = append([]Task{childTask}, tasks...)
		}
	}
	return tasks
}

func (p *Planner) Plan(state *State) (*Plan, error) {
	log.Debugf("building plan %s", p.Name)
	plan := &Plan{
		Name:  p.Name,
		Tasks: make(PrioritizedTasks, 0),
	}
	// walk the Task graph, starting at the root, and find the executable plan
	node := p.Tasks.Root
	if node != nil {
		tasks := evaluateNode(node, state)
		for _, task := range tasks {
			plan.Tasks = append(plan.Tasks, task)
		}
	}
	log.Debugf("plan %s contains %d tasks", plan.Name, len(plan.Tasks))
	return plan, nil
}

func Execute(plan *Plan, state *State) (*State, error) {
	log.Debugf("executing plan %s with %d tasks", plan.Name, len(plan.Tasks))
	for _, task := range plan.Tasks {
		log.Debugf("executing task %s", task.Name())
		_, err := task.Execute(state)
		if err != nil {
			return nil, err
		}
	}
	return state, nil
}
