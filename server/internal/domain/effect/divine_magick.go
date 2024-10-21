package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type DivineMagick struct {
  name string
  description string
}

func NewDivineMagick() *DivineMagick {
  return &DivineMagick{
    name: "Divine Magick",
    description: "You are solely a practitioner of Divine Magick. As a consequence, you may never adopt a Profession which has Arcane Magick as a Special Trait. In addition, you immediately learn three Generalist Magick spells when you enter this Profession. If you gained Divine Magick from a previous Profession, you learn an additional Generalist Magick spell instead.",
  }
}

func (e *DivineMagick) Name() string {
  return e.name
}

func (e *DivineMagick) Description() string {
  return e.description
}

func (e *DivineMagick) Applier() domain.Applier {
  return e.Apply
}

func (e *DivineMagick) Apply(state domain.State) domain.State {
  // - You are solely a practitioner of Divine Magick. As a consequence, you may never adopt a Profession which has Arcane Magick as a Special Trait. In addition, you immediately learn three Generalist Magick spells when you enter this Profession. If you gained Divine Magick from a previous Profession, you learn an additional Generalist Magick spell instead.
  log.Println("applying Divine Magick")
  return state
}

var _ domain.Effect = &DivineMagick{}
