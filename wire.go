//go:build wireinject
// +build wireinject

package gomud

import "github.com/google/wire"
import "github.com/cory-johannsen/gomud/engine"

func InitializeEngine() (*engine.Engine, error) {
	wire.Build(engine.NewConfigFromEnv, engine.NewServer, engine.NewEngine)
	return &engine.Engine{}, nil
}
