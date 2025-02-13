package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type SkillRankSpec struct {
	Job   string `yaml:"job"`
	Skill string `yaml:"skill"`
}

type SkillRankSpecs []SkillRankSpec

type InventorySpec struct {
	MainHand string   `yaml:"mainHand"`
	OffHand  string   `yaml:"offHand"`
	Armor    string   `yaml:"armor"`
	Pack     []string `yaml:"pack"`
	Cash     int      `yaml:"cash"`
}

type NPCSpec struct {
	Name                string                     `yaml:"name"`
	Age                 int                        `yaml:"age"`
	Alignment           domain.AlignmentSpec       `yaml:"alignment"`
	Archetype           string                     `yaml:"archetype"`
	Background          string                     `yaml:"background"`
	BackgroundTrait     string                     `yaml:"backgroundTrait"`
	BirthSeason         domain.Season              `yaml:"birthSeason"`
	Condition           domain.Condition           `yaml:"condition"`
	ConsumedAdvances    domain.ConsumedAdvances    `yaml:"consumedAdvances"`
	DistinguishingMarks domain.DistinguishingMarks `yaml:"distinguishingMarks"`
	Drawback            string                     `yaml:"drawback"`
	FatePoints          int                        `yaml:"fatePoints"`
	Inventory           InventorySpec              `yaml:"inventory"`
	Job                 string                     `yaml:"job"`
	Peril               string                     `yaml:"peril"`
	Poorness            domain.Poorness            `yaml:"poorness"`
	Room                string                     `yaml:"room"`
	Team                string                     `yaml:"team"`
	Tattoo              domain.Tattoo              `yaml:"tattoo"`
	SkillRanks          SkillRankSpecs             `yaml:"skillRanks"`
	Stats               domain.Stats               `yaml:"stats"`
	Talents             []string                   `yaml:"talents"`
	Upbringing          string                     `yaml:"upbringing"`
}

type NPCSpecs map[string]*NPCSpec

func (n *NPCSpec) ToProperties(loaders *Loaders) (map[string]domain.Property, error) {
	props := make(map[string]domain.Property)
	props[domain.AgeProperty] = &domain.BaseProperty{
		Val: n.Age,
	}
	alignment, err := loaders.AlignmentLoader.GetAlignment(n.Alignment.Order)
	if err != nil {
		return nil, err
	}
	alignment.AddOrderRank(n.Alignment.OrderRank)
	alignment.AddChaosRank(n.Alignment.ChaosRank)
	props[domain.AlignmentProperty] = alignment
	archetype, err := loaders.ArchetypeLoader.GetArchetype(n.Archetype)
	if err != nil {
		return nil, err
	}
	props[domain.ArchetypeProperty] = archetype
	background, err := loaders.BackgroundLoader.GetBackground(n.Background)
	if err != nil {
		return nil, err
	}
	props[domain.BackgroundProperty] = background
	backgroundTrait, err := loaders.TraitLoader.GetTrait(n.BackgroundTrait)
	if err != nil {
		return nil, err
	}
	props[domain.BackgroundTraitProperty] = backgroundTrait
	props[domain.BirthSeasonProperty] = n.BirthSeason
	props[domain.ConditionProperty] = n.Condition
	props[domain.ConsumedAdvancesProperty] = n.ConsumedAdvances
	props[domain.DistinguishingMarksProperty] = n.DistinguishingMarks
	drawback, err := loaders.AppearanceLoader.GetDrawback(n.Drawback)
	if err != nil {
		return nil, err
	}
	props[domain.DrawbackProperty] = drawback
	props[domain.FatePointsProperty] = &domain.BaseProperty{
		Val: n.FatePoints,
	}
	// TODO: inventory
	job, err := loaders.JobLoader.GetJob(n.Job)
	if err != nil {
		return nil, err
	}
	props[domain.JobProperty] = job
	props[domain.PoornessProperty] = n.Poorness
	props[domain.PerilProperty] = &domain.Peril{
		Threshold: 0, // TODO: 3 + StatBonuses.Grit
		Condition: domain.PerilConditionFromString(n.Peril),
	}
	room := loaders.RoomLoader.GetRoom(n.Room)
	props[domain.RoomProperty] = room
	team, err := loaders.TeamLoader.GetTeam(n.Team)
	if err != nil {
		return nil, err
	}
	// TODO: skill ranks
	props[domain.StatsProperty] = &n.Stats
	props[domain.TeamProperty] = team
	props[domain.TattooProperty] = &n.Tattoo
	talents := make(domain.Talents, 0)
	for _, talent := range n.Talents {
		t, err := loaders.TalentLoader.GetTalent(talent)
		if err != nil {
			return nil, err
		}
		talents = append(talents, t)
	}
	props[domain.TalentsProperty] = talents
	upbringing, err := loaders.UpbringingLoader.GetUpbringing(n.Upbringing)
	if err != nil {
		return nil, err
	}
	props[domain.UpbringingProperty] = upbringing
	return props, nil
}

type NPCLoader struct {
	config          *config.Config
	alignmentLoader *AlignmentLoader
	npcs            NPCSpecs
}

func NewNPCLoader(cfg *config.Config, alignmentLoader *AlignmentLoader) *NPCLoader {
	return &NPCLoader{
		config:          cfg,
		alignmentLoader: alignmentLoader,
		npcs:            make(NPCSpecs),
	}
}

func (l *NPCLoader) LoadNPCs() (NPCSpecs, error) {
	if l.npcs != nil && len(l.npcs) > 0 {
		return l.npcs, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/npcs")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if strings.HasSuffix(item.Name(), "tmpl.yaml") {
			continue
		}
		spec := &NPCSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/npcs/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}
		l.npcs[spec.Name] = spec
	}
	return l.npcs, nil
}

func (l *NPCLoader) GetNPC(name string) (*NPCSpec, error) {
	if l.npcs == nil || len(l.npcs) == 0 {
		_, err := l.LoadNPCs()
		if err != nil {
			return nil, err
		}
	}
	if npc, ok := l.npcs[name]; ok {
		return npc, nil
	}
	return nil, nil
}
