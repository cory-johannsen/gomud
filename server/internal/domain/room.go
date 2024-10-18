package domain

import (
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
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
	ID          int64               `yaml:"id"`
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

func (r *Room) Value() interface{} {
	return r
}

func (r *Room) String() string {
	return fmt.Sprintf("%s: %s", r.Name, r.Description)
}

var _ Property = &Room{}

type RoomResolver func(name string) *Room

type Rooms map[string]*Room

func NewRoom(spec *RoomSpec, resolver RoomResolver, eventBus eventbus.Bus) *Room {
	eventBus.Subscribe("", func(player *Player, action string) {
		log.Printf("room %s received action %s from player %s", spec.Name, action, player.Name)
	})

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
		log.Printf("player %s has no id", player.Name)
		return
	}
	id := *player.Id
	r.Players[id] = player

	r.broadcast(player, RoomEventEnter)
}

func (r *Room) RemovePlayer(player *Player) {
	if player.Id == nil {
		log.Printf("player %s has no id", player.Name)
		return
	}
	id := *player.Id
	delete(r.Players, id)

	r.broadcast(player, RoomEventExit)
}

func (r *Room) broadcast(player *Player, action string) {
	r.eventBus.Publish(r.Name, player, action)
}
