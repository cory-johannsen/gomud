package engine

import (
	"bufio"
	"errors"
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/api/grpc"
	"github.com/cory-johannsen/gomud/internal/api/rest"
	"github.com/cory-johannsen/gomud/internal/cli"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net"
	"strings"
	"time"
)

type PlayerState struct {
	player  *domain.Player
	sensors htn.Sensors
	vars    map[string]domain.Property
}

func (s *PlayerState) Domain() *htn.Domain {
	props := make(map[string]any)
	for k, v := range s.vars {
		props[k] = v
	}
	return &htn.Domain{
		Owner:      s.player,
		Sensors:    s.sensors,
		Properties: props,
	}
}

func (s *PlayerState) Player() *domain.Player {
	return s.player
}

func (s *PlayerState) Sensor(name string) any {
	return s.sensors[name]
}

func (s *PlayerState) AddSensor(name string, value any) {
	s.sensors[name] = value
}

func (s *PlayerState) Property(key string) domain.Property {
	return s.vars[key]
}

var _ domain.GameState = &PlayerState{}

func NewState(player *domain.Player, sensors htn.Sensors) domain.GameState {
	return &PlayerState{
		player:  player,
		sensors: sensors,
		vars:    make(map[string]domain.Property),
	}
}

type ClientConnection struct {
	Connection net.Conn
	eventBus   eventbus.Bus
	sensors    htn.Sensors
}

func NewClientConnection(conn net.Conn, bus eventbus.Bus, sensors htn.Sensors) *ClientConnection {
	return &ClientConnection{
		Connection: conn,
		eventBus:   bus,
		sensors:    sensors,
	}
}

func (c *ClientConnection) Read() string {
	netData, err := bufio.NewReader(c.Connection).ReadString('\n')
	if err != nil {
		log.Debug(err)
		return ""
	}
	return strings.TrimSuffix(netData, "\r\n")
}

func (c *ClientConnection) Write(data string) int {
	written, err := c.Connection.Write([]byte(data))
	if err != nil {
		log.Debug(err)
	}
	return written
}

func (c *ClientConnection) Writeln(data string) int {
	return c.Write(fmt.Sprintf("%s\n", data))
}

func (c *ClientConnection) EventBus() eventbus.Bus {
	return c.eventBus
}

func (c *ClientConnection) Sensors() htn.Sensors {
	return c.sensors
}

type Client struct {
	Connection net.Conn
	EventBus   eventbus.Bus
	Dispatcher *cli.Dispatcher
}

func NewClient(players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, rooms *loader.RoomLoader,
	skills *loader.SkillLoader, conn net.Conn, eventBus eventbus.Bus, sensors htn.Sensors) *Client {
	dispatcher := cli.NewDispatcher(NewState, players, generator, teams, rooms, skills, NewClientConnection(conn, eventBus, sensors), eventBus, sensors)
	return &Client{
		Connection: conn,
		Dispatcher: dispatcher,
	}
}

