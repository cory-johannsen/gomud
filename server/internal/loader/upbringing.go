package loader

import (
	"errors"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"math/rand"
	"os"
	"strings"
)

type UpbringingLoader struct {
	config      *config.Config
	upbringings map[string]*domain.Upbringing
}

func NewUpbringingLoader(config *config.Config) *UpbringingLoader {
	return &UpbringingLoader{
		config:      config,
		upbringings: make(map[string]*domain.Upbringing),
	}
}

func (l *UpbringingLoader) LoadUpbringings() (map[string]*domain.Upbringing, error) {
	if len(l.upbringings) > 0 {
		return l.upbringings, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/upbringing")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if !strings.HasSuffix(item.Name(), ".yaml") || strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		log.Debugf("loading upbringing %s", item.Name())
		upbringing := &domain.Upbringing{}
		data, err := os.ReadFile(l.config.AssetPath + "/upbringing/" + item.Name())
		if err != nil {
			log.Errorf("error reading upbringing file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, upbringing)
		if err != nil {
			log.Errorf("error unmarshalling upbringing file %s: %v", item.Name(), err)
			continue
		}
		l.upbringings[upbringing.Name] = upbringing
	}
	return l.upbringings, nil
}

func (l *UpbringingLoader) GetUpbringing(name string) (*domain.Upbringing, error) {
	if len(l.upbringings) == 0 {
		_, err := l.LoadUpbringings()
		if err != nil {
			return nil, err
		}
	}
	upbringing, ok := l.upbringings[name]
	if !ok {
		return nil, errors.New("upbringing not found")
	}
	return upbringing, nil
}

func (l *UpbringingLoader) Random() *domain.Upbringing {
	upbringings, err := l.LoadUpbringings()
	if err != nil {
		return nil
	}
	u := make([]*domain.Upbringing, 0)
	for _, upbringing := range upbringings {
		u = append(u, upbringing)
	}
	return u[rand.Intn(len(u))]
}
