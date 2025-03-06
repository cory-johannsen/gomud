package loader

import (
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync"
)

type TaskGraphLoader struct {
	mutex      sync.Mutex
	config     *config.Config
	taskGraphs htn.TaskGraphs
	taskLoader *TaskLoader
}

func NewTaskGraphLoader(cfg *config.Config, taskLoader *TaskLoader) *TaskGraphLoader {
	return &TaskGraphLoader{
		config:     cfg,
		taskGraphs: make(htn.TaskGraphs),
		taskLoader: taskLoader,
	}
}

func (l *TaskGraphLoader) LoadTaskGraphs() (htn.TaskGraphs, error) {
	items, err := os.ReadDir(l.config.AssetPath + "/htn/domain")
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
		spec := &htn.TaskGraphSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/htn/domain/" + item.Name())
		if err != nil {
			log.Errorf("error reading task graph file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling task graph file %s: %v", item.Name(), err)
			continue
		}
		resolvers, err := l.taskLoader.LoadTaskResolvers()
		if err != nil {
			return nil, err
		}
		root, err := loadTaskNode(spec.Root, resolvers)
		l.mutex.Lock()
		l.taskGraphs[spec.Name] = &htn.TaskGraph{
			Name: spec.Name,
			Root: root,
		}
		l.mutex.Unlock()
	}
	return l.taskGraphs, nil
}

func (l *TaskGraphLoader) GetTaskGraph(name string) (*htn.TaskGraph, error) {
	taskGraphs, err := l.LoadTaskGraphs()
	if err != nil {
		return nil, err
	}
	if taskGraph, ok := taskGraphs[name]; ok {
		return taskGraph, nil
	}
	return nil, nil
}

func loadTaskNode(spec *htn.TaskNodeSpec, resolvers htn.TaskResolvers) (*htn.TaskNode, error) {
	taskResolver, ok := resolvers[spec.Task]
	if !ok {
		return nil, fmt.Errorf("taskResolver %s not found", spec.Task)
	}
	children := make([]*htn.TaskNode, 0)
	for _, childSpec := range spec.Children {
		child, err := loadTaskNode(childSpec, resolvers)
		if err != nil {
			return nil, err
		}
		children = append(children, child)
	}
	node := &htn.TaskNode{
		TaskResolver: taskResolver,
		Children:     children,
	}
	return node, nil
}
