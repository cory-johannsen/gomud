package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type ManifestDestiny struct {
  name string
  description string
}

func NewManifestDestiny() *ManifestDestiny {
  return &ManifestDestiny{
    name: "Manifest Destiny",
    description: "Effect1",
  }
}

func (e *ManifestDestiny) Name() string {
  return e.name
}

func (e *ManifestDestiny) Description() string {
  return e.description
}

func (e *ManifestDestiny) Applier() domain.Applier {
  return e.Apply
}

func (e *ManifestDestiny) Apply(state domain.State) domain.State {
  // - Effect1
  log.Println("applying Manifest Destiny")
  return state
}

var _ domain.Effect = &ManifestDestiny{}
