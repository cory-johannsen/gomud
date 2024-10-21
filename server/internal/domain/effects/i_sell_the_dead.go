package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ISelltheDead struct {
  Name string
  Description string
}

func (e *ISelltheDead) Apply(state domain.State) domain.State {
  // - When facing Supernatural creatures, you always succeed at Resolve Tests to save yourself from Stress and Fear provoked by them. In addition, you are immune to specific Diseases such as Red Death and Tomb Rot.
  log.Println("applying I Sell the Dead")
  return state
}
