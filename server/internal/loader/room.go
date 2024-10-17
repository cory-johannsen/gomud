package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type RoomLoader struct {
	config   *config.Config
	eventBus goeventbus.EventBus
	rooms    map[string]*domain.Room
}

func NewRoomLoader(cfg *config.Config, eventBus goeventbus.EventBus) *RoomLoader {
	return &RoomLoader{
		config:   cfg,
		eventBus: eventBus,
		rooms:    make(map[string]*domain.Room),
	}
}

func (l *RoomLoader) LoadRooms() (map[string]*domain.Room, error) {
	if len(l.rooms) > 0 {
		return l.rooms, nil
	}
	log.Printf("loading rooms from %s", l.config.AssetPath+"/rooms")
	items, err := os.ReadDir(l.config.AssetPath + "/rooms")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			log.Printf("skipping template file %s", item.Name())
			continue
		}
		log.Printf("loading room %s", item.Name())
		spec := &domain.RoomSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/rooms/" + item.Name())
		if err != nil {
			log.Printf("error reading file %s: %s", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Printf("error unmarshalling file %s: %s", item.Name(), err)
			continue
		}
		room := domain.NewRoom(spec, l.GetRoom, l.eventBus)
		l.rooms[room.Name] = room
	}
	return l.rooms, nil
}

func (l *RoomLoader) GetRoom(name string) *domain.Room {
	if len(l.rooms) == 0 {
		_, err := l.LoadRooms()
		if err != nil {
			log.Printf("error loading rooms: %s", err)
			return nil
		}
	}
	return l.rooms[name]
}
