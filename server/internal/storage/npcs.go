package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cory-johannsen/gomud/internal/domain"
	"github.com/cory-johannsen/gomud/internal/loader"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

type NPCs struct {
	database  *Database
	loaders   *loader.Loaders
	equipment *Equipment
	npcs      map[string]*domain.Character
}

func NewNPCs(db *Database, loaders *loader.Loaders, equipment *Equipment) *NPCs {
	return &NPCs{
		database:  db,
		loaders:   loaders,
		equipment: equipment,
		npcs:      make(map[string]*domain.Character),
	}
}

func (n *NPCs) CreateNPC(ctx context.Context, name string, data map[string]domain.Property) (*domain.Character, error) {
	if _, ok := n.npcs[name]; ok {
		return nil, errors.New(fmt.Sprintf("npc %s already exists", name))
	}
	specData := n.PropertiesToData(data)
	encoded, err := json.Marshal(specData)
	if err != nil {
		log.Errorf("failed to marshal player data: %s", err)
		return nil, err
	}
	var id int
	err = n.database.Conn.QueryRow(ctx, "INSERT INTO npcs (name, data) VALUES ($1, $2) RETURNING id", name, encoded).Scan(&id)
	if err != nil {
		log.Errorf("failed to insert player: %s", err)
		return nil, err
	}
	npc := domain.NewCharacter(&id, name, data)
	n.npcs[name] = npc
	return npc, nil
}

func (n *NPCs) FetchNPCById(ctx context.Context, id int) (*domain.Character, error) {
	for _, npc := range n.npcs {
		if npc.Id != nil && *npc.Id == id {
			return npc, nil
		}
	}
	var name, data string
	err := n.database.Conn.QueryRow(ctx, "SELECT name, data FROM npcs WHERE id = $1", id).Scan(&name, &data)
	if err != nil {
		log.Errorf("failed to fetch npc: %s", err)
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Errorf("failed to unmarshal npc data: %s", err)
		return nil, err
	}
	spec, err := n.loaders.NPCLoader.GetNPC(name)
	if err != nil {
		log.Errorf("failed to load npc: %s", err)
		return nil, err
	}
	npc := n.NPCFromSpec(ctx, spec, id, specProps)

	// todo: load equipment

	return npc, nil
}

func (n *NPCs) FetchNPCByName(ctx context.Context, name string) (*domain.Character, error) {
	if npc, ok := n.npcs[name]; ok {
		return npc, nil
	}
	var id int
	var data string
	err := n.database.Conn.QueryRow(ctx, "SELECT id, data FROM npcs WHERE name = $1", name).Scan(&id, &data)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		log.Errorf("failed to fetch npc: %s", err)
		return nil, err
	}
	specProps := make(map[string]interface{})
	err = json.Unmarshal([]byte(data), &specProps)
	if err != nil {
		log.Errorf("failed to unmarshal npc data: %s", err)
		return nil, err
	}
	props := n.DataToProperties(ctx, specProps)
	npc := domain.NewCharacter(&id, name, props)
	// Peril threshold is calculated from Grit Bonus
	npc.Peril().Threshold = npc.StatBonuses().Grit + 3
	n.npcs[name] = npc
	return npc, nil
}

func (n *NPCs) Exists(ctx context.Context, name string) (bool, error) {
	var count int
	row := n.database.Conn.QueryRow(ctx, "SELECT count(*) FROM npcs WHERE name = $1", name)
	err := row.Scan(&count)
	if err != nil {
		log.Errorf("failed to check if player exists: %s", err)
		return false, err
	}
	return count > 0, nil
}

func (n *NPCs) NPCFromSpec(ctx context.Context, spec *loader.NPCSpec, id int, data map[string]interface{}) *domain.Character {
	props := n.DataToProperties(ctx, data)

	// todo: prop loading/overloading?

	npc := domain.NewCharacter(&id, spec.Name, props)
	return npc
}

