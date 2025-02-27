package htn

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Task interface {
	Execute(state *State) (*State, error)
	IsComplete() bool
	Name() string
	String() string
}

type Tasks map[string]Task
type TaskResolver func() (Task, error)
type TaskResolvers map[string]TaskResolver

type TaskType string

const (
	PrimitiveTaskType TaskType = "primitive"
	CompoundTaskType  TaskType = "compound"
	GoalTaskType      TaskType = "goal"
)

type TaskSpec struct {
	Preconditions []string `yaml:"preconditions"`
	Complete      bool     `yaml:"complete,omitempty"`
	Action        string   `yaml:"action,omitempty"`
	TaskName      string   `yaml:"name"`
	TaskType      TaskType `yaml:"type,omitempty"`
}

type TaskSpecs map[string]*TaskSpec

// Action is an action applied by a Task.
type Action func(state *State) error
type Actions map[string]Action

// PrimitiveTask implements the HTN primitive Task.   It contains a set of preconditions that must be met
// before it will execute.  Once the preconditions are met, the Action is applied, then the completion flag is set.
type PrimitiveTask struct {
	Preconditions []Condition `yaml:"preconditions"`
	Complete      bool        `yaml:"complete"`
	Action        Action      `yaml:"action"`
	TaskName      string      `yaml:"name"`
}

func (t *PrimitiveTask) Execute(state *State) (*State, error) {
	preconditions := make([]string, 0)
	for _, condition := range t.Preconditions {
		preconditions = append(preconditions, condition.String())
	}
	log.Debugf("executing task {%s}, preconditions {%s}", t.Name(), strings.Join(preconditions, ","))
	// Determine if the Task preconditions have been met
	var ready = true
	for _, condition := range t.Preconditions {
		log.Debugf("evaluating condition {%s}", condition.String())
		if !condition.IsMet(state) {
			ready = false
			break
		}
	}
	if ready {
		log.Printf("task {%s} preconditions met, applying task action", t.Name())
		// Apply the Task action and update the state
		err := t.Action(state)
		if err != nil {
			return nil, err
		}
		// Set the Task to 'complete' so it doesn't execute again
		t.Complete = true
	}
	return state, nil
}

func (t *PrimitiveTask) IsComplete() bool {
	return t.Complete
}

func (t *PrimitiveTask) Name() string {
	return t.TaskName
}

func (t *PrimitiveTask) String() string {
	preconditions := make([]string, 0)
	for _, condition := range t.Preconditions {
		preconditions = append(preconditions, condition.String())
	}
	return fmt.Sprintf("[%s] preconditions: [%s], complete: %t", t.Name(), strings.Join(preconditions, ","), t.Complete)
}

// GoalTask implements the HTN goal Task, composed of preconditions that are other TaskResolvers.  The goal Task is considered
// complete when all condition TaskResolvers are themselves complete.
type GoalTask struct {
	Preconditions []*TaskCondition `yaml:"preconditions"`
	Complete      bool             `yaml:"complete"`
	TaskName      string           `yaml:"name"`
}

func (g *GoalTask) Execute(state *State) (*State, error) {
	log.Printf("executing goal task %s", g.TaskName)
	if !g.Complete {
		log.Printf("goal task %s is not complete checking preconditions", g.TaskName)
		for _, condition := range g.Preconditions {
			if !condition.IsMet(state) {
				log.Printf("goal task %s precondition not met, exiting", g.TaskName)
				return state, nil
			}
		}
		log.Printf("goal %s conditions met, goal Task is complete.", g.TaskName)
		g.Complete = true
	}
	return state, nil
}

func (g *GoalTask) IsComplete() bool {
	return g.Complete
}

func (g *GoalTask) Name() string {
	return g.TaskName
}

func (g *GoalTask) String() string {
	preconditions := make([]string, 0)
	for _, condition := range g.Preconditions {
		preconditions = append(preconditions, fmt.Sprintf("{%s}", condition.String()))
	}
	return fmt.Sprintf("goal: preconditions: [%s], complete: %t", strings.Join(preconditions, ","), g.Complete)
}

type MethodSpec struct {
	Name       string   `yaml:"name"`
	Conditions []string `yaml:"conditions"`
	Tasks      []string `yaml:"tasks"`
}

type Method struct {
	Conditions    []Condition
	TaskResolvers TaskResolvers
	Name          string
}

type Methods map[string]*Method

func (m *Method) Applies(state *State) bool {
	log.Printf("checking if method {%s} applies", m.Name)
	for _, condition := range m.Conditions {
		if !condition.IsMet(state) {
			log.Printf("method {%s} condition {%s} not met, exiting", m.Name, condition.Name())
			return false
		}
	}
	return true
}

func (m *Method) Execute(state *State) (int64, error) {
	log.Printf("executing method {%s}", m.Name)
	var executed = int64(0)
	tasks := make([]Task, 0)
	for _, taskResolver := range m.TaskResolvers {
		task, err := taskResolver()
		if err != nil {
			return 0, err
		}
		tasks = append([]Task{task}, tasks...)
	}
	for _, task := range tasks {
		if !task.IsComplete() {
			log.Printf("method {%s} task {%s} not complete, executing it", m.Name, task.Name())
			_, err := task.Execute(state)
			if err != nil {
				return -1, err
			}
			executed++
		}
	}
	return executed, nil
}

func (m *Method) String() string {
	conditions := make([]string, 0)
	for _, condition := range m.Conditions {
		conditions = append(conditions, fmt.Sprintf("{%s}", condition.String()))
	}
	tasks := make([]string, 0)
	for taskName := range m.TaskResolvers {
		tasks = append(tasks, fmt.Sprintf("{%s}", taskName))
	}
	return fmt.Sprintf("Method %s:\n  conditions: \n   %s\n  tasks: \n   %s", m.Name, strings.Join(conditions, ",\n   "), strings.Join(tasks, ",\n   "))
}

// CompoundTask implements the HTN compound task, which consists of a ranked list of methods and a name.
// The task selects a method at execution time by checking the conditions on each.  Since the method list
// is in priority order, the first match is selected when more than one apply.
type CompoundTask struct {
	Methods  []*Method `yaml:"methods"`
	TaskName string    `yaml:"name"`
	Complete bool      `yaml:"complete"`
}

func (c *CompoundTask) Execute(state *State) (*State, error) {
	log.Printf("executing compound task {%s}", c.Name())
	applicableMethods := make([]*Method, 0)
	for _, method := range c.Methods {
		if method.Applies(state) {
			applicableMethods = append(applicableMethods, method)
		}
	}
	if len(applicableMethods) == 0 {
		log.Println("no applicable methods found")
		c.Complete = true
		return state, nil
	}
	// The methods are stored in priority order, so the first one is the selected choice
	selectedMethod := applicableMethods[0]
	executedTasks, err := selectedMethod.Execute(state)
	if err != nil {
		return nil, err
	}
	if executedTasks == 0 {
		log.Printf("method {%s} execute zero tasks", c.Name())
	}

	return state, nil
}

func (c *CompoundTask) Name() string {
	return c.TaskName
}

func (c *CompoundTask) IsComplete() bool {
	return c.Complete
}

func (c *CompoundTask) String() string {
	methods := make([]string, 0)
	for _, method := range c.Methods {
		methods = append(methods, fmt.Sprintf("{%s}", method.String()))
	}
	return fmt.Sprintf("CompoundTask %s: methods: \n %s\n", c.Name(), strings.Join(methods, ",\n "))
}
