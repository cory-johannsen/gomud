package domain

type InteractiveObject interface {
	Interact(state *GameState, user *Character, target *string) (string, error)
}

type InteractiveObjectResolver func(name string) (InteractiveObject, error)

type InteractiveObjects map[string]InteractiveObject

type InteractiveObjectSpec struct {
	Name       string `yaml:"name"`
	ObjectType string `yaml:"objectType"`
	Action     string `yaml:"action"`
}
