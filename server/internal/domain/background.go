package domain

type BackgroundSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Traits      []string `yaml:"traits"`
}

type Background struct {
	Name        string
	Description string
	Traits      Traits
}

type Backgrounds []*Background
