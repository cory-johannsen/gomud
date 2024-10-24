package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type TalentLoader struct {
	config       *config.Config
	talents      domain.Talents
	effectLoader *EffectLoader
}

func NewTalentLoader(config *config.Config, effectLoader *EffectLoader) *TalentLoader {
	return &TalentLoader{
		config:       config,
		talents:      make(domain.Talents, 0),
		effectLoader: effectLoader,
	}
}

func (l *TalentLoader) LoadTalents() (domain.Talents, error) {
	if l.talents != nil && len(l.talents) > 0 {
		return l.talents, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/talents")
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
		//log.Printf("loading talent %s", item.Name())
		talent := &domain.Talent{}
		data, err := os.ReadFile(l.config.AssetPath + "/talents/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, talent)
		if err != nil {
			return nil, err
		}

		effect, err := l.effectLoader.GetEffect(talent.Name)
		if err != nil {
			log.Errorf("error loading effect for %s: %s", talent.Name, err)
		}
		talent.Effect = effect

		l.talents = append(l.talents, talent)
	}
	return l.talents, nil
}
func (l *TalentLoader) GetTalent(name string) (*domain.Talent, error) {
	talents, err := l.LoadTalents()
	if err != nil {
		return nil, err
	}
	for _, talent := range talents {
		if talent.Name == name {
			return talent, nil
		}
	}
	return nil, nil
}
