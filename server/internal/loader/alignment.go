package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
)

type AlignmentLoader struct {
	config     *config.Config
	alignments domain.Alignments
}

func NewAlignmentLoader(cfg *config.Config) *AlignmentLoader {
	return &AlignmentLoader{
		config: cfg,
	}
}

func (l *AlignmentLoader) LoadAlignments() (domain.Alignments, error) {
	if l.alignments != nil {
		return l.alignments, nil
	}

	// read the Order alignments
	items, err := os.ReadDir(l.config.AssetPath + "/alignments/order")
	orderAlignments := make(map[string]*domain.Order)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if item.Name() == "tmpl.yaml" {
			continue
		}
		o := &domain.Order{}
		data, err := os.ReadFile(l.config.AssetPath + "/alignments/order/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, o)
		if err != nil {
			return nil, err
		}
		orderAlignments[o.Name] = o
	}
	// read the Chaos alignments
	chaosAlignments := make(map[string]*domain.Chaos)
	items, err = os.ReadDir(l.config.AssetPath + "/alignments/chaos")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if item.Name() == "tmpl.yaml" {
			continue
		}
		c := &domain.Chaos{}
		data, err := os.ReadFile(l.config.AssetPath + "/alignments/chaos/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, c)
		if err != nil {
			return nil, err
		}
		chaosAlignments[c.Name] = c
	}

	// read the alignment map
	data, err := os.ReadFile(l.config.AssetPath + "/alignments/alignments.yaml")
	if err != nil {
		return nil, err
	}

	alignments := make(map[string]string)
	err = yaml.Unmarshal(data, &alignments)
	if err != nil {
		return nil, err
	}

	for o, c := range alignments {
		a := domain.Alignment{
			Order:      orderAlignments[o],
			Chaos:      chaosAlignments[c],
			Corruption: 0,
		}
		l.alignments = append(l.alignments, a)
	}

	return l.alignments, nil
}

func (l *AlignmentLoader) GetAlignment(name string) (*domain.Alignment, error) {
	alignments, err := l.LoadAlignments()
	if err != nil {
		return nil, err
	}
	for _, a := range alignments {
		if a.Order.Name == name {
			return &a, nil
		}
	}
	return nil, nil
}
