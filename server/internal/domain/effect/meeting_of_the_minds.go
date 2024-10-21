package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type MeetingoftheMinds struct {
  name string
  description string
}

func NewMeetingoftheMinds() *MeetingoftheMinds {
  return &MeetingoftheMinds{
    name: "Meeting of the Minds",
    description: "When attempting to bring compromise between two extreme positions, you gain a +10 Base Chance to Skill Tests.",
  }
}

func (e *MeetingoftheMinds) Name() string {
  return e.name
}

func (e *MeetingoftheMinds) Description() string {
  return e.description
}

func (e *MeetingoftheMinds) Applier() domain.Applier {
  return e.Apply
}

func (e *MeetingoftheMinds) Apply(state domain.State) domain.State {
  // - When attempting to bring compromise between two extreme positions, you gain a +10 Base Chance to Skill Tests.
  log.Println("applying Meeting of the Minds")
  return state
}

var _ domain.Effect = &MeetingoftheMinds{}
