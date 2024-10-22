package loader

import (
	"errors"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type DisorderLoader struct {
	config    *config.Config
	disorders map[string]*domain.Disorder
}

func NewDisorderLoader(config *config.Config) *DisorderLoader {
	return &DisorderLoader{
		config:    config,
		disorders: make(map[string]*domain.Disorder),
	}
}

func (l *DisorderLoader) Load() map[string]*domain.Disorder {
	return l.disorders
}

func (l *DisorderLoader) GetDisorder(name string) (*domain.Disorder, error) {
	disorder, ok := l.disorders[name]
	if !ok {
		return nil, errors.New("disorder not found")
	}
	return disorder, nil
}
