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
	ItemId          int     `json:"ItemId"`
	ItemName        string  `json:"name"`
	ItemDescription string  `json:"description"`
	ItemMass        float64 `json:"mass"`
	ItemCost        float64 `json:"cost"`
}

func (i *BaseItem) Id() int {
	return i.ItemId
}

func (i *BaseItem) SetId(id int) {
	i.ItemId = id
}

func (i *BaseItem) Name() string {
	return i.ItemName
}

func (i *BaseItem) Description() string {
	return i.ItemDescription
}

func (i *BaseItem) Mass() float64 {
	return i.ItemMass
}

func (i *BaseItem) Cost() float64 {
	return i.ItemCost
}

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

var _ Item = &Weapon{}

type Armor struct {
	BaseItem
}

func (a *Armor) Type() ItemType {
	return ItemTypeArmor
}

var _ Item = &Armor{}
