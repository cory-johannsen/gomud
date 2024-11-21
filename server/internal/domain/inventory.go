package domain

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

var PackFullError = errors.New("Pack is full")

type Pack struct {
	items   Items
	maxSize int
}

func (p *Pack) Items() Items {
	return p.items
}

func (p *Pack) AddItem(item Item) error {
	if len(p.items) < p.maxSize {
		p.items = append(p.items, item)
	} else {
		return PackFullError
	}
	return nil
}

func (p *Pack) RemoveItem(item Item) error {
	for i, packItem := range p.items {
		if packItem == item {
			p.items = append(p.items[:i], p.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found in pack")
}

type Inventory struct {
	mainHand *Weapon
	offHand  *Weapon
	armor    *Armor
	pack     *Pack
	cash     int
}

func (i *Inventory) MainHand() *Weapon {
	return i.mainHand
}

func (i *Inventory) OffHand() *Weapon {
	return i.offHand
}

func (i *Inventory) Armor() *Armor {
	return i.armor
}

func (i *Inventory) Pack() *Pack {
	return i.pack
}

func (i *Inventory) Cash() int {
	return i.cash
}
func (i *Inventory) AddCash(amount int) {
	i.cash += amount
}
func (i *Inventory) RemoveCash(amount int) error {
	if i.cash < amount {
		return errors.New("not enough cash")
	}
	i.cash -= amount
	return nil
}

func (i *Inventory) EquipMainHand(weapon *Weapon) error {
	if weapon == nil {
		return errors.New("weapon is nil")
	}
	if i.mainHand != nil {
		err := i.pack.AddItem(i.mainHand)
		if err != nil {
			if errors.Is(err, PackFullError) {
				return err
			}
			log.Errorf("error adding main hand weapon to pack: %v", err)
			return err
		}
	}
	i.mainHand = weapon
	return nil
}

func (i *Inventory) EquipOffHand(weapon *Weapon) error {
	if weapon == nil {
		return errors.New("weapon is nil")
	}
	if i.offHand != nil {
		err := i.pack.AddItem(i.offHand)
		if err != nil {
			if errors.Is(err, PackFullError) {
				return err
			}
			log.Errorf("error adding off hand weapon to pack: %v", err)
			return err
		}
	}
	i.offHand = weapon
	return nil
}

func (i *Inventory) EquipArmor(armor *Armor) error {
	if armor == nil {
		return errors.New("armor is nil")
	}
	if i.armor != nil {
		err := i.pack.AddItem(i.armor)
		if err != nil {
			if errors.Is(err, PackFullError) {
				return err
			}
			log.Errorf("error adding armor to pack: %v", err)
			return err
		}
	}
	i.armor = armor
	return nil
}

func (i *Inventory) UnequipMainHand() (*Weapon, error) {
	if i.mainHand == nil {
		return nil, errors.New("main hand weapon is nil")
	}
	weapon := i.mainHand
	i.mainHand = nil
	return weapon, nil
}

func (i *Inventory) UnequipOffHand() (*Weapon, error) {
	if i.offHand == nil {
		return nil, errors.New("off hand weapon is nil")
	}
	weapon := i.offHand
	i.offHand = nil
	return weapon, nil
}

func (i *Inventory) UnequipArmor() (*Armor, error) {
	if i.armor == nil {
		return nil, errors.New("armor is nil")
	}
	armor := i.armor
	i.armor = nil
	return armor, nil
}

func (i *Inventory) Value() interface{} {
	return i
}

func (i *Inventory) String() string {
	return "Inventory"
}

type InventorySpec struct {
	MainHand int   `yaml:"main_hand"`
	OffHand  int   `yaml:"off_hand"`
	Armor    int   `yaml:"armor"`
	Pack     []int `yaml:"pack"`
	Cash     int   `yaml:"cash"`
}

func SpecFromInventory(inventory *Inventory) *InventorySpec {
	pack := make([]int, 0)
	for _, item := range inventory.pack.items {
		pack = append(pack, item.Id())
	}
	var mainHand = 0
	if inventory.mainHand != nil {
		mainHand = inventory.mainHand.Id()
	}
	var offHand = 0
	if inventory.offHand != nil {
		offHand = inventory.offHand.Id()
	}
	var armor = 0
	if inventory.armor != nil {
		armor = inventory.armor.Id()
	}
	return &InventorySpec{
		MainHand: mainHand,
		OffHand:  offHand,
		Armor:    armor,
		Pack:     pack,
		Cash:     inventory.cash,
	}
}

func NewInventory() *Inventory {
	return &Inventory{
		pack: &Pack{
			maxSize: 25,
		},
	}
}

var _ Property = &Inventory{}
