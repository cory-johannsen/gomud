package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type SkillLoader struct {
	config *config.Config
	skills domain.Skills
}

func NewSkillLoader(config *config.Config) *SkillLoader {
	return &SkillLoader{
		config: config,
		skills: make(domain.Skills, 0),
	}
}

func (l *SkillLoader) Load() (domain.Skills, error) {
	if l.skills != nil && len(l.skills) > 0 {
		return l.skills, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/skills")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") ||
			strings.HasSuffix(item.Name(), ".txt") ||
			strings.HasSuffix(item.Name(), ".sh") {
			continue
		}
		log.Printf("loading skill %s", item.Name())
		skill := &domain.Skill{}
		data, err := os.ReadFile(l.config.AssetPath + "/skills/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, skill)
		if err != nil {
			return nil, err
		}
		l.skills = append(l.skills, skill)
	}
	return l.skills, nil
}
