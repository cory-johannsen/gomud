package generator

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"sync"
	"time"
)

type NPCGenerator struct {
	mutex             sync.Mutex
	npcPool           map[int]*domain.NPC
	running           bool
	Name              string
	loaders           *loader.Loaders
	Spec              *domain.NPCSpec
	npcs              *storage.NPCs
	Minimum           int
	Maximum           int
	SpawnDelaySeconds int
}

type CharacterLifecycle interface {
	CreateCharacter(ctx context.Context, spec *domain.NPCSpec) (*domain.Character, error)
	DeleteCharacter(ctx context.Context, character *domain.Character) error
}

func NewNPCGenerator(spec *domain.GeneratorSpec, loaders *loader.Loaders, npcSpec *domain.NPCSpec, npcs *storage.NPCs) *NPCGenerator {
	return &NPCGenerator{
		running:           false,
		npcPool:           make(map[int]*domain.NPC),
		Name:              spec.Name,
		Spec:              npcSpec,
		npcs:              npcs,
		loaders:           loaders,
		Minimum:           spec.Minimum,
		Maximum:           spec.Maximum,
		SpawnDelaySeconds: spec.SpawnDelaySeconds,
	}
}

func (g *NPCGenerator) Start() error {
	log.Printf("Starting generator %s for NPC %s and room %s with minumum %d and maximum %d", g.Name, g.Spec.Name, g.Spec.Room, g.Minimum, g.Maximum)
	if g.running {
		return nil
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
	room := g.loaders.RoomLoader.GetRoom(g.Spec.Room)
	g.running = true
	for {
		if !g.running {
			log.Printf("generator %s has been stopped", g.Name)
			break
		}
		count := room.NPCCount(g.Spec.Name)
		if count < g.Minimum {
			log.Printf("room %s has %d NPCs for generator %s, minimum is %d", room.Name, count, g.Name, g.Minimum)
			newNPC, err := g.npcs.CreateNPC(context.Background(), g.Spec)
			if err != nil {
				return err
			}
			// assign the NPC the Planner
			g.npcPool[*newNPC.Id] = newNPC
			err = room.AddNPC(newNPC)
			if err != nil {
				return err
			}
		} else if count > g.Maximum {
			log.Printf("room %s has %d NPCs for generator %s, maximum is %d", room.Name, count, g.Name, g.Maximum)
			index := rand.Intn(len(g.npcPool))
			var toRemove *domain.NPC
			for _, npc := range g.npcPool {
				if index == 0 {
					toRemove = npc
					break
				}
				index--
			}
			err := room.RemoveNPC(toRemove)
			if err != nil {
				return err
			}
			err = g.npcs.DeleteNPC(context.Background(), toRemove)
			if err != nil {
				return err
			}
		}
		time.Sleep(time.Duration(g.SpawnDelaySeconds) * time.Second)
	}
	log.Printf("Generator %s exiting", g.Name)
	return nil
}

func (g *NPCGenerator) Stop() error {
	if !g.running {
		return nil
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.running = false
	return nil
}
