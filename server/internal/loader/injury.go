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
	config   *config.Config
	injuries map[domain.Severity]domain.Injuries
}

func NewInjuryLoader(cfg *config.Config) *InjuryLoader {
	return &InjuryLoader{
		config:   cfg,
		injuries: make(map[domain.Severity]domain.Injuries),
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
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			log.Printf("skipping template file %s", item.Name())
			continue
		}
		log.Printf("loading injury %s", item.Name())
		spec := &domain.InjurySpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/injuries/" + item.Name())
		if err != nil {
			log.Errorf("error reading file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling file %s: %v", item.Name(), err)
			continue
		}
		injury := domain.Injury{
			Name:     spec.Name,
			Severity: domain.Severity(spec.Severity),
			Effect: domain.Effect{
				Name: spec.Effect,
			},
		}
		if _, ok := l.injuries[injury.Severity]; !ok {
			l.injuries[injury.Severity] = make(domain.Injuries, 0)
		}
		l.injuries[injury.Severity] = append(l.injuries[injury.Severity], &injury)
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
