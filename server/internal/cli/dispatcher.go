package cli

import (
	"context"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/io"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"strings"
	"time"
)

type StateProvider func() State

type Dispatcher struct {
	handlers map[string]Handler
	ctx      context.Context
	state    State
	eventBus goeventbus.EventBus
}

func (d *Dispatcher) State() State {
	return d.state
}

func NewDispatcher(stateConstructor StateConstructor, players *storage.Players, generator *generator.PlayerGenerator,
	teams *loader.TeamLoader, rooms *loader.RoomLoader, conn io.Connection, eventBus goeventbus.EventBus) *Dispatcher {
	dispatcher := &Dispatcher{
		handlers: make(map[string]Handler),
		ctx:      context.Background(),
		eventBus: eventBus,
	}
	quit := &QuitHandler{
		stateProvider: dispatcher.State,
		players:       players,
	}
	quitAliases := CreateAliases(quit, "exit", "q")
	dispatcher.Register("quit", quit)
	for _, alias := range quitAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	helpHandler := &HelpHandler{stateProvider: dispatcher.State, dispatcher: dispatcher}
	dispatcher.Register("help", helpHandler)
	helpAliases := CreateAliases(helpHandler, "?", "h")
	for _, alias := range helpAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	dispatcher.Register("login", NewLoginHandler(stateConstructor, players, generator, teams, rooms, conn))

	characterHandler := CharacterHandler{stateProvider: dispatcher.State}
	dispatcher.Register("character", &characterHandler)
	characterAliases := CreateAliases(&characterHandler, "c", "char", "me", "self")
	for _, alias := range characterAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	lookHandler := NewLookHandler(dispatcher.State)
	dispatcher.Register("look", lookHandler)
	lookAliases := CreateAliases(lookHandler, "l")
	for _, alias := range lookAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	moveHandler := NewMoveHandler(dispatcher.State, players, rooms)
	dispatcher.Register("move", moveHandler)
	moveAliases := CreateAliases(moveHandler, "m")
	for _, alias := range moveAliases {
		dispatcher.Register(alias.Alias, alias)
	}

	directions := map[string][]string{
		"North":     {"n", "N", "north"},
		"South":     {"s", "S", "south"},
		"East":      {"e", "E", "east"},
		"West":      {"w", "W", "west"},
		"Northeast": {"ne", "NE", "northeast"},
		"Northwest": {"nw", "NW", "northwest"},
		"Southeast": {"se", "SE", "southeast"},
		"Southwest": {"sw", "SW", "southwest"},
		"Up":        {"u", "U", "up"},
		"Down":      {"d", "D", "down"},
	}
	for direction, aliases := range directions {
		handler := NewDirectionalMoveHandler(dispatcher.State, players, rooms, direction)
		dispatcher.Register(direction, handler)
		handlerAliases := CreateAliases(handler, aliases...)
		for _, alias := range handlerAliases {
			dispatcher.Register(alias.Alias, alias)
		}
	}

	eventBus.Channel("tick").Subscriber().Listen(func(ctx goeventbus.Context) {
		msg := ctx.Result()
		t := msg.Data.(time.Time)
		log.Debugf("tick %v", t)
	})

	return dispatcher
}

func (d *Dispatcher) Register(name string, handler Handler) {
	d.handlers[name] = handler
}

func (d *Dispatcher) Prompt() string {
	if d.State() == nil || d.State().Player() == nil {
		return "> "
	}
	player := d.State().Player()
	return fmt.Sprintf("%s [%s, %s]> ", player.Name, player.Condition(), player.Peril().Condition.String())
}

func (d *Dispatcher) Tab(buffer string) string {
	return "" // do nothing
}

func (d *Dispatcher) Eval(buffer string) string {
	fields := strings.Fields(buffer)

	if len(fields) == 0 {
		return ""
	} else {
		cmd := fields[0]

		handler, ok := d.handlers[cmd]
		if !ok {
			return fmt.Sprintf("unrecognized command \"%s\"", cmd)
		}

		result, err := handler.Handle(d.ctx, fields[1:])
		if err != nil {
			return fmt.Sprintf("error: %s", err)
		}
		state := handler.State()
		if state != nil {
			d.state = state
		}
		return result
	}
}
