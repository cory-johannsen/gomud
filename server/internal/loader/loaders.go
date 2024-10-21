package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	AlignmentLoader  *AlignmentLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	EffectLoader     *EffectLoader
	InjuryLoader     *InjuryLoader
	InventoryLoader  *InventoryLoader
	JobLoader        *JobLoader
	RoomLoader       *RoomLoader
	SkillLoader      *SkillLoader
	TalentLoader     *TalentLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, alignmentLoader *AlignmentLoader, archetypeLoader *ArchetypeLoader,
	backgroundLoader *BackgroundLoader, effectLoader *EffectLoader, injuryLoader *InjuryLoader, inventoryLoader *InventoryLoader, jobLoader *JobLoader, roomLoader *RoomLoader, skillLoader *SkillLoader, talentLoader *TalentLoader,
	traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		AlignmentLoader:  alignmentLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		EffectLoader:     effectLoader,
		InjuryLoader:     injuryLoader,
		InventoryLoader:  inventoryLoader,
		JobLoader:        jobLoader,
		RoomLoader:       roomLoader,
		SkillLoader:      skillLoader,
		TalentLoader:     talentLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
	}
}
