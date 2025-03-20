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

type NpcInitializer interface {
	Initialize(ctx context.Context, spec *domain.NPCSpec) (*domain.NPC, error)
}

type NpcInitializers map[string]NpcInitializer

func (i NpcInitializers) GetInitializer(name string) NpcInitializer {
	init, ok := i[name]
	if !ok {
		return i["Default"]
	}
	return init
}

func NewNpcInitializers(loaders *loader.Loaders, npcs *storage.NPCs, domainGenerator *DomainGenerator, plannerGenerator *PlannerGenerator) NpcInitializers {
	initializers := make(NpcInitializers)
	initializers["Default"] = &BaseNpcInitializer{
		Loaders:          loaders,
		DomainGenerator:  domainGenerator,
		NPCs:             npcs,
		PlannerGenerator: plannerGenerator,
	}
	initializers["Wayne Dawg"] = &WayneDawgInitializer{
		BaseNpcInitializer: BaseNpcInitializer{
			Loaders:          loaders,
			DomainGenerator:  domainGenerator,
			NPCs:             npcs,
			PlannerGenerator: plannerGenerator,
		},
	}
	return initializers
}

// BaseNpcInitializer is the default NPC initializer, which creates a standard domain and planner.  This initializer is used when no custom initializer is provided.
// This initializer sets up the standard sensors for the NPC, including HourOfDay, PlayersEngaged, and PlayersInRange.
type BaseNpcInitializer struct {
	Loaders          *loader.Loaders
	DomainGenerator  *DomainGenerator
	NPCs             *storage.NPCs
	PlannerGenerator *PlannerGenerator
}

func (i *BaseNpcInitializer) Initialize(ctx context.Context, spec *domain.NPCSpec) (*domain.NPC, error) {
	npcDomain := initializeDomain()
	i.DomainGenerator.AddDomain(spec.Name, npcDomain)

	// fetch the NPC task graph
	taskGraph, err := i.Loaders.TaskGraphLoader.GetTaskGraph(spec.Name)
	if err != nil {
		return nil, err
	}
	// create a new planner for the NPC and add it to the planner generator
	planner := &htn.Planner{
		Name:  spec.Name,
		Tasks: taskGraph,
	}
	i.PlannerGenerator.AddPlanner(spec.Name, planner)

	newNPC, err := i.NPCs.CreateNPC(context.Background(), spec)
	if err != nil {
		return nil, err
	}
	// set the domain owner to the new NPC
	npcDomain.Owner = newNPC

	// initialize the NPC sensors
	hodSensor, err := i.Loaders.SensorLoader.GetSensor("HourOfDay")
	if err != nil {
		return nil, err
	}
	npcDomain.Sensors["HourOfDay"] = hodSensor
	npcDomain.Sensors["PlayersEngaged"] = &domain.PlayersEngagedSensor{
		NPC: newNPC,
	}
	npcDomain.Sensors["PlayersInRange"] = &domain.PlayersInRangeSensor{
		NPC: newNPC,
	}
	return newNPC, nil
}

type WayneDawgInitializer struct {
	BaseNpcInitializer
}

func (i *WayneDawgInitializer) Initialize(ctx context.Context, spec *domain.NPCSpec) (*domain.NPC, error) {
	instance, err := i.BaseNpcInitializer.Initialize(ctx, spec)
	if err != nil {
		return nil, err
	}
	instance.Domain.Sensors["Drunkenness"] = &domain.IntoxicationSensor{
		SobrietySensor: domain.SobrietySensor{
			NPC:       instance,
			Substance: "Alcohol",
		},
	}
	instance.Domain.Sensors["Stoned"] = &domain.SobrietySensor{
		NPC:       instance,
		Substance: "Weed",
	}
	return instance, nil
}

type NpcGenerator struct {
	mutex             sync.Mutex
	initializers      NpcInitializers
	npcPool           map[int]*domain.NPC
	running           bool
	loaders           *loader.Loaders
	npcs              *storage.NPCs
	domainGenerator   *DomainGenerator
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

func NewNpcGenerator(spec *domain.GeneratorSpec, loaders *loader.Loaders, npcSpec *domain.NPCSpec, npcs *storage.NPCs,
	domainGenerator *DomainGenerator, plannerGenerator *PlannerGenerator, initializers NpcInitializers) *NpcGenerator {
	return &NpcGenerator{
		running:           false,
		initializers:      initializers,
		npcPool:           make(map[int]*domain.NPC),
		Name:              spec.Name,
		Spec:              npcSpec,
		npcs:              npcs,
		loaders:           loaders,
		domainGenerator:   domainGenerator,
		plannerGenerator:  plannerGenerator,
		Minimum:           spec.Minimum,
		Maximum:           spec.Maximum,
		SpawnDelaySeconds: spec.SpawnDelaySeconds,
	}
}

func (g *NpcGenerator) Start() error {
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

			initializer := g.initializers.GetInitializer(g.Spec.Name)
			newNPC, err := initializer.Initialize(context.Background(), g.Spec)

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
			// clean up the domain and planner
			g.domainGenerator.DeleteDomain(toRemove.Name)
			g.plannerGenerator.DeletePlanner(toRemove.Name)
		}
		time.Sleep(time.Duration(g.SpawnDelaySeconds) * time.Second)
	}
	log.Printf("Generator %s exiting", g.Name)
	return nil
}

func initializeDomain() *htn.Domain {
	properties := make(map[string]any)
	properties["HourOfDay"] = &htn.Property[int64]{
		Name: "HourOfDay",
		Value: func(state *htn.Domain) int64 {
			sensor, err := state.Sensor("HourOfDay")
			if err != nil {
				log.Fatal(err)
			}
			val, err := sensor.(*htn.HourOfDaySensor).Get()
			if err != nil {
				log.Fatal(err)
			}
			return val
		},
	}
	properties["PlayersInRange"] = &htn.Property[int64]{
		Name: "PlayersInRange",
		Value: func(state *htn.Domain) int64 {
			sensor, err := state.Sensor("PlayersInRange")
			if err != nil {
				log.Fatal(err)
			}
			val, err := sensor.(*domain.PlayersInRangeSensor).Get()
			if err != nil {
				log.Fatal(err)
			}
			log.Debugf("PlayersInRange property: %d players in range", val)
			return int64(val)
		},
	}
	properties["PlayersAvailable"] = &htn.Property[int64]{
		Name: "PlayersAvailable",
		Value: func(state *htn.Domain) int64 {
			sensor, err := state.Sensor("PlayersEngaged")
			if err != nil {
				log.Fatal(err)
			}
			val, err := sensor.(*domain.PlayersEngagedSensor).Get()
			if err != nil {
				log.Fatal(err)
			}
			log.Debugf("PlayersAvailable property: %d players available", val)
			npc := state.Owner.(*domain.NPC)
			available := npc.Room().PlayerCount() - val
			return int64(available)
		},
	}
	properties["PlayersEngaged"] = &htn.Property[int64]{
		Name: "PlayersEngaged",
		Value: func(state *htn.Domain) int64 {
			npc := state.Owner.(*domain.NPC)
			engaged := npc.PlayersEngaged()
			return int64(engaged)
		},
	}
	state := &htn.Domain{
		Sensors:    make(htn.Sensors),
		Properties: properties,
	}
	return state
}

func (g *NpcGenerator) Stop() error {
	if !g.running {
		return nil
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()
	g.running = false
	return nil
}
