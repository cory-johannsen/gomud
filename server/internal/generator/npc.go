package generator

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
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
	loaders           *loader.Loaders
	npcs              *storage.NPCs
	stateGenerator    *StateGenerator
	plannerGenerator  *PlannerGenerator
	Name              string
	Spec              *domain.NPCSpec
	Minimum           int
	Maximum           int
	SpawnDelaySeconds int
}

type CharacterLifecycle interface {
	CreateCharacter(ctx context.Context, spec *domain.NPCSpec) (*domain.Character, error)
	DeleteCharacter(ctx context.Context, character *domain.Character) error
}

func NewNPCGenerator(spec *domain.GeneratorSpec, loaders *loader.Loaders, npcSpec *domain.NPCSpec, npcs *storage.NPCs,
	stateGenerator *StateGenerator, plannerGenerator *PlannerGenerator) *NPCGenerator {
	return &NPCGenerator{
		running:           false,
		npcPool:           make(map[int]*domain.NPC),
		Name:              spec.Name,
		Spec:              npcSpec,
		npcs:              npcs,
		loaders:           loaders,
		stateGenerator:    stateGenerator,
		plannerGenerator:  plannerGenerator,
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
			// create the new state for the NPC and add it to the state generator
			sensors := make(htn.Sensors)
			hodSensor, err := g.loaders.SensorLoader.GetSensor("HourOfDay")
			if err != nil {
				return err
			}
			sensors["HourOfDay"] = hodSensor
			sensors["PlayersEngaged"] = domain.PlayersEngagedSensor{
				NPC: newNPC,
			}
			sensors["PlayersInRange"] = domain.PlayersInRangeSensor{
				NPC: newNPC,
			}
			g.stateGenerator.AddState(newNPC, &htn.State{
				Sensors:    sensors,
				Properties: make(map[string]interface{}),
			})
			// fetch the NPC task graph
			taskGraph, err := g.loaders.TaskGraphLoader.GetTaskGraph(g.Spec.Name)
			if err != nil {
				return err
			}
			// create a new planner for the NPC and add it to the planner generator
			planner := &htn.Planner{
				Name:  g.Spec.Name,
				Tasks: taskGraph,
			}
			g.plannerGenerator.AddPlanner(newNPC, planner)
			// Start the NPC running
			err = newNPC.Start()
			if err != nil {
				return err
			}
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
			// stop the NPC
			err := toRemove.Stop()
			if err != nil {
				return err
			}
			// remove the NPC from the room
			err = room.RemoveNPC(toRemove)
			if err != nil {
				return err
			}
			// delete the NPC from storage
			err = g.npcs.DeleteNPC(context.Background(), toRemove)
			if err != nil {
				return err
			}
			// clean up the state and planner
			g.stateGenerator.DeleteState(toRemove)
			g.plannerGenerator.DeletePlanner(toRemove)
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
