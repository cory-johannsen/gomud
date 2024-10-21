package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type StentorianVoice struct {
  Name string
  Description string
}

func (e *StentorianVoice) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Stentorian Voice")
  return state
}
