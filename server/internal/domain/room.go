package domain

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

type ExitSpec struct {
	Direction   string
	Name        string
	Description string
	Target      string
}

type Exit struct {
	Direction   string
	Name        string
	Description string
	Target      *Room
}
type Exits map[string]*Exit

type RoomSpec struct {
	ID          int64
	Name        string
	Description string
	Exits       map[string]ExitSpec
}

type Room struct {
	ID          int64
	Name        string
	Description string
	Players     Players
	exitSpecs   map[string]ExitSpec
	exits       Exits
	resolver    RoomResolver
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

func NewRoom(spec *RoomSpec, resolver RoomResolver) *Room {
	return &Room{
		ID:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Players:     make(Players),
		exitSpecs:   spec.Exits,
		exits:       make(Exits),
		resolver:    resolver,
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
}

func (r *Room) RemovePlayer(player *Player) {
	if player.Id == nil {
		log.Printf("player %s has no id", player.Name)
		return
	}
	id := *player.Id
	delete(r.Players, id)
}
