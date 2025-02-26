package loader

import log "github.com/sirupsen/logrus"

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	AlignmentLoader  *AlignmentLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	DisorderLoader   *DisorderLoader
	EffectLoader     *EffectLoader
	EquipmentLoader  *EquipmentLoader
	GeneratorLoader  *GeneratorLoader
	InjuryLoader     *InjuryLoader
	InventoryLoader  *InventoryLoader
	JobLoader        *JobLoader
	NPCLoader        *NPCLoader
	QualityLoader    *QualityLoader
	RoomLoader       *RoomLoader
	SkillLoader      *SkillLoader
	TalentLoader     *TalentLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
	UpbringingLoader *UpbringingLoader

	ActionLoader    *ActionLoader
	ConditionLoader *ConditionLoader
	MethodLoader    *MethodLoader
	TaskLoader      *TaskLoader
	TaskGraphLoader *TaskGraphLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, alignmentLoader *AlignmentLoader, archetypeLoader *ArchetypeLoader,
	backgroundLoader *BackgroundLoader, disorderLoader *DisorderLoader, effectLoader *EffectLoader, equipmentLoader *EquipmentLoader,
	generatorLoader *GeneratorLoader, injuryLoader *InjuryLoader,
	inventoryLoader *InventoryLoader, jobLoader *JobLoader, npcLoader *NPCLoader, qualityLoader *QualityLoader, roomLoader *RoomLoader, skillLoader *SkillLoader,
	talentLoader *TalentLoader, traitLoader *TraitLoader, teamLoader *TeamLoader, upbringingLoader *UpbringingLoader,
	actionLoader *ActionLoader, conditionLoader *ConditionLoader, methodLoader *MethodLoader, taskLoader *TaskLoader, taskGraphLoader *TaskGraphLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		AlignmentLoader:  alignmentLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		DisorderLoader:   disorderLoader,
		EffectLoader:     effectLoader,
		EquipmentLoader:  equipmentLoader,
		GeneratorLoader:  generatorLoader,
		InjuryLoader:     injuryLoader,
		InventoryLoader:  inventoryLoader,
		JobLoader:        jobLoader,
		NPCLoader:        npcLoader,
		QualityLoader:    qualityLoader,
		RoomLoader:       roomLoader,
		SkillLoader:      skillLoader,
		TalentLoader:     talentLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
		UpbringingLoader: upbringingLoader,

		ActionLoader:    actionLoader,
		ConditionLoader: conditionLoader,
		MethodLoader:    methodLoader,
		TaskLoader:      taskLoader,
		TaskGraphLoader: taskGraphLoader,
	}
}

func (l *Loaders) Preload() error {
	log.Println("Pre-loading assets")
	log.Info("loading alignments")
	_, err := l.AlignmentLoader.LoadAlignments()
	if err != nil {
		return err
	}
	log.Info("loading tattoo locations")
	_, err = l.AppearanceLoader.LoadTattooLocations()
	if err != nil {
		return err
	}
	log.Info("loading tattoos")
	_, err = l.AppearanceLoader.LoadTattoos()
	if err != nil {
		return err
	}
	log.Info("loading distinguishing marks")
	_, err = l.AppearanceLoader.LoadDistinguishingMarks()
	if err != nil {
		return err
	}
	log.Info("loading drawbacks")
	_, err = l.AppearanceLoader.LoadDrawbacks()
	if err != nil {
		return err
	}
	log.Info("loading archetypes")
	_, err = l.ArchetypeLoader.LoadArchetypes()
	if err != nil {
		return err
	}
	log.Info("loading backgrounds")
	_, err = l.BackgroundLoader.LoadBackgrounds()
	if err != nil {
		return err
	}
	log.Info("loading injuries")
	_, err = l.InjuryLoader.LoadInjuries()
	if err != nil {
		return err
	}
	log.Info("loading skills")
	_, err = l.SkillLoader.LoadSkills()
	if err != nil {
		return err
	}
	log.Info("loading talents")
	_, err = l.TalentLoader.LoadTalents()
	if err != nil {
		return err
	}
	log.Info("loading traits")
	_, err = l.TraitLoader.LoadTraits()
	if err != nil {
		return err
	}
	log.Info("loading jobs")
	_, err = l.JobLoader.LoadJobs()
	if err != nil {
		return err
	}
	log.Info("loading teams")
	_, err = l.TeamLoader.LoadTeams()
	if err != nil {
		return err
	}
	log.Info("loading qualities")
	_, err = l.QualityLoader.LoadQualities()
	if err != nil {
		return err
	}
	log.Info("loading weapons")
	_, err = l.EquipmentLoader.LoadWeapons()
	if err != nil {
		return err
	}
	log.Info("loading rooms")
	_, err = l.RoomLoader.LoadRooms()
	if err != nil {
		return err
	}
	log.Info("loading upbringings")
	_, err = l.UpbringingLoader.LoadUpbringings()
	if err != nil {
		return err
	}
	log.Info("loading NPCs")
	_, err = l.NPCLoader.LoadNPCs()
	if err != nil {
		return err
	}
	log.Info("loading actions")
	_, err = l.ActionLoader.LoadActions()
	if err != nil {
		return err
	}
	log.Info("loading conditions")
	_, err = l.ConditionLoader.LoadConditions()
	if err != nil {
		return err
	}
	log.Info("loading tasks")
	_, err = l.TaskLoader.LoadTaskResolvers()
	if err != nil {
		return err
	}
	log.Info("loading methods")
	_, err = l.MethodLoader.LoadMethods(l.TaskLoader)
	if err != nil {
		return err
	}
	log.Info("loading generators")
	_, err = l.GeneratorLoader.LoadGenerators()
	if err != nil {
		return err
	}
	return nil
}