func (c *Client) Connect() {
	log.Printf("Serving client %s\n", c.Connection.RemoteAddr().String())
	prompt := fmt.Sprintf("\n%s", c.Dispatcher.Prompt())
	_, err := c.Connection.Write([]byte(cli.Welcome()))
	if err != nil {
		panic(err)
	}
	_, err = c.Connection.Write([]byte(prompt))
	if err != nil {
		panic(err)
	}

	for {
		netData, err := bufio.NewReader(c.Connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		result := fmt.Sprintf("%s\n", c.Dispatcher.Eval(netData))
		_, err = c.Connection.Write([]byte(result))
		if err != nil {
			log.Println(err)
			break
		}
		if strings.HasPrefix(result, cli.QuitMessage) {
			break
		}
		prompt = fmt.Sprintf("\n%s", c.Dispatcher.Prompt())
		_, err = c.Connection.Write([]byte(prompt))
		if err != nil {
			log.Println(err)
			break
		}
	}
	log.Printf("Client %s disconnected\n", c.Connection.RemoteAddr().String())
	c.Connection.Close()
}

type Server struct {
	port             string
	db               *storage.Database
	players          *storage.Players
	npcs             *storage.NPCs
	loaders          *loader.Loaders
	playerGenerator  *generator.PlayerGenerator
	domainGenerator  *generator.DomainGenerator
	plannerGenerator *generator.PlannerGenerator
	dispatcher       *cli.Dispatcher
	eventBus         eventbus.Bus
	clock            *event.Clock
	config           *config.Config
	grpcServer       *grpc.GrpcServer
	restServer       *rest.RestServer
}

func NewServer(config *config.Config, db *storage.Database, players *storage.Players, npcs *storage.NPCs, loaders *loader.Loaders,
	playerGenerator *generator.PlayerGenerator, stateGenerator *generator.DomainGenerator, plannerGenerator *generator.PlannerGenerator, eventBus eventbus.Bus, clock *event.Clock,
	grpcServer *grpc.GrpcServer, restServer *rest.RestServer) *Server {
	return &Server{
		port:             config.Port,
		db:               db,
		players:          players,
		npcs:             npcs,
		loaders:          loaders,
		playerGenerator:  playerGenerator,
		domainGenerator:  stateGenerator,
		plannerGenerator: plannerGenerator,
		eventBus:         eventBus,
		clock:            clock,
		config:           config,
		grpcServer:       grpcServer,
		restServer:       restServer,
	}
}

func (s *Server) Start() {
	// Define the non-asset conditions, global actions, and general purpose sensors
	conditions := initializeConditions()
	actions := initializeActions()
	sensors := s.initializeSensors()

	// preload the assets
	err := s.loaders.Preload(conditions, actions, sensors, s.npcs)
	if err != nil {
		panic(err)
	}

	initializers := generator.NewNpcInitializers(s.loaders, s.npcs, s.domainGenerator, s.plannerGenerator)
	// TODO: populate the NPC-specific initializers

	// start the generators
	specs, err := s.loaders.GeneratorLoader.LoadGenerators()
	if err != nil {
		panic(err)
	}
	for _, spec := range specs {
		s.startGenerator(spec, initializers)
	}

	go func() {
		log.Printf("Starting grpc server on %s\n", s.config.GrpcAddress)
		err = grpc.StartGRPCServer(s.grpcServer)
		if err != nil {
			panic(err)
		}
	}()

	go func() {
		log.Printf("Starting rest server on %s\n", s.config.RestAddress)
		err = rest.StartRestServer(s.restServer)
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting server on port %s\n", s.port)
	l, err := net.Listen("tcp4", fmt.Sprintf(":%s", s.port))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	go func() {
		s.clock.Start()
		defer s.clock.Stop()
	}()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		client := NewClient(s.players, s.playerGenerator, s.loaders.TeamLoader, s.loaders.RoomLoader, s.loaders.SkillLoader, c, s.eventBus, sensors)
		go client.Connect()
	}
}

func (s *Server) initializeSensors() htn.Sensors {
	now := time.Now()
	sensors := htn.Sensors{
		"HourOfDay": &htn.HourOfDaySensor{
			TickSensor: htn.TickSensor{
				StartedAt:    now,
				TickDuration: time.Duration(s.config.TickDurationMillis) * time.Millisecond,
			},
			TicksPerHour: 60,
			Offset:       8,
		},
		"TimeOfDay": &htn.TimeOfDaySensor{
			TickSensor: htn.TickSensor{
				StartedAt:    now,
				TickDuration: time.Duration(s.config.TickDurationMillis) * time.Millisecond,
			},
			TicksPerHour:   60,
			TicksPerMinute: 1,
			OffSet: htn.TimeOfDay{
				Hour:   8,
				Minute: 0,
			},
		},
	}
	return sensors
}

func initializeConditions() htn.Conditions {
	awakeHours := &htn.FuncCondition{
		ConditionName: "awakeHours",
		Evaluator: func(state *htn.Domain) bool {
			sensor, err := state.Sensor("HourOfDay")
			if err != nil {
				log.Errorf("AfterSleep: could not get HourOfDay sensor")
				return false
			}
			hourOfDaySensor := sensor.(*htn.HourOfDaySensor)
			hourOfDay, err := hourOfDaySensor.Get()
			if err != nil {
				log.Errorf("AfterSleep: could not get HourOfDay sensor value")
				return false
			}
			return hourOfDay >= 9 && hourOfDay < 22
		},
	}
	awake := &htn.FuncCondition{
		ConditionName: "Awake",
		Evaluator: func(state *htn.Domain) bool {
			if owner, ok := state.Owner.(*domain.NPC); ok {
				return !owner.Sleeping()
			}
			return false
		},
	}
	bedAvailable := &htn.FuncCondition{
		ConditionName: "BedAvailable",
		Evaluator: func(state *htn.Domain) bool {
			var owner *domain.Character
			if npc, ok := state.Owner.(*domain.NPC); ok {
				owner = &npc.Character
			} else if player, ok := state.Owner.(*domain.Player); ok {
				owner = &player.Character
			}
			room := owner.Room()
			for _, obj := range room.Objects {
				if obj.Type() == domain.ObjectTypeFurniture {
					if subtype, ok := obj.Tags()[domain.ObjectTagSubtype]; ok {
						if subtype == domain.ObjectTagSubtypeBed && !obj.Busy() {
							return true
						}
					}
				}
			}
			return false
		},
	}
	bedNotAvailable := &htn.FuncCondition{
		ConditionName: "BedUnavailable",
		Evaluator: func(state *htn.Domain) bool {
			return !bedAvailable.IsMet(state)
		},
	}
	chairAvailable := &htn.FuncCondition{
		ConditionName: "ChairAvailable",
		Evaluator: func(state *htn.Domain) bool {
			var owner *domain.Character
			if npc, ok := state.Owner.(*domain.NPC); ok {
				owner = &npc.Character
			} else if player, ok := state.Owner.(*domain.Player); ok {
				owner = &player.Character
			}
			room := owner.Room()
			for _, obj := range room.Objects {
				if obj.Type() == domain.ObjectTypeFurniture {
					if subtype, ok := obj.Tags()[domain.ObjectTagSubtype]; ok {
						if subtype == domain.ObjectTagSubtypeChair && !obj.Busy() {
							return true
						}
					}
				}
			}
			return false
		},
	}
	chairNotAvailable := &htn.FuncCondition{
		ConditionName: "ChairUnavailable",
		Evaluator: func(state *htn.Domain) bool {
			return !bedAvailable.IsMet(state)
		},
	}
	conditions := htn.Conditions{
		"AwakeHours": awakeHours,
		"Awake":      awake,
		"AsleepHours": &htn.FuncCondition{
			ConditionName: "AsleepHours",
			Evaluator: func(state *htn.Domain) bool {
				return !awakeHours.IsMet(state)
			},
		},
		"Asleep": &htn.FuncCondition{
			ConditionName: "Asleep",
			Evaluator: func(state *htn.Domain) bool {
				return !awake.IsMet(state)
			},
		},
		"BedAvailable":     bedAvailable,
		"BedUnavailable":   bedNotAvailable,
		"ChairAvailable":   chairAvailable,
		"ChairUnavailable": chairNotAvailable,
		"PlayerEngaged": &htn.FuncCondition{
			ConditionName: "PlayerEngaged",
			Evaluator: func(state *htn.Domain) bool {
				if owner, ok := state.Owner.(*domain.Player); ok {
					return owner.Engaged()
				}
				return false
			},
		},
		"PlayerNotEngaged": &htn.FuncCondition{
			ConditionName: "PlayerNotEngaged",
			Evaluator: func(state *htn.Domain) bool {
				if owner, ok := state.Owner.(*domain.Player); ok {
					return !owner.Engaged()
				}
				return false
			},
		},
		"PlayersEngaged": &htn.FuncCondition{
			ConditionName: "PlayersEngaged",
			Evaluator: func(state *htn.Domain) bool {
				if owner, ok := state.Owner.(*domain.NPC); ok {
					return owner.PlayersEngaged() > 0
				}
				return false
			},
		},
		"PlayersNotEngaged": &htn.ComparisonCondition[int64]{
			ConditionName: "PlayersNotEngaged",
			Comparison:    htn.EQ,
			Value:         0,
			Property:      "PlayersEngaged",
			Comparator: func(value int64, property int64, comparison htn.Comparison) bool {
				return property <= value
			},
		},
		"PlayersInRange": &htn.ComparisonCondition[int64]{
			ConditionName: "PlayersInRange",
			Comparison:    htn.GT,
			Value:         0,
			Property:      "PlayersInRange",
			Comparator: func(value int64, property int64, comparison htn.Comparison) bool {
				return property > value
			},
		},
		"NoPlayersInRange": &htn.ComparisonCondition[int64]{
			ConditionName: "NoPlayersInRange",
			Comparison:    htn.EQ,
			Value:         0,
			Property:      "PlayersInRange",
			Comparator: func(value int64, property int64, comparison htn.Comparison) bool {
				return property <= value
			},
		},
	}
	return conditions
}

func initializeActions() htn.Actions {
	actions := htn.Actions{
		"Wait": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			log.Debugf("%s waiting", owner.Name)
			return nil
		},
		"WakeUp": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			owner.SetSleeping(false)
			dialog := owner.Dialog
			msg := "I'm up."
			wakeUpDialog, ok := dialog["WakeUp"]
			if !ok {
				log.Errorf("No WakeUp dialog for %s", owner.Name)
			} else {
				msg = wakeUpDialog.Text[rand.Intn(len(wakeUpDialog.Text))]
			}
			log.Printf("%s waking up", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventSay,
				Args:      []interface{}{msg},
			})
			return nil
		},
		"SleepOnGround": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			owner.SetSleeping(true)
			dialog := owner.Dialog
			msg := "I'm out."
			sleepDialog, ok := dialog["Sleep"]
			if !ok {
				log.Errorf("No Sleep dialog for %s", owner.Name)
			} else {
				msg = sleepDialog.Text[rand.Intn(len(sleepDialog.Text))]
			}
			log.Printf("%s sleeping", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventSay,
				Args:      []interface{}{msg},
			})
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventAction,
				Args:      []interface{}{"grumbles and lies down on the ground to sleep."},
			})
			return nil
		},
		"UseBed": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			owner.SetSleeping(true)
			room := owner.Room()
			var bed domain.InteractiveObject
			for _, obj := range room.Objects {
				if !obj.Busy() && obj.Type() == domain.ObjectTypeFurniture {
					if subtype, ok := obj.Tags()[domain.ObjectTagSubtype]; ok {
						if subtype == domain.ObjectTagSubtypeBed {
							bed = obj
						}
					}
				}
			}
			dialog := owner.Dialog
			msg := "I'm out."
			sleepDialog, ok := dialog["Sleep"]
			if !ok {
				log.Errorf("No Sleep dialog for %s", owner.Name)
			} else {
				msg = sleepDialog.Text[rand.Intn(len(sleepDialog.Text))]
			}
			log.Printf("%s sleeping", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventSay,
				Args:      []interface{}{msg},
			})
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventAction,
				Args:      []interface{}{fmt.Sprintf("flops down onto the %s to sleep.", bed.Name())},
			})
			return nil
		},
		"UseChair": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			owner.SetSleeping(true)
			room := owner.Room()
			var chair domain.InteractiveObject
			for _, obj := range room.Objects {
				if !obj.Busy() && obj.Type() == domain.ObjectTypeFurniture {
					if subtype, ok := obj.Tags()[domain.ObjectTagSubtype]; ok {
						if subtype == domain.ObjectTagSubtypeChair {
							chair = obj
						}
					}
				}
			}
			dialog := owner.Dialog
			msg := "I'm out."
			sleepDialog, ok := dialog["Sleep"]
			if !ok {
				log.Errorf("No Sleep dialog for %s", owner.Name)
			} else {
				msg = sleepDialog.Text[rand.Intn(len(sleepDialog.Text))]
			}
			log.Printf("%s sleeping", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventSay,
				Args:      []interface{}{msg},
			})
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventAction,
				Args:      []interface{}{fmt.Sprintf("flops down onto the %s to sleep.", chair.Name())},
			})
			return nil
		},
		"Sleep": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			owner.SetSleeping(true)
			dialog := owner.Dialog
			msg := "I'm out."
			sleepDialog, ok := dialog["Sleep"]
			if !ok {
				log.Errorf("No Sleep dialog for %s", owner.Name)
			} else {
				msg = sleepDialog.Text[rand.Intn(len(sleepDialog.Text))]
			}
			log.Printf("%s sleeping", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    event.RoomEventSay,
				Args:      []interface{}{msg},
			})
			// find a location to sleep
			return nil
		},
		"Greet": func(state *htn.Domain) error {
			owner := state.Owner.(*domain.NPC)
			dialog := owner.Dialog
			players := owner.Room().Players
			msg := "{TARGET}! My dawg! Whattup, yo!"
			wakeUpDialog, ok := dialog["Greet"]
			if !ok {
				log.Errorf("No Greet dialog for %s", owner.Name)
			}
			for _, player := range players {
				lastGreeted := owner.PlayerLastGreeted(player)
				if time.Since(lastGreeted) > 5*time.Minute {
					owner.SetPlayerLastGreeted(player, time.Now())
					log.Debugf("%s issuing greeting to %s", owner.Name, player.Name)
					if ok {
						msg = wakeUpDialog.Text[rand.Intn(len(wakeUpDialog.Text))]
					}
					formatted := strings.Replace(msg, "{TARGET}", player.Name, -1)
					owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
						Room:      owner.Room(),
						Character: &owner.Character,
						Action:    event.RoomEventSay,
						Args:      []interface{}{formatted},
					})
				}
			}
			return nil
		},
		"Heal": func(state *htn.Domain) error {
			if npc, ok := state.Owner.(*domain.NPC); ok {
				log.Printf("healing NPC %s", npc.Name)
				npc.Heal()
			} else if player, ok := state.Owner.(*domain.Player); ok {
				log.Printf("healing player %s", player.Name)
				player.Connection.Writeln("You hit that button and inhale, and soon you are fully healed")
				player.Heal()
			} else {
				return errors.New("heal action not implemented for this type")
			}
			return nil
		},
	}
	return actions
}

func (s *Server) startGenerator(spec *domain.GeneratorSpec, initializers generator.NpcInitializers) {
	npcSpec, err := s.loaders.NPCLoader.GetNPC(spec.NPC)
	if err != nil {
		panic(err)
	}
	g := generator.NewNpcGenerator(spec, s.loaders, npcSpec, s.npcs, s.domainGenerator, s.plannerGenerator, initializers)
	go func() {
		err := g.Start()
		if err != nil {
			panic(err)
		}
		defer g.Stop()
	}()
}

type Engine struct {
	Config   *config.Config
	Server   *Server
	EventBus eventbus.Bus
}

func NewEngine(config *config.Config, server *Server, eventBus eventbus.Bus) *Engine {
	return &Engine{
		Config:   config,
		Server:   server,
		EventBus: eventBus,
	}
}
