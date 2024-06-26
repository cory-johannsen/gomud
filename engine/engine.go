package engine

import (
	"bufio"
	"fmt"
	"github.com/cory-johannsen/gomud/cli"
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
	c.Connection.Write([]byte(fmt.Sprintf("\n%s", c.Handler.Prompt())))
	for {
		netData, err := bufio.NewReader(c.Connection).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		result := fmt.Sprintf("%s\n", c.Handler.Eval(netData))
		written, err := c.Connection.Write([]byte(result))
		if err != nil {
			panic(err)
		}
		if written != len(result) {
			panic("Expected to write the result")
		}
		c.Connection.Write([]byte(fmt.Sprintf("\n%s", c.Handler.Prompt())))
	}
	c.Connection.Close()
}

type Server struct {
	port string
}

func NewServer(config *Config) *Server {
	return &Server{
		port: config.Port,
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
	Config *Config
	Server *Server
}

func NewEngine(config *Config, server *Server) *Engine {
	return &Engine{
		Config: config,
		Server: server,
	}
}
