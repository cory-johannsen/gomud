package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DivineMagick struct {
  Name string
  Description string
}

func (e *DivineMagick) Apply(state domain.State) domain.State {
  // - You are solely a practitioner of Divine Magick. As a consequence, you may never adopt a Profession which has Arcane Magick as a Special Trait. In addition, you immediately learn three Generalist Magick spells when you enter this Profession. If you gained Divine Magick from a previous Profession, you learn an additional Generalist Magick spell instead.
  log.Println("applying Divine Magick")
  return state
}
