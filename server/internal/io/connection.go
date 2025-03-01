package io

import (
	eventbus "github.com/asaskevich/EventBus"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
)

type Connection interface {
	Read() string
	Write(string) int
	Writeln(string) int
	EventBus() eventbus.Bus
	Sensors() htn.Sensors
}
