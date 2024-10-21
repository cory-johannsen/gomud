package domain

type ItemType string

const (
	ItemTypeWeapon     ItemType = "weapon"
	ItemTypeArmor      ItemType = "armor"
	ItemTypeConsumable ItemType = "consumable"
)

type ItemSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Type        ItemType `yaml:"type"`
	Mass        float64  `yaml:"mass"`
}

type Item interface {
	Id() int
	Name() string
	Description() string
	Mass() float64
	Type() ItemType
}

type Items []Item

type BaseItem struct {
	id          int
	name        string
	description string
	mass        float64
}

func (i *BaseItem) Id() int {
	return i.id
}

func (i *BaseItem) Name() string {
	return i.name
}

func (i *BaseItem) Description() string {
	return i.description
}

func (i *BaseItem) Mass() float64 {
	return i.mass
}

type Weapon struct {
	BaseItem
}

func (w *Weapon) Type() ItemType {
	return ItemTypeWeapon
}

var _ Item = &Weapon{}

type Armor struct {
	BaseItem
}

func (a *Armor) Type() ItemType {
	return ItemTypeArmor
}

var _ Item = &Armor{}
