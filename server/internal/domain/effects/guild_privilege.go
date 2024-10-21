package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type GuildPrivilege struct {
  Name string
  Description string
}

func (e *GuildPrivilege) Apply(state domain.State) domain.State {
  // - Whenever you purchase a Focus in the Tradecraft Skill, you instead gain three Focuses. This means you may exceed the normal limits for Focuses set by your [IB], but for Tradecraft only.
  log.Println("applying Guild Privilege")
  return state
}
