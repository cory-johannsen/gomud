package loader

import (
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type TaskLoader struct {
	config          *config.Config
	specs           htn.TaskSpecs
	resolvers       htn.TaskResolvers
	tasks           htn.Tasks
	actionLoader    *ActionLoader
	conditionLoader *ConditionLoader
	methodLoader    *MethodLoader
}

func NewTaskLoader(cfg *config.Config, actionLoader *ActionLoader, conditionLoader *ConditionLoader, methodLoader *MethodLoader) *TaskLoader {
	return &TaskLoader{
		config:          cfg,
		specs:           make(htn.TaskSpecs),
		resolvers:       make(htn.TaskResolvers),
		tasks:           make(htn.Tasks),
		actionLoader:    actionLoader,
		conditionLoader: conditionLoader,
		methodLoader:    methodLoader,
	}
}

func (l *TaskLoader) LoadTaskResolvers() (htn.TaskResolvers, error) {
	if len(l.resolvers) > 0 {
		return l.resolvers, nil
	}

	primitive, err := l.loadPrimitiveSpecs()
	if err != nil {
		return nil, err
	}
	for k, v := range primitive {
		l.specs[k] = v
		resolver := func() (htn.Task, error) {
			existing, ok := l.tasks[k]
			if ok {
				return existing, nil
			}
			log.Printf("instantiating primitive task %s", k)
			t, err := l.loadPrimitiveTask(v)
			if err != nil {
				return nil, err
			}
			l.tasks[k] = t
			return t, nil
		}
		l.resolvers[k] = resolver
	}

	compound, err := l.loadCompoundSpecs()
	if err != nil {
		return nil, err
	}
	for k, v := range compound {
		l.specs[k] = v
		resolver := func() (htn.Task, error) {
			existing, ok := l.tasks[k]
			if ok {
				return existing, nil
			}
			log.Printf("instantiating compound task %s", k)
			t, err := l.loadCompoundTask(v)
			if err != nil {
				return nil, err
			}
			l.tasks[k] = t
			return t, nil
		}
		l.resolvers[k] = resolver
	}
	return l.resolvers, nil
}

func (l *TaskLoader) loadPrimitiveSpecs() (htn.TaskSpecs, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/tasks/primitive")
	if err != nil {
		return nil, err
	}
	tasks := make(htn.TaskSpecs)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}

		spec := &htn.TaskSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/tasks/primitive/" + item.Name())
		if err != nil {
			log.Errorf("error reading task spec file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling task spec file %s: %v", item.Name(), err)
			continue
		}
		spec.TaskType = htn.PrimitiveTaskType
		tasks[spec.TaskName] = spec
	}
	return tasks, nil
}

func (l *TaskLoader) loadCompoundSpecs() (htn.TaskSpecs, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/tasks/compound")
	if err != nil {
		return nil, err
	}
	tasks := make(htn.TaskSpecs)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}

		spec := &htn.TaskSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/tasks/compound/" + item.Name())
		if err != nil {
			log.Errorf("error reading task spec file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling task spec file %s: %v", item.Name(), err)
			continue
		}
		spec.TaskType = htn.CompoundTaskType
		tasks[spec.TaskName] = spec
	}
	return tasks, nil
}

func (l *TaskLoader) loadGoalSpecs() (htn.TaskSpecs, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/tasks/goal")
	if err != nil {
		return nil, err
	}
	tasks := make(htn.TaskSpecs)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}

		spec := &htn.TaskSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/tasks/goal/" + item.Name())
		if err != nil {
			log.Errorf("error reading task spec file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling task spec file %s: %v", item.Name(), err)
			continue
		}
		spec.TaskType = htn.GoalTaskType
		tasks[spec.TaskName] = spec
	}
	return tasks, nil
}

func (l *TaskLoader) loadPrimitiveTask(spec *htn.TaskSpec) (htn.Task, error) {
	conditions := make([]htn.Condition, 0)
	for _, conditionName := range spec.Preconditions {
		condition, err := l.conditionLoader.GetCondition(conditionName)
		if err != nil {
			return nil, err
		}
		if condition == nil {
			log.Errorf("condition %s not found", conditionName)
			continue
		}
		conditions = append(conditions, condition)
	}
	action, err := l.actionLoader.GetAction(spec.Action)
	if err != nil {
		return nil, err
	}
	return &htn.PrimitiveTask{
		Preconditions: conditions,
		Complete:      spec.Complete,
		Action:        action,
		TaskName:      spec.TaskName,
	}, nil
}

func (l *TaskLoader) loadCompoundTask(spec *htn.TaskSpec) (htn.Task, error) {
	methods := make([]*htn.Method, 0)
	for _, methodName := range spec.Preconditions {
		method, err := l.methodLoader.GetMethod(methodName, l)
		if err != nil {
			return nil, err
		}
		if method == nil {
			log.Errorf("method %s not found", methodName)
			continue
		}
		methods = append(methods, method)
	}
	return &htn.CompoundTask{
		Methods:  methods,
		Complete: spec.Complete,
		TaskName: spec.TaskName,
	}, nil
}

func (l *TaskLoader) loadGoalTask(spec *htn.TaskSpec) (htn.Task, error) {
	taskConditions := make([]*htn.TaskCondition, 0)
	for _, taskName := range spec.Preconditions {
		resolver, ok := l.resolvers[taskName]
		if !ok {
			return nil, errors.New(fmt.Sprintf("task resolver not found %s", taskName))
		}
		task, err := resolver()
		if err != nil {
			return nil, err
		}
		taskConditions = append(taskConditions, &htn.TaskCondition{
			Task: task,
		})
	}
	return &htn.GoalTask{
		Preconditions: taskConditions,
		Complete:      spec.Complete,
		TaskName:      spec.TaskName,
	}, nil
}

func (l *TaskLoader) GetTaskResolver(name string) (htn.TaskResolver, error) {
	if len(l.resolvers) == 0 {
		_, err := l.LoadTaskResolvers()
		if err != nil {
			return nil, err
		}
	}
	if task, ok := l.resolvers[name]; ok {
		return task, nil
	}
	return nil, nil
}
