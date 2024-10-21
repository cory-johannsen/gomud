package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LilyLivered struct {
  name string
  description string
}

func NewLilyLivered() *LilyLivered {
  return &LilyLivered{
    name: "Lily-Livered",
    description: "Whenever you fail to Resist Stress, Fear or Terror, you temporarily reduce your Initiative and Movement by -3 (to a minimum of 1). This lasts until you get a good night's rest.",
  }
}

func (e *LilyLivered) Name() string {
  return e.name
}

func (e *LilyLivered) Description() string {
  return e.description
}

func (e *LilyLivered) Applier() domain.Applier {
  return e.Apply
}

func (e *LilyLivered) Apply(state domain.State) domain.State {
  // - Until fully Recuperated, you cannot see as you’re blinded. You must undergo a successful surgery or suffer the consequences. Once a Vitreous Hemorrhage has undergone a failed surgery, you gain the Black Cataract Drawback. If you already have this Drawback, you permanently lose 9% Perception.
  log.Println("applying Lily-Livered")
  return state
}

var _ domain.Effect = &LilyLivered{}
