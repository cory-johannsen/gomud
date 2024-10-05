package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
	"strings"
)

type BackgroundLoader struct {
	config      *config.Config
	backgrounds domain.Backgrounds
	traitLoader *TraitLoader
}

func NewBackgroundLoader(cfg *config.Config, traitLoader *TraitLoader) *BackgroundLoader {
	return &BackgroundLoader{
		config:      cfg,
		backgrounds: make(domain.Backgrounds, 0),
		traitLoader: traitLoader,
	}
}

func (l *BackgroundLoader) LoadBackgrounds() (domain.Backgrounds, error) {
	if len(l.backgrounds) > 0 {
		return l.backgrounds, nil
	}
	log.Printf("loading backgrounds from %s", l.config.AssetPath+"/backgrounds")
	items, err := os.ReadDir(l.config.AssetPath + "/backgrounds")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			log.Printf("skipping template file %s", item.Name())
			continue
		}
		log.Printf("loading background %s", item.Name())
		spec := &domain.BackgroundSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/backgrounds/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}

		// iterates the trait names and resolve the traits
		traits := make(domain.Traits, 0)
		for _, traitName := range spec.Traits {
			log.Printf("loading trait %s", traitName)
			trait, err := l.traitLoader.GetTrait(traitName)
			if err != nil {
				return nil, err
			}
			if trait == nil {
				log.Printf("could not find trait %s for background %s", traitName, spec.Name)
				continue
			}
			traits = append(traits, trait)
		}
		l.backgrounds = append(l.backgrounds, &domain.Background{
			Name:        spec.Name,
			Description: spec.Description,
			Traits:      traits,
		})
	}
	return l.backgrounds, nil
}

func (l *BackgroundLoader) RandomBackground() (*domain.Background, error) {
	backgrounds, err := l.LoadBackgrounds()
	if err != nil {
		return nil, err
	}
	index := rand.Intn(len(backgrounds))
	return backgrounds[index], nil
}

func (l *BackgroundLoader) GetBackground(name string) (*domain.Background, error) {
	backgrounds, err := l.LoadBackgrounds()
	if err != nil {
		return nil, err
	}
	for _, background := range backgrounds {
		if background.Name == name {
			return background, nil
		}
	}
	return nil, nil
}
