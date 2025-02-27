package engine

import (
	"bufio"
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/cli"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	"net"
	"strings"
	"time"
)

type State struct {
	player *domain.Player
	vars   map[string]domain.Property
}

func (s *State) Player() *domain.Player {
	return s.player
}

func (s *State) Property(key string) domain.Property {
	return s.vars[key]
}

func NewState(player *domain.Player) domain.State {
	return &State{
		player: player,
		vars:   make(map[string]domain.Property),
	}
}

type ClientConnection struct {
	Connection net.Conn
	eventBus   eventbus.Bus
}

func NewClientConnection(conn net.Conn, bus eventbus.Bus) *ClientConnection {
	return &ClientConnection{
		Connection: conn,
		eventBus:   bus,
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

type Client struct {
	Connection net.Conn
	EventBus   eventbus.Bus
	Dispatcher *cli.Dispatcher
}

func NewClient(players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, rooms *loader.RoomLoader,
	skills *loader.SkillLoader, conn net.Conn, eventBus eventbus.Bus) *Client {
	dispatcher := cli.NewDispatcher(NewState, players, generator, teams, rooms, skills, NewClientConnection(conn, eventBus), eventBus)
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
	stateGenerator   *generator.StateGenerator
	plannerGenerator *generator.PlannerGenerator
	dispatcher       *cli.Dispatcher
	eventBus         eventbus.Bus
	clock            *event.Clock
}

func NewServer(config *config.Config, db *storage.Database, players *storage.Players, npcs *storage.NPCs, loaders *loader.Loaders,
	playerGenerator *generator.PlayerGenerator, stateGenerator *generator.StateGenerator, plannerGenerator *generator.PlannerGenerator, eventBus eventbus.Bus, clock *event.Clock) *Server {
	return &Server{
		port:             config.Port,
		db:               db,
		players:          players,
		npcs:             npcs,
		loaders:          loaders,
		playerGenerator:  playerGenerator,
		stateGenerator:   stateGenerator,
		plannerGenerator: plannerGenerator,
		eventBus:         eventBus,
		clock:            clock,
	}
}

func (s *Server) Start() {
	// Define the non-asset conditions
	conditions := htn.Conditions{
		"AfterWake": &htn.ComparisonCondition[int64]{
			ConditionName: "AfterAwake",
			Comparison:    htn.GTE,
			Value:         8,
			Property:      "HourOfDay",
			Comparator: func(value int64, property int64, comparison htn.Comparison) bool {
				return property >= value
			},
		},
		"BeforeSleep": &htn.ComparisonCondition[int64]{
			ConditionName: "BeforeSleep",
			Comparison:    htn.GTE,
			Value:         10,
			Property:      "HourOfDay",
			Comparator: func(value int64, property int64, comparison htn.Comparison) bool {
				return property <= value
			},
		},
		"PlayerNotEngaged": &htn.ComparisonCondition[int64]{
			ConditionName: "PlayerNotEngaged",
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
				return property <= value
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
		"IsPlayer": &htn.FuncCondition{
			ConditionName: "IsPlayer",
			Evaluator: func(state *htn.State) bool {
				// TODO: fetch the current customer for the vendor and check if they are the player
				return true
			},
		},
	}

	actions := htn.Actions{
		"Wait": func(state *htn.State) error {
			owner := state.Owner.(*domain.NPC)
			log.Printf("%s waiting", owner.Name)
			return nil
		},
		"WakeUp": func(state *htn.State) error {
			owner := state.Owner.(*domain.NPC)
			log.Printf("%s waking up", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    "say",
				Args:      []interface{}{"Mornin' dawgs!"},
			})
			return nil
		},
		"Sleep": func(state *htn.State) error {
			owner := state.Owner.(*domain.NPC)
			log.Printf("%s sleeping", owner.Name)
			owner.EventBus.Publish(event.RoomChannel, &domain.RoomEvent{
				Room:      owner.Room(),
				Character: &owner.Character,
				Action:    "say",
				Args:      []interface{}{"Imma crash now, feelin' busted."},
			})
			return nil
		},
		"Greet": func(state *htn.State) error {
			owner := state.Owner.(*domain.NPC)
			log.Printf("%s issuing greeting", owner.Name)
			return nil
		},
	}

	now := time.Now()
	sensors := htn.Sensors{
		"HourOfDay": &htn.HourOfDaySensor{
			TickSensor: htn.TickSensor{
				StartedAt:    now,
				TickDuration: 10 * time.Second,
			},
		},
	}

	// preload the assets
	err := s.loaders.Preload(conditions, actions, sensors)
	if err != nil {
		panic(err)
	}

	// start the generators
	specs, err := s.loaders.GeneratorLoader.LoadGenerators()
	if err != nil {
		panic(err)
	}
	for _, spec := range specs {
		npcSpec, err := s.loaders.NPCLoader.GetNPC(spec.NPC)
		if err != nil {
			panic(err)
		}
		g := generator.NewNPCGenerator(spec, s.loaders, npcSpec, s.npcs, s.stateGenerator, s.plannerGenerator)
		go func() {
			err := g.Start()
			if err != nil {
				panic(err)
			}
			defer g.Stop()
		}()
	}

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
		client := NewClient(s.players, s.playerGenerator, s.loaders.TeamLoader, s.loaders.RoomLoader, s.loaders.SkillLoader, c, s.eventBus)
		go client.Connect()
	}
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
