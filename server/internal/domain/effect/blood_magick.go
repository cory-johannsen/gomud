package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type BloodMagick struct {
  name string
  description string
}

func NewBloodMagick() *BloodMagick {
  return &BloodMagick{
    name: "Blood Magick",
    description: "After you have made the appropriate sacrifice of an innocent creature (a beloved animal for Petty Magick, a farm animal for Lesser Magick or a living person for Greater Magick), you can cause any one foe to automatically fail to Resist one cast Magick spell within the next 24 hours. You can only make a sacrifice like this once per day and will likely suffer Corruption for such a heinous act.",
  }
}

func (e *BloodMagick) Name() string {
  return e.name
}

func (e *BloodMagick) Description() string {
  return e.description
}

func (e *BloodMagick) Applier() domain.Applier {
  return e.Apply
}

func (e *BloodMagick) Apply(state domain.State) domain.State {
  // - After you have made the appropriate sacrifice of an innocent creature (a beloved animal for Petty Magick, a farm animal for Lesser Magick or a living person for Greater Magick), you can cause any one foe to automatically fail to Resist one cast Magick spell within the next 24 hours. You can only make a sacrifice like this once per day and will likely suffer Corruption for such a heinous act.
  log.Println("applying Blood Magick")
  return state
}

var _ domain.Effect = &BloodMagick{}
