package domain

type SkillRankSpec struct {
	Job   string `yaml:"job"`
	Skill string `yaml:"skill"`
}

type SkillRankSpecs []SkillRankSpec

type NPCInventorySpec struct {
	MainHand string   `yaml:"mainHand"`
	OffHand  string   `yaml:"offHand"`
	Armor    string   `yaml:"armor"`
	Pack     []string `yaml:"pack"`
	Cash     int      `yaml:"cash"`
}

type NPCSpec struct {
	Name                string              `yaml:"name"`
	Age                 int                 `yaml:"age"`
	Alignment           AlignmentSpec       `yaml:"alignment"`
	Archetype           string              `yaml:"archetype"`
	Background          string              `yaml:"background"`
	BackgroundTrait     string              `yaml:"backgroundTrait"`
	BirthSeason         Season              `yaml:"birthSeason"`
	Condition           Condition           `yaml:"condition"`
	ConsumedAdvances    ConsumedAdvances    `yaml:"consumedAdvances"`
	DistinguishingMarks DistinguishingMarks `yaml:"distinguishingMarks"`
	Drawback            string              `yaml:"drawback"`
	FatePoints          int                 `yaml:"fatePoints"`
	Inventory           NPCInventorySpec    `yaml:"inventory"`
	Job                 string              `yaml:"job"`
	Peril               string              `yaml:"peril"`
	Poorness            Poorness            `yaml:"poorness"`
	Room                string              `yaml:"room"`
	Team                string              `yaml:"team"`
	Tattoo              Tattoo              `yaml:"tattoo"`
	SkillRanks          SkillRankSpecs      `yaml:"skillRanks"`
	Stats               Stats               `yaml:"stats"`
	Talents             []string            `yaml:"talents"`
	Upbringing          string              `yaml:"upbringing"`
}

type NPCSpecs map[string]*NPCSpec
