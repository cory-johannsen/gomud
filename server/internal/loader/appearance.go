package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
)

type AppearanceLoader struct {
	config    *config.Config
	tats      domain.Tattoos
	drawbacks domain.Drawbacks
	marks     domain.DistinguishingMarks
}

func NewAppearanceLoader(cfg *config.Config) *AppearanceLoader {
	return &AppearanceLoader{
		config: cfg,
	}
}

func (l *AppearanceLoader) LoadTattoos() (domain.Tattoos, error) {
	if l.tats != nil {
		return l.tats, nil
	}
	tats := make(domain.Tattoos, 0)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/tattoos.json")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &tats)
	if err != nil {
		return nil, err
	}
	l.tats = tats
	return tats, nil
}

func (l *AppearanceLoader) LoadDrawbacks() (domain.Drawbacks, error) {
	if l.drawbacks != nil {
		return l.drawbacks, nil
	}
	drawbacks := make(domain.Drawbacks, 0)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/drawbacks.json")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &drawbacks)
	if err != nil {
		return nil, err
	}
	l.drawbacks = drawbacks
	return drawbacks, nil
}

func (l *AppearanceLoader) LoadDistinguishingMarks() (domain.DistinguishingMarks, error) {
	if l.marks != nil {
		return l.marks, nil
	}
	marks := make(domain.DistinguishingMarks, 0)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/distinguishing_marks.json")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &marks)
	if err != nil {
		return nil, err
	}
	l.marks = marks
	return marks, nil
}
