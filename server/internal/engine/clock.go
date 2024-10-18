package engine

import (
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/config"
	"time"
)

type Clock struct {
	config   *config.Config
	eventBus eventbus.Bus
	channel  string
	done     chan bool
}

func NewClock(eventBus eventbus.Bus, config *config.Config) *Clock {
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
	go func() {
		for {
			select {
			case <-c.done:
				return
			case t := <-ticker.C:
				c.eventBus.Publish(c.channel, t)
			}
		}
	}()
}

func (c *Clock) Stop() {
	c.done <- true
}
