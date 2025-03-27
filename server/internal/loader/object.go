package loader

import (
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type InteractiveObjectLoader struct {
	config  *config.Config
	objects domain.InteractiveObjects
}

func NewInteractiveObjectLoader(cfg *config.Config) *InteractiveObjectLoader {
	return &InteractiveObjectLoader{
		config:  cfg,
		objects: make(domain.InteractiveObjects),
	}
}

func (l *InteractiveObjectLoader) LoadInteractiveObjects() (domain.InteractiveObjects, error) {
	if len(l.objects) > 0 {
		return l.objects, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/objects")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") || strings.HasSuffix(item.Name(), "txt") {
			continue
		}
		name := item.Name()
		data, err := os.ReadFile(l.config.AssetPath + "/objects/" + name)
		if err != nil {
			log.Printf("error reading file %s: %s", name, err)
			continue
		}
		spec := &domain.InteractiveObjectSpec{}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Printf("error unmarshalling file %s: %s", name, err)
			continue
		}

	}
	return l.objects, nil
}

func (l *InteractiveObjectLoader) GetInteractiveObject(name string) (domain.InteractiveObject, error) {
	objs, err := l.LoadInteractiveObjects()
	if err != nil {
		return nil, err
	}
	obj, ok := objs[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("interactive object %s not found", name))
	}
	return obj, nil
}
