package loader

import (
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type RoomLoader struct {
	config       *config.Config
	objectLoader *InteractiveObjectLoader
	eventBus     eventbus.Bus
	rooms        map[string]*domain.Room
}

func NewRoomLoader(cfg *config.Config, objectLoader *InteractiveObjectLoader, eventBus eventbus.Bus) *RoomLoader {
	return &RoomLoader{
		config:       cfg,
		objectLoader: objectLoader,
		eventBus:     eventBus,
		rooms:        make(map[string]*domain.Room),
	}
}

func (l *RoomLoader) LoadRooms() (map[string]*domain.Room, error) {
	if len(l.rooms) > 0 {
		return l.rooms, nil
	}
	//log.Printf("loading rooms from %s", l.config.AssetPath+"/rooms")
	items, err := os.ReadDir(l.config.AssetPath + "/rooms")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") || strings.HasSuffix(item.Name(), "txt") {
			continue
		}
		//log.Printf("loading room %s", item.Name())
		name := item.Name()
		spec := &domain.RoomSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/rooms/" + name)
		if err != nil {
			log.Printf("error reading file %s: %s", name, err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Printf("error unmarshalling file %s: %s", name, err)
			continue
		}
		room := domain.NewRoom(spec, l.GetRoom, l.objectLoader.GetInteractiveObject, l.eventBus)
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
