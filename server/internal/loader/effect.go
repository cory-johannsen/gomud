package loader

import (
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type EffectLoader struct {
	config  *config.Config
	effects map[string]domain.Effect
}

func (l *EffectLoader) GetEffect(name string) (domain.Effect, error) {
	if len(l.effects) == 0 {
		return nil, fmt.Errorf("no effects loaded")
	}
	effect, ok := l.effects[name]
	if !ok {
		return nil, fmt.Errorf("effect %s not found", name)
	}
	return effect, nil
}

func NewEffectLoader(cfg *config.Config, effects domain.Effects) *EffectLoader {
	e := make(map[string]domain.Effect)
	for _, effect := range effects {
		e[effect.Name()] = effect
	}
	return &EffectLoader{
		config:  cfg,
		effects: e,
	}
}
