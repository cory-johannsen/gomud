package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
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

func (l *SkillLoader) LoadSkills() (domain.Skills, error) {
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
func (l *SkillLoader) GetSkill(name string) (*domain.Skill, error) {
	skills, err := l.LoadSkills()
	if err != nil {
		return nil, err
	}
	for _, skill := range skills {
		if skill.Name == name {
			return skill, nil
		}
	}
	return nil, nil
}
