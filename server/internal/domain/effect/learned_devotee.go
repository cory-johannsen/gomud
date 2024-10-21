package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type LearnedDevotee struct {
  name string
  description string
}

func NewLearnedDevotee() *LearnedDevotee {
  return &LearnedDevotee{
    name: "Learned Devotee",
    description: "You can read, write and speak the language of any creature that is classified as a Humanoid (including player Ancestries). In addition, whenever you suffer Corruption, decrease the number you gain by 1 (to a minimum of one). This means that if you suffer 3 Corruption, you gain 2 instead.",
  }
}

func (e *LearnedDevotee) Name() string {
  return e.name
}

func (e *LearnedDevotee) Description() string {
  return e.description
}

func (e *LearnedDevotee) Applier() domain.Applier {
  return e.Apply
}

func (e *LearnedDevotee) Apply(state domain.State) domain.State {
  // - You can read, write and speak the language of any creature that is classified as a Humanoid (including player Ancestries). In addition, whenever you suffer Corruption, decrease the number you gain by 1 (to a minimum of one). This means that if you suffer 3 Corruption, you gain 2 instead.
  log.Println("applying Learned Devotee")
  return state
}

var _ domain.Effect = &LearnedDevotee{}
