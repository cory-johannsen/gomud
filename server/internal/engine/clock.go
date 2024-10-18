package engine

import (
	"github.com/cory-johannsen/gomud/internal/config"
	goeventbus "github.com/stanipetrosyan/go-eventbus"
	"time"
)

type Clock struct {
	config   *config.Config
	eventBus goeventbus.EventBus
	channel  string
	done     chan bool
}

func NewClock(eventBus goeventbus.EventBus, config *config.Config) *Clock {
	return &Clock{
		config:   config,
		eventBus: eventBus,
		channel:  "tick",
		done:     make(chan bool),
	}
}

func (c *Clock) Start() {
	ticker := time.NewTicker(time.Duration(c.config.TickDurationMillis) * time.Millisecond)
	defer ticker.Stop()
	channel := c.eventBus.Channel(c.channel)
	go func() {
		for {
			select {
			case <-c.done:
				return
			case t := <-ticker.C:
				message := goeventbus.CreateMessage().SetBody(t)
				channel.Publisher().Publish(message)
			}
		}
	}()
}

func (c *Clock) Stop() {
	c.done <- true
}
