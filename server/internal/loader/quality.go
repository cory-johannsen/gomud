package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type QualityLoader struct {
	config       *config.Config
	qualities    map[string]*domain.Quality
	effectLoader *EffectLoader
}

func NewQualityLoader(cfg *config.Config, effectLoader *EffectLoader) *QualityLoader {
	return &QualityLoader{
		config:       cfg,
		qualities:    make(map[string]*domain.Quality),
		effectLoader: effectLoader,
	}
}

func (l *QualityLoader) LoadQualities() (map[string]*domain.Quality, error) {
	if len(l.qualities) > 0 {
		return l.qualities, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/equipment/qualities")
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
		log.Debugf("loading quality %s", item.Name())
		spec := &domain.QualitySpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/qualities/" + item.Name())
		if err != nil {
			log.Errorf("error reading quality file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling quality file %s: %v", item.Name(), err)
			continue
		}
		quality := &domain.Quality{
			Name:       spec.Name,
			Conditions: make([]string, 0),
			Effects:    make([]domain.Effect, 0),
			Targets:    make([]string, 0),
		}
		for _, condition := range spec.Conditions {
			quality.Conditions = append(quality.Conditions, condition)
		}
		e, err := l.effectLoader.GetEffect(spec.Name)
		if err != nil {
			log.Errorf("error getting effect %s: %v", spec.Name, err)
			continue
		}
		quality.Effects = append(quality.Effects, e)
		for _, target := range spec.Targets {
			quality.Targets = append(quality.Targets, target)
		}
		l.qualities[spec.Name] = quality
	}
	return l.qualities, nil
}

func (l *QualityLoader) GetQuality(name string) (*domain.Quality, error) {
	if len(l.qualities) == 0 {
		_, err := l.LoadQualities()
		if err != nil {
			return nil, err
		}
	}
	return l.qualities[name], nil
}
