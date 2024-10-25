package domain

import (
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/event"
	log "github.com/sirupsen/logrus"
)

const (
	RoomEventEnter = "enter"
	RoomEventExit  = "exit"
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

func NewRoom(spec *RoomSpec, resolver RoomResolver, eventBus eventbus.Bus) *Room {
	err := eventBus.SubscribeAsync(spec.Name, func(player *Player, action string) {
		log.Printf("room %s received action %s from player %s", spec.Name, action, player.Name)
	}, false)
	if err != nil {
		log.Printf("error subscribing to event bus: %s", err)
		panic(err)
	}
	err = eventBus.SubscribeAsync(event.TickChannel, func(tick int64) {
		log.Debugf("room %s received tick %d", spec.Name, tick)
	}, false)
	if err != nil {
		log.Printf("error subscribing to event bus: %s", err)
		panic(err)
	}

	return &Room{
		ID:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Players:     make(Players),
		exitSpecs:   spec.Exits,
		exits:       make(Exits),
		resolver:    resolver,
		eventBus:    eventBus,
	}
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
	err := r.eventBus.SubscribeAsync(r.Name, player.RoomHandler, false)
	if err != nil {
		log.Errorf("error subscribing player %s to room %s: %s", player.Name, r.Name, err)
	}
	r.eventBus.Publish(r.Name, player, RoomEventEnter)
}

func (r *Room) RemovePlayer(player *Player) {
	if player.Id == nil {
		log.Printf("player %s has no ItemId", player.Name)
		return
	}
	id := *player.Id
	delete(r.Players, id)
	err := r.eventBus.Unsubscribe(r.Name, player.RoomHandler)
	if err != nil {
		log.Errorf("error unsubscribing player %s from room %s: %s", player.Name, r.Name, err)
	}

	r.eventBus.Publish(r.Name, player, RoomEventExit)
}

func (r *Room) broadcast(player *Player, action string) {

}
