package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"strings"
)

type InjuryLoader struct {
	config       *config.Config
	injuries     map[domain.Severity]domain.Injuries
	effectLoader *EffectLoader
}

func NewInjuryLoader(cfg *config.Config, effectLoader *EffectLoader) *InjuryLoader {
	return &InjuryLoader{
		config:       cfg,
		injuries:     make(map[domain.Severity]domain.Injuries),
		effectLoader: effectLoader,
	}
}

func (l *InjuryLoader) LoadInjuries() (map[domain.Severity]domain.Injuries, error) {
	if len(l.injuries) > 0 {
		return l.injuries, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/injuries")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			injuries, err := l.loadInjuries(item.Name())
			if err != nil {
				return nil, err
			}
			l.injuries[domain.Severity(strings.ToUpper(item.Name()))] = injuries
		}
	}
	return l.injuries, nil
}

func (l *InjuryLoader) GetInjury(name string) (*domain.Injury, error) {
	if len(l.injuries) == 0 {
		_, err := l.LoadInjuries()
		if err != nil {
			return nil, err
		}
	}
	for _, injuries := range l.injuries {
		for _, i := range injuries {
			if i.Name == name {
				return i, nil
			}
		}
	}
	return nil, nil
}

func (l *InjuryLoader) Random(severity domain.Severity) (*domain.Injury, error) {
	if len(l.injuries) == 0 {
		_, err := l.LoadInjuries()
		if err != nil {
			return nil, err
		}
	}
	if len(l.injuries[severity]) == 0 {
		log.Printf("no injuries found for severity %s", severity)
		return nil, nil
	}
	injuries := l.injuries[severity]
	return injuries[rand.Intn(len(injuries))], nil
}

func (l *InjuryLoader) loadInjuries(path string) (domain.Injuries, error) {
	itemPath := l.config.AssetPath + "/injuries/" + path
	items, err := os.ReadDir(itemPath)
	if err != nil {
		return nil, err
	}
	injuries := make(domain.Injuries, 0)
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			log.Printf("skipping template file %s", item.Name())
			continue
		}
		//log.Printf("loading injury %s", item.Name())
		spec := &domain.InjurySpec{}
		data, err := os.ReadFile(itemPath + "/" + item.Name())
		if err != nil {
			log.Errorf("error reading file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling file %s: %v", item.Name(), err)
			continue
		}
		effect, err := l.effectLoader.GetEffect(spec.Name)
		if err != nil {
			log.Errorf("error loading injury %s effect %s: %s", spec.Name, spec.Effect, err)
		}
		injury := domain.Injury{
			Name:     spec.Name,
			Severity: domain.Severity(spec.Severity),
			Effect:   effect,
		}
		injuries = append(injuries, &injury)
	}
	return injuries, nil
}
