package loader

import (
	"context"
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type NPCResolver func(ctx context.Context, id int) (*domain.NPC, error)

type InteractiveObjectLoader struct {
	config         *config.Config
	objects        domain.InteractiveObjects
	actionResolver ActionResolver
	npcResolver    NPCResolver
}

func NewInteractiveObjectLoader(cfg *config.Config, actionResolver ActionResolver, npcResolver NPCResolver) *InteractiveObjectLoader {
	return &InteractiveObjectLoader{
		config:         cfg,
		actionResolver: actionResolver,
		npcResolver:    npcResolver,
		objects:        make(domain.InteractiveObjects),
	}
}

func (l *InteractiveObjectLoader) LoadInteractiveObjects() (domain.InteractiveObjects, error) {
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

		l.objects[spec.Name] = &ActionInteractiveObject{
			BaseInteractiveObject: domain.BaseInteractiveObject{
				ObjectName: spec.Name,
				ObjectType: domain.InteractiveObjectType(spec.ObjectType),
				ActionName: spec.Action,
			},
			ActionResolver: l.actionResolver,
			NPCResolver:    l.npcResolver,
		}
	}
	return l.objects, nil
}

func (l *InteractiveObjectLoader) GetInteractiveObject(name string) (domain.InteractiveObject, error) {
	objs, err := l.LoadInteractiveObjects()
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
	NPCResolver    NPCResolver
}

func (i *ActionInteractiveObject) Interact(gameState *domain.GameState, user *domain.Character, target *string) (string, error) {
	if target != nil {
		log.Printf("%s is using %s action %s on %s", user.Name, i.Name(), i.ActionName, *target)
	} else {
		log.Printf("%s is using %s action %s", user.Name, i.ActionName, i.Name())
	}
	action, err := i.ActionResolver(i.ObjectName)
	if err != nil {
		return "", err
	}
	if user.IsNPC() {
		npc, err := i.NPCResolver(context.Background(), *user.Id)
		if err != nil {
			return "", err
		}
		err = action(npc.Domain)
		if err != nil {
			return "", err
		}
	} else {
		err = action(nil)
		if err != nil {
			return "", err
		}
	}
	return "ok", nil
}

var _ domain.InteractiveObject = &ActionInteractiveObject{}
