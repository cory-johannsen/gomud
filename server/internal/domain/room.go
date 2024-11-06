package domain

import (
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/event"
	log "github.com/sirupsen/logrus"
)

type ExitSpec struct {
	Direction   string `yaml:"direction"`
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Target      string `yaml:"target"`
}

type Exit struct {
	Direction   string
	Name        string
	Description string
	Target      *Room
}
type Exits map[string]*Exit

type RoomSpec struct {
	ID          int64               `yaml:"ItemId"`
	Name        string              `yaml:"name"`
	Description string              `yaml:"description"`
	Exits       map[string]ExitSpec `yaml:"exits"`
}

type Room struct {
	ID          int64
	Name        string
	Description string
	Players     Players
	exitSpecs   map[string]ExitSpec
	exits       Exits
	resolver    RoomResolver
	eventBus    eventbus.Bus
}
type RoomResolver func(name string) *Room

type Rooms map[string]*Room

type RoomEvent struct {
	Room   *Room
	Player *Player
	Action string
	Args   []interface{}
}

type RoomEventHandler func(*RoomEvent)

func NewRoom(spec *RoomSpec, resolver RoomResolver, eventBus eventbus.Bus) *Room {
	room := &Room{
		ID:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Players:     make(Players),
		exitSpecs:   spec.Exits,
		exits:       make(Exits),
		resolver:    resolver,
		eventBus:    eventBus,
	}
	err := eventBus.SubscribeAsync(event.RoomChannel, func(r *RoomEvent) {
		if r == nil || r.Room == nil || r.Room != room || r.Player == nil {
			return
		}
		log.Debugf("room %s received action %s from player %s", spec.Name, r.Action, r.Player.Name)
	}, false)
	if err != nil {
		log.Printf("error subscribing to event bus: %s", err)
		panic(err)
	}
	err = eventBus.SubscribeAsync(event.TickChannel, func(tick int64) {
		log.Debugf("room %s received tick %d", spec.Name, tick)
		eventBus.Publish(spec.Name, nil, fmt.Sprintf("tick %d", tick))
	}, false)
	if err != nil {
		log.Printf("error subscribing to event bus: %s", err)
		panic(err)
	}

	return room
}

func (r *Room) Value() interface{} {
	return r
}

func (r *Room) String() string {
	return fmt.Sprintf("%s: %s", r.Name, r.Description)
}

func (r *Room) Exits() Exits {
	if len(r.exits) > 0 {
		return r.exits
	}
	for direction, spec := range r.exitSpecs {
		exit := r.resolver(spec.Target)
		r.exits[direction] = &Exit{
			Direction:   direction,
			Name:        spec.Name,
			Description: spec.Description,
			Target:      exit,
		}
	}
	return r.exits
}

func (r *Room) AddPlayer(player *Player) {
	if player.Id == nil {
		log.Printf("player %s has no ItemId", player.Name)
		return
	}
	id := *player.Id
	r.Players[id] = player
	r.eventBus.Publish(event.RoomChannel, &RoomEvent{
		Room:   r,
		Player: player,
		Action: event.RoomEventEnter,
	})
}

func (r *Room) RemovePlayer(player *Player) {
	if player.Id == nil {
		log.Printf("player %s has no ItemId", player.Name)
		return
	}
	id := *player.Id
	delete(r.Players, id)
	r.eventBus.Publish(event.RoomChannel, &RoomEvent{
		Room:   r,
		Player: player,
		Action: event.RoomEventExit,
	})
}
