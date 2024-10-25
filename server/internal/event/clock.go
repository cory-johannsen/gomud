package event

import (
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/config"
	log "github.com/sirupsen/logrus"
	"time"
)

const TickChannel = "tick"

type TickEvent int64
type TickHandler func(int64)

type Clock struct {
	config   *config.Config
	eventBus eventbus.Bus
	done     chan bool
}

func NewClock(eventBus eventbus.Bus, config *config.Config) *Clock {
	return &Clock{
		config:   config,
		eventBus: eventBus,
		done:     make(chan bool),
	}
}

func (c *Clock) Start() {
	start := time.Now()
	ticker := time.NewTicker(time.Second) //time.Duration(c.config.TickDurationMillis) * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-c.done:
			return
		case t := <-ticker.C:
			ticks := int64(t.Sub(start).Seconds())
			log.Debugf("tick: %d", ticks)
			c.eventBus.Publish(TickChannel, ticks)
		}
	}
}

func (c *Clock) Stop() {
	c.done <- true
}
