package cli

import (
	"context"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
)

type SkillsHandler struct {
	stateProvider domain.StateProvider
	skills        *loader.SkillLoader
}

func (s *SkillsHandler) Handle(ctx context.Context, args []string) (string, error) {
	skills, err := s.skills.LoadSkills()
	if err != nil {
		return "", err
	}
	return domain.RankSkillsString(s.stateProvider().Player().Skills(skills)), nil
}

func (s *SkillsHandler) Help(args []string) string {
	return "view your Skills"
}

func (s *SkillsHandler) State() domain.State {
	return s.stateProvider()
}

func NewSkillsHandler(stateProvider domain.StateProvider, skills *loader.SkillLoader) *SkillsHandler {
	return &SkillsHandler{
		stateProvider: stateProvider,
		skills:        skills,
	}
}

var _ Handler = &SkillsHandler{}
