package domain

type GeneratorSpec struct {
	Name              string `yaml:"name"`
	Room              string `yaml:"room"`
	NPC               string `yaml:"npc"`
	Minimum           int    `yaml:"minimum"`
	Maximum           int    `yaml:"maximum"`
	SpawnDelaySeconds int    `yaml:"spawnDelaySeconds"`
}
