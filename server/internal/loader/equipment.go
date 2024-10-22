package loader

import (
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
	skillLoader   *SkillLoader
	qualityLoader *QualityLoader
}

func NewEquipmentLoader(cfg *config.Config, skillLoader *SkillLoader, qualityLoader *QualityLoader) *EquipmentLoader {
	return &EquipmentLoader{
		config:        cfg,
		weapons:       make(map[string]*domain.Weapon),
		armor:         make(map[string]*domain.Armor),
		skillLoader:   skillLoader,
		qualityLoader: qualityLoader,
	}
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
		log.Debugf("loading weapon %s", item.Name())
		spec := &domain.WeaponSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/equipment/weapons/" + item.Name())
		if err != nil {
			log.Errorf("error reading weapon file %s: %v", item.Name(), err)
			continue
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			log.Errorf("error unmarshalling weapon file %s: %v", item.Name(), err)
			continue
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
