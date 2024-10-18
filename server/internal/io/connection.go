package io

type Connection interface {
	Read() string
	Write(string) int
	Writeln(string) int
}
