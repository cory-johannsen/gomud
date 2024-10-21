package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type CharacterHandler struct {
	stateProvider StateProvider
}

func (c *CharacterHandler) Handle(ctx context.Context, args []string) (string, error) {
	return c.stateProvider().Player().String(), nil
}

func (c *CharacterHandler) Help(args []string) string {
	return "display your character"
}

func (c *CharacterHandler) State() domain.State {
	return c.stateProvider()
}

var _ Handler = &CharacterHandler{}
