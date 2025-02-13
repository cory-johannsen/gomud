package domain

import "fmt"

type BackgroundSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Traits      []string `yaml:"traits"`
}

func (b *BackgroundSpec) String() string {
	return fmt.Sprintf("%s: %s", b.Name, b.Description)
}

var _ Property = &BackgroundSpec{}

func SpecFromBackground(b *Background) *BackgroundSpec {
	traits := make([]string, 0)
	for _, t := range b.Traits {
		traits = append(traits, t.Name)
	}
	return &BackgroundSpec{
		Name:        b.Name,
		Description: b.Description,
		Traits:      traits,
	}
}

type Background struct {
	Name        string
	Description string
	Traits      Traits
}

func (b *Background) String() string {
	return fmt.Sprintf("%s: %s", b.Name, b.Description)
}

var _ Property = &Background{}

type Backgrounds []*Background
