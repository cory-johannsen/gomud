package domain

type ItemType string

const (
	ItemTypeWeapon        ItemType = "weapon"
	ItemTypeArmor         ItemType = "armor"
	ItemTypeShield        ItemType = "shield"
	ItemTypeMiscellaneous ItemType = "miscellaneous"
	ItemTypeConsumable    ItemType = "consumable"
)

type ItemSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Type        ItemType `yaml:"type"`
	Mass        int      `yaml:"mass"`
	Encumbrance int      `yaml:"encumbrance"`
	Cost        int      `yaml:"cost"`
}

type Item interface {
	Id() int
	Name() string
	Description() string
	Encumbrance() int
	MassInGrams() int
	Type() ItemType
	Cost() int
	NewInstance(id int) Item
}

type Items []Item

type BaseItem struct {
	ItemId          int    `json:"id"`
	ItemName        string `json:"name"`
	ItemDescription string `json:"description"`
	ItemMass        int    `json:"mass"`
	ItemEncumbrance int    `json:"encumbrance"`
	ItemCost        int    `json:"cost"`
}

func (i *BaseItem) Id() int {
	return i.ItemId
}

func (i *BaseItem) Name() string {
	return i.ItemName
}

func (i *BaseItem) Description() string {
	return i.ItemDescription
}

func (i *BaseItem) MassInGrams() int {
	return i.ItemMass
}

func (i *BaseItem) Encumbrance() int {
	return i.ItemEncumbrance
}

func (i *BaseItem) Cost() int {
	return i.ItemCost
}

type MiscellaneousItem struct {
	BaseItem
}

func (m *MiscellaneousItem) NewInstance(id int) Item {
	return &MiscellaneousItem{
		BaseItem: BaseItem{
			ItemId:          id,
			ItemName:        m.ItemName,
			ItemDescription: m.ItemDescription,
			ItemMass:        m.ItemMass,
			ItemEncumbrance: m.ItemEncumbrance,
			ItemCost:        m.ItemCost,
		},
	}
}

func (m *MiscellaneousItem) Type() ItemType {
	return ItemTypeMiscellaneous
}

var _ Item = &MiscellaneousItem{}

type QualitySpec struct {
	Name       string   `json:"name"`
	Conditions []string `json:"conditions"`
	Effects    []string `json:"effects"`
	Targets    []string `json:"targets"`
}

type Quality struct {
	Name       string
	Conditions []string
	Effects    Effects
	Targets    []string
}
type Qualities []*Quality

type WeaponCategory string

const (
	WeaponCategoryMelee  WeaponCategory = "melee"
	WeaponCategoryRanged WeaponCategory = "ranged"
)

type WeaponLevel string

const (
	WeaponLevelSimple  WeaponLevel = "simple"
	WeaponLevelMartial WeaponLevel = "martial"
)

type WeaponType string

const (
	WeaponTypeBladed    WeaponType = "Bladed"
	WeaponTypeBrawling  WeaponType = "Brawling"
	WeaponTypeExplosive WeaponType = "Explosive"
	WeaponTypeCrushing  WeaponType = "Crushing"
	WeaponTypeGunpowder WeaponType = "Gunpowder"
	WeaponTypeMissile   WeaponType = "Missile"
)

type WeaponHandling string

const (
	WeaponHandlingOneHanded      WeaponHandling = "One-handed"
	WeaponHandlingTwoHanded      WeaponHandling = "Two-handed"
	WeaponHandlingOneOrTwoHanded                = "One or two-handed"
)

type WeaponDistance struct {
	BaseRange int
	StatBonus string
}

type WeaponSpec struct {
	BaseItem
	WeaponCategory  WeaponCategory `json:"category"`
	WeaponLevel     WeaponLevel    `json:"level"`
	WeaponType      WeaponType     `json:"type"`
	WeaponHandling  WeaponHandling `json:"handling"`
	WeaponDistance  string         `json:"distance"`
	WeaponSkill     string         `json:"skill"`
	WeaponLoad      int            `json:"load"`
	WeaponQualities []string       `json:"qualities"`
}

