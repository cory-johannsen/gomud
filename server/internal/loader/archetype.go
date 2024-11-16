package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type ArchetypeLoader struct {
	config          *config.Config
	archetypes      domain.Archetypes
	traitLoader     *TraitLoader
	equipmentLoader *EquipmentLoader
}

func NewArchetypeLoader(cfg *config.Config, traitLoader *TraitLoader, equipmentLoader *EquipmentLoader) *ArchetypeLoader {
	return &ArchetypeLoader{
		config:          cfg,
		archetypes:      make(domain.Archetypes, 0),
		traitLoader:     traitLoader,
		equipmentLoader: equipmentLoader,
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
			StartingEquipment: domain.StartingEquipment{
				OneEach: make(domain.Items, 0),
				OneOf:   make(domain.Items, 0),
			},
			Traits: make(domain.Traits, 0),
		}
		for _, traitName := range spec.Traits {
			trait, err := l.traitLoader.GetTrait(traitName)
			if err != nil {
				return nil, err
			}
			if trait == nil {
				log.Printf("trait %s not found", traitName)
				continue
			}
			archetype.Traits = append(archetype.Traits, trait)
		}
		for _, itemName := range spec.StartingEquipment.OneEach {
			i, err := l.equipmentLoader.ItemFromName(itemName)
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.Printf("item %s not found", itemName)
				continue
			}
			archetype.StartingEquipment.OneEach = append(archetype.StartingEquipment.OneEach, i)
		}
		for _, itemName := range spec.StartingEquipment.OneOf {
			i, err := l.equipmentLoader.ItemFromName(itemName)
			if err != nil {
				return nil, err
			}
			if i == nil {
				log.Printf("item %s not found", itemName)
				continue
			}
			archetype.StartingEquipment.OneOf = append(archetype.StartingEquipment.OneOf, i)
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
