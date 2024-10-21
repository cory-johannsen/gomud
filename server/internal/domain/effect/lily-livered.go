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
  // - When gossiping or spinning a tale, you gain a +20 Base Chance to Rumor Tests.
  log.Println("applying Lily-Livered")
  return state
}

var _ domain.Effect = &LilyLivered{}
