package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GuildPrivilege struct {
  name string
  description string
}

func NewGuildPrivilege() *GuildPrivilege {
  return &GuildPrivilege{
    name: "Guild Privilege",
    description: "Whenever you purchase a Focus in the Tradecraft Skill, you instead gain three Focuses. This means you may exceed the normal limits for Focuses set by your [IB], but for Tradecraft only.",
  }
}

func (e *GuildPrivilege) Name() string {
  return e.name
}

func (e *GuildPrivilege) Description() string {
  return e.description
}

func (e *GuildPrivilege) Applier() domain.Applier {
  return e.Apply
}

func (e *GuildPrivilege) Apply(state domain.State) domain.State {
  // - Whenever you purchase a Focus in the Tradecraft Skill, you instead gain three Focuses. This means you may exceed the normal limits for Focuses set by your [IB], but for Tradecraft only.
  log.Println("applying Guild Privilege")
  return state
}

var _ domain.Effect = &GuildPrivilege{}