type Weapon struct {
	BaseItem
	category         WeaponCategory
	level            WeaponLevel
	weaponType       WeaponType
	handling         WeaponHandling
	distance         WeaponDistance
	skill            *Skill
	loadActionPoints int
	qualities        Qualities
}

func NewWeapon(spec *WeaponSpec, skill *Skill, qualities Qualities) *Weapon {
	// parse the distance
	return &Weapon{
		BaseItem:         spec.BaseItem,
		category:         spec.WeaponCategory,
		level:            spec.WeaponLevel,
		weaponType:       spec.WeaponType,
		handling:         spec.WeaponHandling,
		distance:         WeaponDistance{},
		skill:            skill,
		loadActionPoints: spec.WeaponLoad,
		qualities:        qualities,
	}
}

func (w *Weapon) Type() ItemType {
	return ItemTypeWeapon
}

func (w *Weapon) Category() WeaponCategory {
	return w.category
}

func (w *Weapon) Level() WeaponLevel {
	return w.level
}

func (w *Weapon) WeaponType() WeaponType {
	return w.weaponType
}

func (w *Weapon) Handling() WeaponHandling {
	return w.handling
}

func (w *Weapon) Distance() WeaponDistance {
	return w.distance
}

func (w *Weapon) Skill() *Skill {
	return w.skill
}

func (w *Weapon) LoadActionPoints() int {
	return w.loadActionPoints
}

func (w *Weapon) Qualities() Qualities {
	return w.qualities
}

func (w *Weapon) NewInstance(id int) Item {
	return &Weapon{
		BaseItem: BaseItem{
			ItemId:          id,
			ItemName:        w.ItemName,
			ItemDescription: w.ItemDescription,
			ItemMass:        w.ItemMass,
			ItemEncumbrance: w.ItemEncumbrance,
			ItemCost:        w.ItemCost,
		},
		category:         w.category,
		level:            w.level,
		weaponType:       w.weaponType,
		handling:         w.handling,
		distance:         w.distance,
		skill:            w.skill,
		loadActionPoints: w.loadActionPoints,
		qualities:        w.qualities,
	}
}

var _ Item = &Weapon{}

type ArmorSpec struct {
	BaseItem
	DamageThresholdModifier int      `json:"damageThresholdModifier"`
	Qualities               []string `json:"qualities"`
}

type Armor struct {
	BaseItem
	damageThresholdModifier int
	qualities               Qualities
}

func (a *Armor) NewInstance(id int) Item {
	return &Armor{
		BaseItem: BaseItem{
			ItemId:          id,
			ItemName:        a.ItemName,
			ItemDescription: a.ItemDescription,
			ItemMass:        a.ItemMass,
			ItemEncumbrance: a.ItemEncumbrance,
			ItemCost:        a.ItemCost,
		},
		damageThresholdModifier: a.damageThresholdModifier,
		qualities:               a.qualities,
	}
}

func (a *Armor) Type() ItemType {
	return ItemTypeArmor
}

func (a *Armor) DamageThresholdModifier() int {
	return a.damageThresholdModifier
}

func (a *Armor) Qualities() Qualities {
	return a.qualities
}

func NewArmor(spec *ArmorSpec, qualities Qualities) *Armor {
	return &Armor{
		BaseItem:                spec.BaseItem,
		damageThresholdModifier: spec.DamageThresholdModifier,
		qualities:               qualities,
	}
}

var _ Item = &Armor{}

type ShieldSpec struct {
	BaseItem
	WeaponHandling WeaponHandling `json:"handling"`
}

type Shield struct {
	BaseItem
	handling WeaponHandling
}

func (s *Shield) NewInstance(id int) Item {
	return &Shield{
		BaseItem: BaseItem{
			ItemId:          id,
			ItemName:        s.ItemName,
			ItemDescription: s.ItemDescription,
			ItemMass:        s.ItemMass,
			ItemEncumbrance: s.ItemEncumbrance,
			ItemCost:        s.ItemCost,
		},
		handling: s.handling,
	}
}

func (s *Shield) Type() ItemType {
	return ItemTypeShield
}

func NewShield(spec *ShieldSpec) *Shield {
	return &Shield{
		BaseItem: spec.BaseItem,
		handling: spec.WeaponHandling,
	}
}

var _ Item = &Shield{}
