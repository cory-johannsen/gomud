package domain

import (
	"fmt"
	"math/rand"
)

type Applier func() interface{}

type Effect struct {
	Name        string
	Description string
	Applier     Applier
}
type Effects []*Effect

type TraitSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effects"`
}

type Trait struct {
	Name        string
	Description string
	Effects     Effects
}

func (t *Trait) Value() interface{} {
	return t
}

func (t *Trait) String() string {
	return fmt.Sprintf("%s - %s", t.Name, t.Description)
}

type Traits []*Trait

func (t Traits) Random() *Trait {
	if len(t) == 0 {
		return nil
	}
	return t[rand.Intn(len(t))]
}

var _ Property = &Trait{}
