package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
)

type TeamLoader struct {
	config *config.Config
	teams  domain.Teams
}

func NewTeamLoader(cfg *config.Config) *TeamLoader {
	return &TeamLoader{
		config: cfg,
		teams:  make(domain.Teams, 0),
	}
}

func (l *TeamLoader) LoadTeams() (domain.Teams, error) {
	if l.teams != nil && len(l.teams) > 0 {
		return l.teams, nil
	}
	teams := make(domain.Teams, 0)
	items, err := os.ReadDir(l.config.AssetPath + "/teams")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			log.Printf("skipping template file %s", item.Name())
			continue
		}
		team := &domain.Team{}
		data, err := os.ReadFile(l.config.AssetPath + "/teams/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, team)
		if err != nil {
			return nil, err
		}
		teams = append(teams, team)
	}
	l.teams = teams
	return teams, nil
}
