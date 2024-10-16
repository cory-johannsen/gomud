package engine

import (
	"bufio"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/cli"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	log "github.com/sirupsen/logrus"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
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

func NewState(player *domain.Player) cli.State {
	return &State{
		player: player,
		vars:   make(map[string]domain.Property),
	}
}

type ClientConnection struct {
	Connection net.Conn
}

func NewClientConnection(conn net.Conn) *ClientConnection {
	return &ClientConnection{
		Connection: conn,
	}
}

func (c *ClientConnection) Read() string {
	netData, err := bufio.NewReader(c.Connection).ReadString('\n')
	if err != nil {
		log.Println(err)
		return ""
	}
	return strings.TrimSuffix(netData, "\r\n")
}

func (c *ClientConnection) Write(data string) int {
	written, err := c.Connection.Write([]byte(data))
	if err != nil {
		log.Println(err)
	}
	return written
}

func (c *ClientConnection) Writeln(data string) int {
	return c.Write(fmt.Sprintf("%s\n", data))
}

type Client struct {
	Connection net.Conn
	EventBus   goeventbus.EventBus
	Dispatcher *cli.Dispatcher
}

func NewClient(players *storage.Players, generator *generator.PlayerGenerator, teams *loader.TeamLoader, rooms *loader.RoomLoader,
	conn net.Conn, eventBus goeventbus.EventBus) *Client {
	dispatcher := cli.NewDispatcher(NewState, players, generator, teams, rooms, NewClientConnection(conn), eventBus)
	return &Client{
		Connection: conn,
		Dispatcher: dispatcher,
	}
}

func (c *Client) Connect() {
	log.Printf("Serving client %s\n", c.Connection.RemoteAddr().String())
	prompt := fmt.Sprintf("\n%s", c.Dispatcher.Prompt())
	_, err := c.Connection.Write([]byte(cli.WelcomeMessage))
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
			return
		}
		if strings.TrimSuffix(result, "\n") == cli.QuitMessage {
			break
		}
		prompt = fmt.Sprintf("\n%s", c.Dispatcher.Prompt())
		_, err = c.Connection.Write([]byte(prompt))
		if err != nil {
			log.Println(err)
			return
		}
	}
	log.Printf("Client %s disconnected\n", c.Connection.RemoteAddr().String())
	c.Connection.Close()
}

type Server struct {
	port            string
	db              *storage.Database
	players         *storage.Players
	loaders         *loader.Loaders
	playerGenerator *generator.PlayerGenerator
	dispatcher      *cli.Dispatcher
	eventBus        goeventbus.EventBus
	clock           *Clock
}

func NewServer(config *config.Config, db *storage.Database, players *storage.Players, loaders *loader.Loaders,
	playerGenerator *generator.PlayerGenerator, eventBus goeventbus.EventBus, clock *Clock) *Server {
	return &Server{
		port:            config.Port,
		db:              db,
		players:         players,
		loaders:         loaders,
		playerGenerator: playerGenerator,
		eventBus:        eventBus,
		clock:           clock,
	}
}

func (s *Server) Start() {
	log.Println("Pre-loading assets")
	_, err := s.loaders.AlignmentLoader.LoadAlignments()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.AppearanceLoader.LoadTattooLocations()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.AppearanceLoader.LoadTattoos()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.AppearanceLoader.LoadDistinguishingMarks()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.AppearanceLoader.LoadDrawbacks()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.ArchetypeLoader.LoadArchetypes()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.BackgroundLoader.LoadBackgrounds()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.InjuryLoader.LoadInjuries()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.SkillLoader.LoadSkills()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.TalentLoader.LoadTalents()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.TraitLoader.LoadTraits()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.JobLoader.LoadJobs()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.TeamLoader.LoadTeams()
	if err != nil {
		panic(err)
	}
	_, err = s.loaders.RoomLoader.LoadRooms()
	if err != nil {
		panic(err)
	}

	log.Printf("Starting server on port %s\n", s.port)
	l, err := net.Listen("tcp4", fmt.Sprintf(":%s", s.port))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	s.clock.Start()
	defer s.clock.Stop()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		client := NewClient(s.players, s.playerGenerator, s.loaders.TeamLoader, s.loaders.RoomLoader, c, s.eventBus)
		go client.Connect()
	}
}

type Engine struct {
	Config   *config.Config
	Server   *Server
	EventBus goeventbus.EventBus
}

func NewEngine(config *config.Config, server *Server, eventBus goeventbus.EventBus) *Engine {
	return &Engine{
		Config:   config,
		Server:   server,
		EventBus: eventBus,
	}
}
