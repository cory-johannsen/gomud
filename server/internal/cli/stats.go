package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type StatsHandler struct {
	stateProvider domain.StateProvider
}

func (s *StatsHandler) Help(args []string) string {
	return "display your stats"
}

func (s *StatsHandler) State() domain.State {
	return s.stateProvider()
}

func (s *StatsHandler) Handle(ctx context.Context, args []string) (string, error) {
	stats := s.stateProvider().Player().Stats()
	bonuses := s.stateProvider().Player().StatBonuses()
	advances := s.stateProvider().Player().ConsumedBonusAdvances()
	return domain.StatsString(stats, bonuses, advances), nil
}

var _ Handler = &StatsHandler{}
