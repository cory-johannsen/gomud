package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ScrivenersSpeed struct {
  Name string
  Description string
}

func (e *ScrivenersSpeed) Apply(state domain.State) domain.State {
  // - You reduce the time required to use Education Tests by half (examples include research, speed-reading, writing, etc.). In addition, you automatically succeed at all Skill Tests to decipher cryptographs in languages they understand.
  log.Println("applying Scrivener's Speed")
  return state
}
