package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
)

type ActionLoader struct {
	config  *config.Config
	actions htn.Actions
}

func NewActionLoader(cfg *config.Config) *ActionLoader {
	return &ActionLoader{
		config:  cfg,
		actions: make(htn.Actions),
	}
}

func (l *ActionLoader) LoadActions() (htn.Actions, error) {
	if len(l.actions) > 0 {
		return l.actions, nil
	}

	return l.actions, nil
}

func (l *ActionLoader) GetAction(name string) (htn.Action, error) {
	if len(l.actions) == 0 {
		_, err := l.LoadActions()
		if err != nil {
			return nil, err
		}
	}
	return l.actions[name], nil
}

func (l *ActionLoader) SetAction(name string, action htn.Action) {
	l.actions[name] = action
}
