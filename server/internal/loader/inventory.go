package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
)

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

func (l *InventoryLoader) LoadInventory() (map[int]domain.Item, error) {
	if len(l.inventory) > 0 {
		return l.inventory, nil
	}
	return l.inventory, nil
}

func (l *InventoryLoader) GetItem(id int) (domain.Item, error) {
	if len(l.inventory) == 0 {
		_, err := l.LoadInventory()
		if err != nil {
			return nil, err
		}
	}
	return l.inventory[id], nil
}

func (l *InventoryLoader) InventoryFromSpec(spec *domain.InventorySpec) (*domain.Inventory, error) {
	inventory := domain.NewInventory()
	if spec.MainHand != 0 {
		item, err := l.GetItem(spec.MainHand)
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
		item, err := l.GetItem(spec.OffHand)
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
		item, err := l.GetItem(spec.Armor)
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
		item, err := l.GetItem(id)
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
