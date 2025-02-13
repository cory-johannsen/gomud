package domain

import (
	"fmt"
	"math/rand"
)

type TraitSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effect"`
	GoType      string   `yaml:"goType"`
}

type Trait struct {
	Name        string
	Description string
	Effects     Effects
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
