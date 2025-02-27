package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

func NPCSpecToProperties(n *domain.NPCSpec, loaders *Loaders) (map[string]domain.Property, error) {
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
	inventory := domain.NewInventory()
	if n.Inventory.MainHand != "" {
		item, err := loaders.EquipmentLoader.ItemFromName(n.Inventory.MainHand)
		if err != nil {
			return nil, err
		}
		mainHand := item.(*domain.Weapon)
		err = inventory.EquipMainHand(mainHand)
		if err != nil {
			return nil, err
		}
	}
	if n.Inventory.OffHand != "" {
		item, err := loaders.EquipmentLoader.ItemFromName(n.Inventory.OffHand)
		if err != nil {
			return nil, err
		}
		offHand := item.(*domain.Weapon)
		err = inventory.EquipOffHand(offHand)
		if err != nil {
			return nil, err
		}
	}
	if n.Inventory.Armor != "" {
		item, err := loaders.EquipmentLoader.ItemFromName(n.Inventory.Armor)
		if err != nil {
			return nil, err
		}
		armor := item.(*domain.Armor)
		err = inventory.EquipArmor(armor)
		if err != nil {
			return nil, err
		}
	}
	for _, name := range n.Inventory.Pack {
		if name != "" {
			item, err := loaders.EquipmentLoader.ItemFromName(n.Inventory.Armor)
			if err != nil {
				return nil, err
			}
			err = inventory.Pack().AddItem(item)
			if err != nil {
				return nil, err
			}
		}
	}
	inventory.AddCash(n.Inventory.Cash)
	props[domain.InventoryProperty] = inventory
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
	npcs            domain.NPCSpecs
}

func NewNPCLoader(cfg *config.Config, alignmentLoader *AlignmentLoader) *NPCLoader {
	return &NPCLoader{
		config:          cfg,
		alignmentLoader: alignmentLoader,
		npcs:            make(domain.NPCSpecs),
	}
}

func (l *NPCLoader) LoadNPCs() (domain.NPCSpecs, error) {
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
		spec := &domain.NPCSpec{}
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

func (l *NPCLoader) GetNPC(name string) (*domain.NPCSpec, error) {
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
