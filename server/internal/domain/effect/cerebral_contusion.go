package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type CerebralContusion struct {
  name string
  description string
}

func NewCerebralContusion() *CerebralContusion {
  return &CerebralContusion{
    name: "Cerebral Contusion",
    description: "Until fully Recuperated, you remain unconscious. You must undergo a successful surgery or suffer the consequences. Once a Cerebral Contusion has undergone a failed surgery, you gain the Dunderhead Drawback. If you already have this Drawback, you permanently lose 9% Intelligence.",
  }
}

func (e *CerebralContusion) Name() string {
  return e.name
}

func (e *CerebralContusion) Description() string {
  return e.description
}

func (e *CerebralContusion) Applier() domain.Applier {
  return e.Apply
}

func (e *CerebralContusion) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you remain unconscious. You must undergo a successful surgery or suffer the consequences. Once a Cerebral Contusion has undergone a failed surgery, you gain the Dunderhead Drawback. If you already have this Drawback, you permanently lose 9% Intelligence.
  log.Println("applying Cerebral Contusion")
  return state
}

var _ domain.Effect = &CerebralContusion{}
