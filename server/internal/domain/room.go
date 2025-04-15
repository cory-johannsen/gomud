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
	Objects     []string            `yaml:"objects"`
	Exits       map[string]ExitSpec `yaml:"exits"`
}

type RoomNPCs map[int]*NPC

type Room struct {
	ID          int64
	Name        string
	Description string
	Players     Players
	NPCs        RoomNPCs
	Objects     InteractiveObjects
	exitSpecs   map[string]ExitSpec
	exits       Exits
	resolver    RoomResolver
	eventBus    eventbus.Bus
}

func (r Room) String() string {
	return fmt.Sprintf("%s: %s", r.Name, r.Description)
}

var _ Property = &Room{}

type RoomResolver func(name string) *Room

type Rooms map[string]*Room

type RoomEvent struct {
	Room      *Room
	Character *Character
	Action    string
	Args      []interface{}
}

type RoomEventHandler func(*RoomEvent)

func NewRoom(spec *RoomSpec, resolver RoomResolver, objectResolver InteractiveObjectResolver, eventBus eventbus.Bus) *Room {
	room := &Room{
		ID:          spec.ID,
		Name:        spec.Name,
		Description: spec.Description,
		Players:     make(Players),
		NPCs:        make(RoomNPCs),
		Objects:     make(InteractiveObjects),
		exitSpecs:   spec.Exits,
		exits:       make(Exits),
		resolver:    resolver,
		eventBus:    eventBus,
	}
	err := eventBus.SubscribeAsync(event.RoomChannel, func(r *RoomEvent) {
		if r == nil || r.Room == nil || r.Room != room || r.Character == nil {
			return
		}
		log.Debugf("room %s received action %s from player %s", spec.Name, r.Action, r.Character.Name)
	}, false)
	if err != nil {
		log.Printf("error subscribing to event bus: %s", err)
		panic(err)
	}

	for _, objectName := range spec.Objects {
		object, err := objectResolver(objectName, nil)
		if err != nil {
			log.Printf("error resolving interactive object: %s", err)
			continue
		}
		room.Objects[objectName] = object
	}

	return room
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
		Room:      r,
		Character: &player.Character,
		Action:    event.RoomEventEnter,
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
		Room:      r,
		Character: &player.Character,
		Action:    event.RoomEventExit,
	})
}

func (r *Room) PlayerCount() int {
	return len(r.Players)
}

func (r *Room) AddNPC(npc *NPC) error {
	if npc.Id == nil {
		log.Printf("npc %s has no Id", npc.Name)
		return fmt.Errorf("npc %s has no Id", npc.Name)
	}
	if _, ok := r.NPCs[*npc.Id]; ok {
		log.Printf("npc %s already exists in room %s", npc.Name, r.Name)
		return fmt.Errorf("npc %s already exists in room %s", npc.Name, r.Name)
	}
	log.Printf("adding npc %s to room %s", npc.Name, r.Name)
	id := *npc.Id
	r.NPCs[id] = npc
	r.eventBus.Publish(event.RoomChannel, &RoomEvent{
		Room:      r,
		Character: &npc.Character,
		Action:    event.RoomEventEnter,
	})
	return nil
}

func (r *Room) RemoveNPC(npc *NPC) error {
	if npc.Id == nil {
		log.Printf("npc %s has no Id", npc.Name)
		return fmt.Errorf("npc %s has no Id", npc.Name)
	}
	if _, ok := r.NPCs[*npc.Id]; !ok {
		log.Printf("npc %s does not exist in room %s", npc.Name, r.Name)
		return fmt.Errorf("npc %s does not exist in room %s", npc.Name, r.Name)
	}
	id := *npc.Id
	delete(r.NPCs, id)
	r.eventBus.Publish(event.RoomChannel, &RoomEvent{
		Room:      r,
		Character: &npc.Character,
		Action:    event.RoomEventExit,
	})
	return nil
}

func (r *Room) HasNPC(npc *Character) bool {
	if npc.Id == nil {
		return false
	}
	_, ok := r.NPCs[*npc.Id]
	return ok
}

func (r *Room) NPCCount(name string) int {
	count := 0
	for _, n := range r.NPCs {
		if n.Name == name {
			count++
		}
	}
	return count
}
