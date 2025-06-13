package loader

import (
	"context"
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/domain/htn"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
	"strings"
	"sync"
)

var defaultObjectType = reflect.TypeOf(&ActionInteractiveObject{})
var actionObjectType = defaultObjectType
var singleUserObjectType = reflect.TypeOf(&SingleUserInteractiveObject{})

var objectTypes map[string]reflect.Type = map[string]reflect.Type{
	"Old busted couch": singleUserObjectType,
	"Shitty bed":       singleUserObjectType,
	"Penjamin":         actionObjectType,
}

type InteractiveObjectLoader struct {
	config         *config.Config
	objects        domain.InteractiveObjects
	actionResolver ActionResolver
}

func NewInteractiveObjectLoader(cfg *config.Config, actionResolver ActionResolver) *InteractiveObjectLoader {
	return &InteractiveObjectLoader{
		config:         cfg,
		actionResolver: actionResolver,
		objects:        make(domain.InteractiveObjects),
	}
}

func (l *InteractiveObjectLoader) LoadInteractiveObjects(npcResolver domain.NPCResolver) (domain.InteractiveObjects, error) {
	if len(l.objects) > 0 {
		return l.objects, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/objects")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") || strings.HasSuffix(item.Name(), "txt") {
			continue
		}
		name := item.Name()
		data, err := os.ReadFile(l.config.AssetPath + "/objects/" + name)
		if err != nil {
			log.Printf("error reading file %s: %s", name, err)
			continue
		}
		spec := &domain.InteractiveObjectSpec{}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Printf("error unmarshalling file %s: %s", name, err)
			continue
		}

		objType, ok := objectTypes[name]
		if !ok {
			objType = defaultObjectType
		}
		var obj domain.InteractiveObject
		switch objType {
		case singleUserObjectType:
			obj = &SingleUserInteractiveObject{
				barrier: sync.WaitGroup{},
				ActionInteractiveObject: ActionInteractiveObject{
					BaseInteractiveObject: domain.BaseInteractiveObject{
						ObjectName: spec.Name,
						ObjectType: domain.InteractiveObjectType(spec.ObjectType),
						ActionName: spec.Action,
					},
					ActionResolver: l.actionResolver,
					NPCResolver:    npcResolver,
				},
			}
		case actionObjectType:
			fallthrough
		case defaultObjectType:
			obj = &ActionInteractiveObject{
				BaseInteractiveObject: domain.BaseInteractiveObject{
					ObjectName: spec.Name,
					ObjectType: domain.InteractiveObjectType(spec.ObjectType),
					ActionName: spec.Action,
				},
				ActionResolver: l.actionResolver,
				NPCResolver:    npcResolver,
			}
		}

		l.objects[spec.Name] = obj
	}
	return l.objects, nil
}

func (l *InteractiveObjectLoader) GetInteractiveObject(name string, npcResolver domain.NPCResolver) (domain.InteractiveObject, error) {
	objs, err := l.LoadInteractiveObjects(npcResolver)
	if err != nil {
		return nil, err
	}
	obj, ok := objs[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("interactive object %s not found", name))
	}
	return obj, nil
}

type ActionInteractiveObject struct {
	domain.BaseInteractiveObject
	ActionResolver ActionResolver
	NPCResolver    domain.NPCResolver
}

func (i *ActionInteractiveObject) Start(gameState domain.GameState, user *domain.Character, target *string, starting domain.InteractionStartingCallback) (htn.Action, error) {
	starting(i, gameState, user, target)
	if target != nil {
		log.Printf("%s is using %s action %s on %s", user.Name, i.Name(), i.ActionName, *target)
	} else {
		log.Printf("%s is using %s action %s", user.Name, i.Name(), i.ActionName)
	}
	action, err := i.ActionResolver.GetAction(i.ActionName)
	if err != nil {
		return nil, err
	}
	return action, nil
}

func (i *ActionInteractiveObject) Apply(gameState domain.GameState, user *domain.Character, action htn.Action) (string, error) {
	if user.IsNPC() {
		npc, err := i.NPCResolver.FetchNPCById(context.Background(), *user.Id)
		if err != nil {
			return "", err
		}
		err = action(npc.Domain)
		if err != nil {
			return "", err
		}
	} else {
		err := action(gameState.Domain())
		if err != nil {
			return "", err
		}
	}
	return "", nil
}

func (i *ActionInteractiveObject) Interact(gameState domain.GameState, user *domain.Character, target *string, starting domain.InteractionStartingCallback, complete domain.InteractionCompleteCallback) (string, error) {
	action, err := i.Start(gameState, user, target, starting)
	if err != nil {
		complete(i, gameState, user, target, "", err)
		return "", err
	}
	result, err := i.Apply(gameState, user, action)
	if err != nil {
		complete(i, gameState, user, target, "", err)
		return "", err
	}
	complete(i, gameState, user, target, result, nil)
	return result, nil
}

var _ domain.InteractiveObject = &ActionInteractiveObject{}

type SingleUserInteractiveObject struct {
	barrier sync.WaitGroup
	ActionInteractiveObject
}

func (i *SingleUserInteractiveObject) Interact(gameState domain.GameState, user *domain.Character, target *string, starting domain.InteractionStartingCallback, complete domain.InteractionCompleteCallback) (string, error) {
	action, err := i.Start(gameState, user, target, starting)
	if err != nil {
		complete(i, gameState, user, target, "", err)
		return "", err
	}
	result, err := i.Apply(gameState, user, action)
	if err != nil {
		complete(i, gameState, user, target, "", err)
		return "", err
	}
	i.barrier.Add(1)
	i.barrier.Wait()

	complete(i, gameState, user, target, result, nil)
	return result, nil
}
func (i *SingleUserInteractiveObject) Release() {
	i.barrier.Done()
}