func (n *NPCs) PropertiesToData(props map[string]domain.Property) map[string]interface{} {
	data := make(map[string]interface{})
	for k, v := range props {
		switch k {
		case domain.AgeProperty:
			data[k] = v.(*domain.BaseProperty).Val
		case domain.AlignmentProperty:
			data[k] = domain.SpecFromAlignment(v.(*domain.Alignment))
		case domain.ArchetypeProperty:
			data[k] = v.(*domain.Archetype).Name
		case domain.BackgroundProperty:
			data[k] = v.(*domain.Background).Name
		case domain.BackgroundTraitProperty:
			data[k] = v.(*domain.Trait).Name
		case domain.ConditionProperty:
			data[k] = v.(domain.Condition)
		case domain.ConsumedAdvancesProperty:
			data[k] = v
		case domain.DistinguishingMarksProperty:
			data[k] = v
		case domain.DrawbackProperty:
			data[k] = v.(*domain.Drawback).Name
		case domain.DisordersProperty:
			disorders := make([]string, 0)
			for _, disorder := range v.(domain.Disorders) {
				disorders = append(disorders, disorder.Name)
			}
			data[k] = disorders
		case domain.InventoryProperty:
			inv := v.(*domain.Inventory)
			// ensure the inventory items have been persisted
			if inv.MainHand() != nil && inv.MainHand().Id() == 0 {
				item, err := n.equipment.CreateItem(context.Background(), inv.MainHand())
				if err != nil {
					log.Printf("failed to create main hand item: %s", err)
				}
				err = inv.EquipMainHand(item.(*domain.Weapon))
				if err != nil {
					log.Printf("failed to update main hand item: %s", err)
				}
			}
			if inv.OffHand() != nil && inv.OffHand().Id() == 0 {
				item, err := n.equipment.CreateItem(context.Background(), inv.OffHand())
				if err != nil {
					log.Printf("failed to create off hand item: %s", err)
				}
				err = inv.EquipOffHand(item.(*domain.Weapon))
				if err != nil {
					log.Printf("failed to update off hand item: %s", err)
				}
			}
			if inv.Armor() != nil && inv.Armor().Id() == 0 {
				item, err := n.equipment.CreateItem(context.Background(), inv.Armor())
				if err != nil {
					log.Printf("failed to create armor item: %s", err)
				}
				err = inv.EquipArmor(item.(*domain.Armor))
				if err != nil {
					log.Printf("failed to update armor item: %s", err)
				}
			}
			for _, item := range inv.Pack().Items() {
				if item.Id() == 0 {
					i, err := n.equipment.CreateItem(context.Background(), item)
					if err != nil {
						log.Printf("failed to create pack item: %s", err)
					}
					err = inv.Pack().RemoveItem(item)
					if err != nil {
						log.Printf("failed to remove item from pack: %s", err)
					}
					err = inv.Pack().AddItem(i)
					if err != nil {
						log.Printf("failed to add item to pack: %s", err)
					}
				}
			}
			data[k] = domain.SpecFromInventory(inv)
		case domain.JobProperty:
			data[k] = v.(*domain.Job).Name
		case domain.PoornessProperty:
			data[k] = v.(domain.Poorness)
		case domain.RoomProperty:
			data[k] = v.(*domain.Room).Name
		case domain.SkillRanksProperty:
			skillRanks := make(map[string][]string)
			for _, skill := range v.(domain.SkillRanks) {
				jobName := skill.Job.Name
				if _, ok := skillRanks[jobName]; !ok {
					skillRanks[jobName] = make([]string, 0)
				}
				skillRanks[jobName] = append(skillRanks[jobName], skill.Skill.Name)
			}
			data[k] = skillRanks
		case domain.TeamProperty:
			data[k] = domain.SpecFromTeam(v.(*domain.Team))
		case domain.TalentsProperty:
			talents := make([]string, 0)
			for _, talent := range v.(domain.Talents) {
				talents = append(talents, talent.Name)
			}
			data[k] = talents
		case domain.UpbringingProperty:
			data[k] = v.(*domain.Upbringing).Name
		case domain.TattooProperty:
			fallthrough
		case domain.BirthSeasonProperty:
			fallthrough
		case domain.InjuriesProperty:
			fallthrough
		case domain.PerilProperty:
			fallthrough
		case domain.StatsProperty:
			data[k] = v
		case domain.ReputationPointsProperty:
			fallthrough
		case domain.FatePointsProperty:
			fallthrough
		default:
			data[k] = v.(*domain.BaseProperty).Val
		}
	}
	return data
}

