package engine

import (
	"bufio"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/cli"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/generator"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/cory-johannsen/gomud/internal/storage"
	"github.com/openengineer/go-repl"
	"log"
	"net"
)

type State struct {
}

type Client struct {
	Connection net.Conn
	Handler    *cli.CommandHandler
}

func NewClient(conn net.Conn) *Client {
	handler := &cli.CommandHandler{}
	handler.R = repl.NewRepl(handler)
	return &Client{
		Connection: conn,
		Handler:    handler,
	}
}

func (c *Client) Connect() {
	log.Printf("Serving client %s\n", c.Connection.RemoteAddr().String())
	prompt := fmt.Sprintf("\n%s", c.Handler.Prompt())
	written, err := c.Connection.Write([]byte(prompt))
	if err != nil {
		panic(err)
	}
	if written != len(prompt) {
		log.Printf("Expected to write %d bytes, wrote %d", len(c.Handler.Prompt()), written)
		return
	}
	for {
		netData, err := bufio.NewReader(c.Connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		result := fmt.Sprintf("%s\n", c.Handler.Eval(netData))
		written, err := c.Connection.Write([]byte(result))
		if err != nil {
			log.Println(err)
			return
		}
		if written != len(result) {
			log.Printf("Expected to write %d bytes, wrote %d", len(result), written)
			return
		}
		prompt = fmt.Sprintf("\n%s", c.Handler.Prompt())
		written, err = c.Connection.Write([]byte(prompt))
		if err != nil {
			log.Println(err)
			return
		}
		if written != len(result) {
			log.Printf("Expected to write %d bytes, wrote %d", len(prompt), written)
			return
		}
	}
	c.Connection.Close()
}

type Server struct {
	port             string
	db               *storage.Database
	players          *storage.Players
	appearanceLoader *loader.AppearanceLoader
	playerGenerator  *generator.PlayerGenerator
}

func NewServer(config *config.Config, db *storage.Database, players *storage.Players, appearanceLoader *loader.AppearanceLoader,
	playerGenerator *generator.PlayerGenerator) *Server {
	return &Server{
		port:             config.Port,
		db:               db,
		players:          players,
		appearanceLoader: appearanceLoader,
		playerGenerator:  playerGenerator,
	}
}

func (s *Server) Start() {
	l, err := net.Listen("tcp4", fmt.Sprintf(":%s", s.port))
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		client := NewClient(c)
		go client.Connect()
	}
}

type Engine struct {
	Config *config.Config
	Server *Server
}

func NewEngine(config *config.Config, server *Server) *Engine {
	return &Engine{
		Config: config,
		Server: server,
	}
}
