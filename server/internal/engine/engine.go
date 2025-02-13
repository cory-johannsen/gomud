package engine

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/cli"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/event"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"net"
	"strings"
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
	port            string
	db              *storage.Database
	players         *storage.Players
	npcs            *storage.NPCs
	loaders         *loader.Loaders
	playerGenerator *generator.PlayerGenerator
	dispatcher      *cli.Dispatcher
	eventBus        eventbus.Bus
	clock           *event.Clock
}

func NewServer(config *config.Config, db *storage.Database, players *storage.Players, npcs *storage.NPCs, loaders *loader.Loaders,
	playerGenerator *generator.PlayerGenerator, eventBus eventbus.Bus, clock *event.Clock) *Server {
	return &Server{
		port:            config.Port,
		db:              db,
		players:         players,
		npcs:            npcs,
		loaders:         loaders,
		playerGenerator: playerGenerator,
		eventBus:        eventBus,
		clock:           clock,
	}
}

func (s *Server) Start() {
	// preload the assets
	err := s.loaders.Preload()
	if err != nil {
		panic(err)
	}

	// load the NPCs
	npcs, err := s.loaders.NPCLoader.LoadNPCs()
	if err != nil {
		panic(err)
	}
	for _, npcSpec := range npcs {
		npc, err := s.npcs.FetchNPCByName(context.Background(), npcSpec.Name)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			panic(err)
		}
		if npc == nil {
			properties, err := npcSpec.ToProperties(s.loaders)
			if err != nil {
				panic(err)
			}

			newNpc, err := s.npcs.CreateNPC(context.Background(), npcSpec.Name, properties)
			if err != nil {
				panic(err)
			}
			log.Printf("created npc %s\n", newNpc.Name)
			npc = newNpc
		} else {
			log.Printf("npc %s loaded with ID %d\n", npc.Name, npc.Id)
		}
		room := npc.Room()
		if room == nil {
			panic(errors.New(fmt.Sprintf("npc %s room not found", npc.Name)))
		}
		err = room.AddNPC(npc)
		if err != nil {
			panic(err)
		}
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
