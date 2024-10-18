package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type TraitLoader struct {
	config *config.Config
	traits domain.Traits
}

func NewTraitLoader(cfg *config.Config) *TraitLoader {
	return &TraitLoader{
		config: cfg,
		traits: make(domain.Traits, 0),
	}
}

func (l *TraitLoader) LoadTraits() (domain.Traits, error) {
	if l.traits != nil && len(l.traits) > 0 {
		return l.traits, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/traits")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			//log.Printf("skipping template file %s", item.Name())
			continue
		}
		//log.Printf("loading trait %s", item.Name())
		spec := &domain.TraitSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/traits/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}
		trait := domain.Trait{
			Name:        spec.Name,
			Description: spec.Description,
			Effects:     make(domain.Effects, 0),
		}
		for _, effectName := range spec.Effects {
			// TODO resolve the effects
			//log.Printf("loading trait %s effect %s", spec.Name, effectName)
			effect := domain.Effect{
				Name: effectName,
			}
			trait.Effects = append(trait.Effects, &effect)
		}
		l.traits = append(l.traits, &trait)
	}
	return l.traits, nil
}

func (l *TraitLoader) GetTrait(name string) (*domain.Trait, error) {
	traits, err := l.LoadTraits()
	if err != nil {
		return nil, err
	}
	for _, trait := range traits {
		if trait.Name == name {
			return trait, nil
		}
	}
	return nil, nil
}
