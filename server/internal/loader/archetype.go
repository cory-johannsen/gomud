package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type ArchetypeLoader struct {
	config      *config.Config
	archetypes  domain.Archetypes
	traitLoader *TraitLoader
}

func NewArchetypeLoader(cfg *config.Config, traitLoader *TraitLoader) *ArchetypeLoader {
	return &ArchetypeLoader{
		config:      cfg,
		archetypes:  make(domain.Archetypes, 0),
		traitLoader: traitLoader,
	}
}

func (l *ArchetypeLoader) LoadArchetypes() (domain.Archetypes, error) {
	if l.archetypes != nil && len(l.archetypes) > 0 {
		return l.archetypes, nil
	}
	archetypes := make(domain.Archetypes, 0)
	items, err := os.ReadDir(l.config.AssetPath + "/archetypes")
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
		spec := &domain.ArchetypeSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/archetypes/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}
		archetype := &domain.Archetype{
			Name:        spec.Name,
			Description: spec.Description,
			Traits:      make(domain.Traits, 0),
		}
		for _, traitName := range spec.Traits {
			trait, err := l.traitLoader.GetTrait(traitName)
			if err != nil {
				return nil, err
			}
			archetype.Traits = append(archetype.Traits, trait)
		}
		archetypes = append(archetypes, archetype)
	}
	return archetypes, nil
}

func (l *ArchetypeLoader) GetArchetype(name string) (*domain.Archetype, error) {
	archetypes, err := l.LoadArchetypes()
	if err != nil {
		return nil, err
	}

	for _, archetype := range archetypes {
		if archetype.Name == name {
			return archetype, nil
		}
	}
	return nil, nil
}
