package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
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
	items, err := os.ReadDir(l.config.AssetPath + "/teams")
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
		team := &domain.Team{}
		data, err := os.ReadFile(l.config.AssetPath + "/teams/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, team)
		if err != nil {
			return nil, err
		}
		l.teams = append(l.teams, team)
	}
	return l.teams, nil
}

func (l *TeamLoader) GetTeam(name string) (*domain.Team, error) {
	teams, err := l.LoadTeams()
	if err != nil {
		return nil, err
	}
	for _, team := range teams {
		if team.Name == name {
			return team, nil
		}
	}
	return nil, nil
}
