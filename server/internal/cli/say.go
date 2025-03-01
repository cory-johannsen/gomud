package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/event"
)

type SayHandler struct {
	stateProvider domain.StateProvider
}

func (s *SayHandler) Handle(ctx context.Context, args []string) (string, error) {
	if len(args) < 1 {
		return "say what?", nil
	}
	a := make([]interface{}, 0)
	for _, arg := range args {
		a = append(a, arg)
	}
	s.stateProvider().Player().Connection.EventBus().Publish(event.RoomChannel, &domain.RoomEvent{
		Room:      s.stateProvider().Player().Room(),
		Character: &s.stateProvider().Player().Character,
		Action:    event.RoomEventSay,
		Args:      a,
	})
	return "", nil
}

func (s *SayHandler) Help(args []string) string {
	return "say something.  Usage: say <message>"
}

func (s *SayHandler) State() domain.State {
	return s.stateProvider()
}

func NewSayHandler(stateProvider domain.StateProvider) *SayHandler {
	return &SayHandler{
		stateProvider: stateProvider,
	}
}

var _ Handler = &SayHandler{}
