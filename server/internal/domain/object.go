package domain

import (
	log "github.com/sirupsen/logrus"
	"sync"
)

const (
	ObjectTypeFurniture = "Furniture"

	ObjectTagSubtype = "subtype"

	ObjectTagSubtypeBed   = "Bed"
	ObjectTagSubtypeChair = "Chair"
)

type InteractionCompleteCallback func(obj InteractiveObject, state GameState, user *Character, target *string, result string, err error)

type InteractiveObject interface {
	Name() string
	Type() InteractiveObjectType
	Tags() Tags
	Interact(state GameState, user *Character, target *string, callback InteractionCompleteCallback) (string, error)
	Busy() bool
}

type InteractiveObjectResolver func(name string, npcResolver NPCResolver) (InteractiveObject, error)

type InteractiveObjects map[string]InteractiveObject

type InteractiveObjectType string

type InteractiveObjectSpec struct {
	Name       string `yaml:"name"`
	ObjectType string `yaml:"objectType"`
	Action     string `yaml:"action"`
	Tags       Tags   `yaml:"tags"`
}

type BaseInteractiveObject struct {
	Mutex      sync.Mutex
	ObjectBusy bool
	ObjectName string
	ObjectType InteractiveObjectType
	ObjectTags Tags
	ActionName string
}

func (i *BaseInteractiveObject) Interact(state GameState, user *Character, target *string, callback InteractionCompleteCallback) (string, error) {
	i.Mutex.Lock()
	defer i.Mutex.Unlock()
	i.ObjectBusy = true
	if target != nil {
		log.Printf("%s is using %s action %s on %s", user.Name, i.Name(), i.ActionName, *target)
	} else {
		log.Printf("%s is using %s action %s", user.Name, i.ActionName, i.Name())
	}
	i.ObjectBusy = false
	callback(i, state, user, target, "Ok", nil)
	return "Ok", nil
}

func (i *BaseInteractiveObject) Name() string {
	return i.ObjectName
}

func (i *BaseInteractiveObject) Busy() bool {
	return i.ObjectBusy
}

func (i *BaseInteractiveObject) Tags() Tags {
	return i.ObjectTags
}

func (i *BaseInteractiveObject) Type() InteractiveObjectType {
	return i.ObjectType
}

var _ InteractiveObject = &BaseInteractiveObject{}
