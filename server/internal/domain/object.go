package domain

import log "github.com/sirupsen/logrus"

type InteractiveObject interface {
	Name() string
	Interact(state *GameState, user *Character, target *string) (string, error)
}

type InteractiveObjectResolver func(name string, npcResolver NPCResolver) (InteractiveObject, error)

type InteractiveObjects map[string]InteractiveObject

type InteractiveObjectType string

type InteractiveObjectSpec struct {
	Name       string `yaml:"name"`
	ObjectType string `yaml:"objectType"`
	Action     string `yaml:"action"`
}

type BaseInteractiveObject struct {
	ObjectName string
	ObjectType InteractiveObjectType
	ActionName string
}

func (i *BaseInteractiveObject) Interact(state *GameState, user *Character, target *string) (string, error) {
	if target != nil {
		log.Printf("%s is using %s action %s on %s", user.Name, i.Name(), i.ActionName, *target)
	} else {
		log.Printf("%s is using %s action %s", user.Name, i.ActionName, i.Name())
	}
	return "Ok", nil
}

func (i *BaseInteractiveObject) Name() string {
	return i.ObjectName
}

var _ InteractiveObject = &BaseInteractiveObject{}
