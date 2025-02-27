package loader

import (
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type MethodLoader struct {
	config          *config.Config
	methods         htn.Methods
	conditionLoader *ConditionLoader
}

func NewMethodLoader(cfg *config.Config, conditionLoader *ConditionLoader) *MethodLoader {
	return &MethodLoader{
		config:          cfg,
		methods:         make(htn.Methods),
		conditionLoader: conditionLoader,
	}
}

func (l *MethodLoader) LoadMethods(taskLoader *TaskLoader) (htn.Methods, error) {
	if len(l.methods) > 0 {
		return l.methods, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/htn/methods")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}

		spec := &htn.MethodSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/methods/" + item.Name())
		if err != nil {
			log.Errorf("error reading method spec file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling method spec file %s: %v", item.Name(), err)
			continue
		}

		conditions := make([]htn.Condition, 0)
		for _, conditionName := range spec.Conditions {
			condition, err := l.conditionLoader.GetCondition(conditionName)
			if err != nil {
				return nil, fmt.Errorf("error fetching condition: %s - %s", conditionName, err)
			}
			if condition != nil {
				conditions = append(conditions, condition)
			} else {
				log.Errorf("condition %s not found loading method %s", conditionName, spec.Name)
			}
		}
		resolvers := make(htn.TaskResolvers)
		for _, taskName := range spec.Conditions {
			resolver, err := taskLoader.GetTaskResolver(taskName)
			if err != nil {
				return nil, err
			}
			resolvers[taskName] = resolver
		}

		l.methods[spec.Name] = &htn.Method{
			Conditions:    conditions,
			TaskResolvers: resolvers,
			Name:          spec.Name,
		}
	}

	return l.methods, nil
}

func (l *MethodLoader) GetMethod(name string, taskLoader *TaskLoader) (*htn.Method, error) {
	if len(l.methods) == 0 {
		_, err := l.LoadMethods(taskLoader)
		if err != nil {
			return nil, err
		}
	}
	if method, ok := l.methods[name]; ok {
		return method, nil
	}
	return nil, nil
}
