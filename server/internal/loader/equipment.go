package loader

import (
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type EquipmentLoader struct {
	config        *config.Config
	weapons       map[string]*domain.Weapon
	armor         map[string]*domain.Armor
	shields       map[string]*domain.Shield
	misc          map[string]domain.Item
	skillLoader   *SkillLoader
	qualityLoader *QualityLoader
}

func NewEquipmentLoader(cfg *config.Config, skillLoader *SkillLoader, qualityLoader *QualityLoader) *EquipmentLoader {
	return &EquipmentLoader{
		config:        cfg,
		weapons:       make(map[string]*domain.Weapon),
		armor:         make(map[string]*domain.Armor),
		shields:       make(map[string]*domain.Shield),
		misc:          make(map[string]domain.Item),
		skillLoader:   skillLoader,
		qualityLoader: qualityLoader,
	}
}

func (l *EquipmentLoader) LoadEquipment() (map[string]domain.Item, error) {
	weapons, err := l.LoadWeapons()
	if err != nil {
		return nil, err
	}
	armor, err := l.LoadArmor()
	if err != nil {
		return nil, err
	}
	shields, err := l.LoadShields()
	if err != nil {
		return nil, err
	}
	misc, err := l.LoadMisc()
	if err != nil {
		return nil, err
	}
	items := make(map[string]domain.Item)
	for k, v := range weapons {
		items[k] = v
	}
	for k, v := range armor {
		items[k] = v
	}
	for k, v := range shields {
		items[k] = v
	}
	for k, v := range misc {
		items[k] = v
	}
	return items, nil
}

func (l *EquipmentLoader) LoadWeapons() (map[string]*domain.Weapon, error) {
	if len(l.weapons) > 0 {
		return l.weapons, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/equipment/weapons")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		log.Printf("loading weapon %s", item.Name())
		spec := &domain.WeaponSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/weapons/" + item.Name())
		if err != nil {
			log.Errorf("error reading weapon file %s: %v", item.Name(), err)
			continue
		}
		m := map[string]interface{}{}
		err = yaml.Unmarshal(data, m)
		if err != nil {
			log.Errorf("error unmarshalling weapon file %s: %v", item.Name(), err)
			continue
		}
		spec.ItemName = m["name"].(string)
		spec.ItemDescription = m["description"].(string)
		spec.ItemEncumbrance = m["encumbrance"].(int)
		spec.ItemCost = m["cost"].(int)
		spec.ItemMass = m["mass"].(int)
		spec.WeaponSkill = m["skill"].(string)
		spec.WeaponCategory = domain.WeaponCategory(m["category"].(string))
		spec.WeaponType = domain.WeaponType(m["type"].(string))
		spec.WeaponLoad = m["load"].(int)
		spec.WeaponLevel = domain.WeaponLevel(m["level"].(string))
		spec.WeaponHandling = domain.WeaponHandling(m["handling"].(string))
		spec.WeaponDistance = m["distance"].(string)
		spec.WeaponQualities = make([]string, 0)
		for _, quality := range m["qualities"].([]interface{}) {
			spec.WeaponQualities = append(spec.WeaponQualities, quality.(string))
		}

		skill, err := l.skillLoader.GetSkill(spec.WeaponSkill)
		if err != nil {
			log.Errorf("error getting skill %s: %v", spec.WeaponSkill, err)
			continue
		}
		qualities := make([]*domain.Quality, 0)
		for _, qualityName := range spec.WeaponQualities {
			quality, err := l.qualityLoader.GetQuality(qualityName)
			if err != nil {
				log.Errorf("error getting quality %s: %v", qualityName, err)
				continue
			}
			qualities = append(qualities, quality)
		}
		weapon := domain.NewWeapon(spec, skill, qualities)

		l.weapons[weapon.Name()] = weapon
	}
	return l.weapons, nil
}

func (l *EquipmentLoader) LoadArmor() (map[string]*domain.Armor, error) {
	if len(l.armor) > 0 {
		return l.armor, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/equipment/armor")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		log.Printf("loading armor %s", item.Name())
		spec := &domain.ArmorSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/armor/" + item.Name())
		if err != nil {
			log.Errorf("error reading armor file %s: %v", item.Name(), err)
			continue
		}

		m := make(map[string]interface{})
		err = yaml.Unmarshal(data, m)
		if err != nil {
			log.Errorf("error unmarshalling armor file %s: %v", item.Name(), err)
			continue
		}

		spec.ItemName = m["name"].(string)
		spec.ItemDescription = m["description"].(string)
		spec.ItemEncumbrance = m["encumbrance"].(int)
		spec.ItemCost = m["cost"].(int)
		spec.ItemMass = m["mass"].(int)
		spec.DamageThresholdModifier = m["damageThresholdModifier"].(int)
		spec.Qualities = make([]string, 0)
		specQualities, ok := m["qualities"]
		if ok && specQualities != nil {
			for _, quality := range specQualities.([]interface{}) {
				spec.Qualities = append(spec.Qualities, quality.(string))
			}
		}

		qualities := make([]*domain.Quality, 0)
		for _, qualityName := range spec.Qualities {
			quality, err := l.qualityLoader.GetQuality(qualityName)
			if err != nil {
				log.Errorf("error getting quality %s: %v", qualityName, err)
				continue
			}
			qualities = append(qualities, quality)
		}
		armor := domain.NewArmor(spec, qualities)
		l.armor[armor.Name()] = armor
	}
	return l.armor, nil
}

func (l *EquipmentLoader) LoadShields() (map[string]*domain.Shield, error) {
	if len(l.shields) > 0 {
		return l.shields, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/equipment/shields")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		log.Printf("loading shield %s", item.Name())
		spec := &domain.ShieldSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/shields/" + item.Name())
		if err != nil {
			log.Errorf("error reading shield file %s: %v", item.Name(), err)
			continue
		}

		m := make(map[string]interface{})
		err = yaml.Unmarshal(data, m)
		if err != nil {
			log.Errorf("error unmarshalling shield file %s: %v", item.Name(), err)
			continue
		}

		spec.ItemName = m["name"].(string)
		spec.ItemDescription = m["description"].(string)
		spec.ItemEncumbrance = m["encumbrance"].(int)
		spec.ItemCost = m["cost"].(int)
		spec.ItemMass = m["mass"].(int)
		spec.WeaponHandling = domain.WeaponHandling(m["handling"].(string))

		shield := domain.NewShield(spec)
		l.shields[shield.Name()] = shield
	}
	return l.shields, nil
}

func (l *EquipmentLoader) LoadMisc() (map[string]domain.Item, error) {
	if len(l.misc) > 0 {
		return l.misc, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/equipment/misc")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		log.Printf("loading miscellaneous item %s", item.Name())
		spec := &domain.ItemSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/misc/" + item.Name())
		if err != nil {
			log.Errorf("error reading misc file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling misc file %s: %v", item.Name(), err)
			continue
		}
		i := &domain.MiscellaneousItem{
			BaseItem: domain.BaseItem{
				ItemName:        spec.Name,
				ItemDescription: spec.Description,
				ItemMass:        spec.Mass,
				ItemEncumbrance: spec.Encumbrance,
				ItemCost:        spec.Cost,
			},
		}
		l.misc[i.Name()] = i
	}
	return l.misc, nil
}

func (l *EquipmentLoader) ItemFromName(name string) (domain.Item, error) {
	weapons, err := l.LoadWeapons()
	if err != nil {
		return nil, err
	}
	weapon, ok := weapons[name]
	if ok {
		return weapon, nil
	}
	armor, err := l.LoadArmor()
	if err != nil {
		return nil, err
	}
	a, ok := armor[name]
	if ok {
		return a, nil
	}
	misc, err := l.LoadMisc()
	if err != nil {
		return nil, err
	}
	m, ok := misc[name]
	if ok {
		return m, nil
	}
	shields, err := l.LoadShields()
	if err != nil {
		return nil, err
	}
	s, ok := shields[name]
	if ok {
		return s, nil
	}
	log.Printf("item %s not found", name)
	return nil, errors.New(fmt.Sprintf("item %s not found", name))
}
