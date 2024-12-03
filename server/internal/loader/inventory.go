package loader

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
)

type ItemResolver func(ctx context.Context, id int) (domain.Item, error)

type InventoryLoader struct {
	config    *config.Config
	inventory map[int]domain.Item
}

func NewInventoryLoader(cfg *config.Config) *InventoryLoader {
	return &InventoryLoader{
		config:    cfg,
		inventory: make(map[int]domain.Item),
	}
}

func (l *InventoryLoader) GetItem(id int) (domain.Item, error) {
	return l.inventory[id], nil
}

func (l *InventoryLoader) InventoryFromSpec(ctx context.Context, spec *domain.InventorySpec, resolver ItemResolver) (*domain.Inventory, error) {
	inventory := domain.NewInventory()
	if spec.MainHand != 0 {
		item, err := resolver(ctx, spec.MainHand)
		if err != nil {
			return nil, err
		}
		mainHand := item.(*domain.Weapon)
		err = inventory.EquipMainHand(mainHand)
		if err != nil {
			return nil, err
		}
	}
	if spec.OffHand != 0 {
		item, err := resolver(ctx, spec.OffHand)
		if err != nil {
			return nil, err
		}
		offHand := item.(*domain.Weapon)
		err = inventory.EquipOffHand(offHand)
		if err != nil {
			return nil, err
		}
	}
	if spec.Armor != 0 {
		item, err := resolver(ctx, spec.Armor)
		if err != nil {
			return nil, err
		}
		armor := item.(*domain.Armor)
		err = inventory.EquipArmor(armor)
		if err != nil {
			return nil, err
		}
	}
	for _, id := range spec.Pack {
		item, err := resolver(ctx, id)
		if err != nil {
			return nil, err
		}
		err = inventory.Pack().AddItem(item)
		if err != nil {
			return nil, err
		}
	}
	inventory.AddCash(spec.Cash)
	return inventory, nil
}
