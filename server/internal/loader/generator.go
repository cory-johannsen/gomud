package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type GeneratorLoader struct {
	config     *config.Config
	generators map[string]*domain.GeneratorSpec
	roomLoader *RoomLoader
	npcLoader  *NPCLoader
}

func NewGeneratorLoader(cfg *config.Config, roomLoader *RoomLoader, npcLoader *NPCLoader) *GeneratorLoader {
	return &GeneratorLoader{
		config:     cfg,
		generators: make(map[string]*domain.GeneratorSpec),
		roomLoader: roomLoader,
		npcLoader:  npcLoader,
	}
}

func (l *GeneratorLoader) LoadGenerators() (map[string]*domain.GeneratorSpec, error) {
	if len(l.generators) > 0 {
		return l.generators, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/generators")
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
		spec := &domain.GeneratorSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/generators/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}

		l.generators[spec.Name] = spec
	}
	return l.generators, nil
}