func (n *NPCs) DataToProperties(ctx context.Context, data map[string]interface{}) map[string]domain.Property {
	props := make(map[string]domain.Property)
	for k, v := range data {
		switch k {
		case domain.AgeProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
			continue
		case domain.AlignmentProperty:
			alignmentMap := v.(map[string]interface{})
			alignment, err := n.loaders.AlignmentLoader.GetAlignment(alignmentMap["order"].(string))
			if err != nil {
				log.Printf("failed to load alignment %s: %s", alignmentMap["order"].(string), err)
				continue
			}
			if alignment == nil {
				log.Printf("alignment %s not found", alignmentMap["order"].(string))
				continue
			}
			props[k] = alignment
		case domain.ArchetypeProperty:
			archetype, err := n.loaders.ArchetypeLoader.GetArchetype(v.(string))
			if err != nil {
				log.Printf("failed to load archetype %s: %s", v.(string), err)
				continue
			}
			if archetype == nil {
				log.Printf("archetype %s not found", v.(string))
				continue
			}
			props[k] = archetype
		case domain.BackgroundProperty:
			name := v.(string)
			background, err := n.loaders.BackgroundLoader.GetBackground(name)
			if err != nil {
				log.Printf("failed to load background %s: %s", v.(string), err)
				continue
			}
			if background == nil {
				log.Printf("background %s not found", v.(string))
				continue
			}
			props[k] = background
		case domain.BackgroundTraitProperty:
			name := v.(string)
			trait, err := n.loaders.TraitLoader.GetTrait(name)
			if err != nil {
				log.Printf("failed to load trait %s: %s", v.(string), err)
				continue
			}
			if trait == nil {
				log.Printf("trait %s not found", v.(string))
				continue
			}
			props[k] = trait
		case domain.BirthSeasonProperty:
			props[k] = domain.Season(v.(string))
		case domain.ConditionProperty:
			props[k] = domain.Condition(v.(string))
		case domain.ConsumedAdvancesProperty:
			consumedAdvances := make(domain.ConsumedAdvances)
			for _, advances := range v.(map[string]interface{}) {
				for _, advance := range advances.([]interface{}) {
					job := advance.(map[string]interface{})["Job"].(string)
					stat := advance.(map[string]interface{})["Stat"].(string)
					amount := advance.(map[string]interface{})["Amount"].(float64)
					if _, ok := consumedAdvances[job]; !ok {
						consumedAdvances[job] = make([]*domain.ConsumedAdvance, 0)
					}

					var consumedAdvance *domain.ConsumedAdvance
					for _, ca := range consumedAdvances[job] {
						if ca.Stat == stat {
							consumedAdvance = ca
							break
						}
					}
					if consumedAdvance == nil {
						consumedAdvances[job] = append(consumedAdvances[job], &domain.ConsumedAdvance{
							Job:    job,
							Stat:   stat,
							Amount: int(amount),
						})
					} else {
						consumedAdvance.Amount += int(amount)
					}
				}
			}
			props[k] = consumedAdvances
		case domain.DisordersProperty:
			disorders := make(domain.Disorders, 0)
			for _, name := range v.([]interface{}) {
				disorder, err := n.loaders.DisorderLoader.GetDisorder(name.(string))
				if err != nil {
					log.Printf("failed to load disorder %s: %s", name, err)
					continue
				}
				if disorder == nil {
					log.Printf("disorder %s not found", name)
					continue
				}
				disorders = append(disorders, disorder)
			}
			props[k] = disorders
		case domain.DistinguishingMarksProperty:
			marks := make(domain.DistinguishingMarks, 0)
			for _, mark := range v.([]interface{}) {
				marks = append(marks, domain.DistinguishingMark(mark.(string)))
			}
			props[k] = marks
		case domain.DrawbackProperty:
			drawback, err := n.loaders.AppearanceLoader.GetDrawback(v.(string))
			if err != nil {
				log.Printf("failed to load drawback %s: %s", v.(string), err)
				continue
			}
			if drawback == nil {
				log.Printf("drawback %s not found", v.(string))
				continue
			}
			props[k] = drawback
		case domain.ExperienceProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
		case domain.FatePointsProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
		case domain.InjuriesProperty:
			injuries := make(domain.Injuries, 0)
			for _, name := range v.([]interface{}) {
				injury, err := n.loaders.InjuryLoader.GetInjury(name.(string))
				if err != nil {
					log.Printf("failed to load injury %s: %s", name, err)
					continue
				}
				if injury == nil {
					log.Printf("injury %s not found", name)
					continue
				}
				injuries = append(injuries, injury)
			}
			props[k] = injuries
		case domain.InventoryProperty:
			m := v.(map[string]interface{})
			pack := make([]int, 0)
			var mainHand = 0
			var offHand = 0
			var armor = 0
			var cash = 0
			if _, ok := m["MainHand"]; ok {
				id := int(m["MainHand"].(float64))
				if id != 0 {
					mainHand = id
				}
			}
			if _, ok := m["OffHand"]; ok {
				id := int(m["OffHand"].(float64))
				if id != 0 {
					offHand = id
				}
			}
			if _, ok := m["Armor"]; ok {
				id := int(m["Armor"].(float64))
				if id != 0 {
					armor = id
				}
			}
			if _, ok := m["Pack"]; ok {
				ids := m["Pack"].([]interface{})
				for _, id := range ids {
					if id != 0 {
						pack = append(pack, int(id.(float64)))
					}
				}
			}
			if _, ok := m["Cash"]; ok {
				cash = int(m["Cash"].(float64))
			}
			spec := &domain.InventorySpec{
				MainHand: mainHand,
				OffHand:  offHand,
				Armor:    armor,
				Pack:     pack,
				Cash:     cash,
			}
			inventory, err := n.loaders.InventoryLoader.InventoryFromSpec(ctx, spec, n.equipment.FetchItemByID)
			if err != nil {
				log.Printf("failed to load inventory: %s", err)
				continue
			}
			if inventory == nil {
				inventory = domain.NewInventory()
			}
			props[k] = inventory
		case domain.JobProperty:
			job, err := n.loaders.JobLoader.GetJob(v.(string))
			if err != nil {
				log.Printf("failed to load job %s: %s", v.(string), err)
				continue
			}
			if job == nil {
				log.Printf("job %s not found", v.(string))
				continue
			}
			props[k] = job
		case domain.PerilProperty:
			peril := v.(map[string]interface{})
			threshold := peril["Threshold"].(float64)
			perilCondition := peril["Condition"].(float64)
			props[k] = &domain.Peril{
				Threshold: int(threshold),
				Condition: domain.PerilCondition(perilCondition),
			}
		case domain.PoornessProperty:
			props[k] = domain.Poorness(v.(string))
		case domain.ReputationPointsProperty:
			props[k] = &domain.BaseProperty{Val: int(v.(float64))}
		case domain.RoomProperty:
			room := n.loaders.RoomLoader.GetRoom(v.(string))
			props[k] = room
		case domain.SkillRanksProperty:
			skillRanks := make(domain.SkillRanks, 0)
			for jobName, skills := range v.(map[string]interface{}) {
				job, err := n.loaders.JobLoader.GetJob(jobName)
				if err != nil {
					log.Printf("failed to load job %s: %s", jobName, err)
					continue
				}
				if job == nil {
					log.Printf("job %s not found", jobName)
					continue
				}
				for _, skillName := range skills.([]interface{}) {
					skill, err := n.loaders.SkillLoader.GetSkill(skillName.(string))
					if err != nil {
						log.Printf("failed to load skill %s: %s", skillName.(string), err)
						continue
					}
					if skill == nil {
						log.Printf("skill %s not found", skillName.(string))
						continue
					}
					skillRanks = append(skillRanks, &domain.SkillRank{
						Job:   job,
						Skill: skill,
					})
				}
			}
			props[k] = skillRanks
		case domain.TeamProperty:
			teamName := v.(map[string]interface{})["Name"].(string)
			team, err := n.loaders.TeamLoader.GetTeam(teamName)
			if err != nil {
				log.Printf("failed to load team %s: %s", v.(string), err)
				continue
			}
			if team == nil {
				log.Printf("team %s not found", v.(string))
				continue
			}
			props[k] = team
		case domain.TattooProperty:
			description := v.(map[string]interface{})["Description"].(string)
			location := v.(map[string]interface{})["Location"].(string)
			tat := &domain.Tattoo{
				Description: description,
				Location:    domain.TattooLocation(location),
				Season:      domain.Season(v.(map[string]interface{})["Season"].(string)),
			}
			props[k] = tat
		case domain.TalentsProperty:
			talents := make(domain.Talents, 0)
			for _, talentName := range v.([]interface{}) {
				talent, err := n.loaders.TalentLoader.GetTalent(talentName.(string))
				if err != nil {
					log.Printf("failed to load talent %s: %s", talentName.(string), err)
					continue
				}
				if talent == nil {
					log.Printf("talent %s not found", talentName.(string))
					continue
				}
				talents = append(talents, talent)
			}
			props[k] = talents
		case domain.UpbringingProperty:
			upbringing, err := n.loaders.UpbringingLoader.GetUpbringing(v.(string))
			if err != nil {
				log.Printf("failed to load upbringing %s: %s", v.(string), err)
				continue
			}
			if upbringing == nil {
				log.Printf("upbringing %s not found", v.(string))
				continue
			}
			props[k] = upbringing
		case domain.StatsProperty:
			brutality := int(v.(map[string]interface{})["brutality"].(float64))
			muscle := int(v.(map[string]interface{})["muscle"].(float64))
			quickness := int(v.(map[string]interface{})["quickness"].(float64))
			savvy := int(v.(map[string]interface{})["savvy"].(float64))
			reasoning := int(v.(map[string]interface{})["reasoning"].(float64))
			grit := int(v.(map[string]interface{})["grit"].(float64))
			flair := int(v.(map[string]interface{})["flair"].(float64))
			stats := &domain.Stats{
				Brutality: brutality,
				Muscle:    muscle,
				Quickness: quickness,
				Savvy:     savvy,
				Reasoning: reasoning,
				Grit:      grit,
				Flair:     flair,
			}
			props[k] = stats
		default:
			log.Printf("unknown property %s: %v", k, v)
		}
	}
	return props
}
