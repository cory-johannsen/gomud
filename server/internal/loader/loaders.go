package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	AlignmentLoader  *AlignmentLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	InjuryLoader     *InjuryLoader
	JobLoader        *JobLoader
	RoomLoader       *RoomLoader
	SkillLoader      *SkillLoader
	TalentLoader     *TalentLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, alignmentLoader *AlignmentLoader, archetypeLoader *ArchetypeLoader,
	backgroundLoader *BackgroundLoader, injuryLoader *InjuryLoader, jobLoader *JobLoader, roomLoader *RoomLoader, skillLoader *SkillLoader, talentLoader *TalentLoader,
	traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		AlignmentLoader:  alignmentLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		InjuryLoader:     injuryLoader,
		JobLoader:        jobLoader,
		RoomLoader:       roomLoader,
		SkillLoader:      skillLoader,
		TalentLoader:     talentLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
	}
}
