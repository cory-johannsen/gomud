package io

import eventbus "github.com/asaskevich/EventBus"

type Connection interface {
	Read() string
	Write(string) int
	Writeln(string) int
	EventBus() eventbus.Bus
}
