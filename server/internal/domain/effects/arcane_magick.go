package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ArcaneMagick struct {
  Name string
  Description string
}

func (e *ArcaneMagick) Apply(state domain.State) domain.State {
  // - You are solely a practitioner of Arcane Magick. As a consequence, you may never adopt a Profession which has Divine Magick as a Special Trait. In addition, you immediately learn three Generalist Magick spells when you enter this Profession. If you gained Arcane Magick from a previous Profession, you learn an additional Generalist Magick spell instead.
  log.Println("applying Arcane Magick")
  return state
}
