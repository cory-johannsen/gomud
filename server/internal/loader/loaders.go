package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	AlignmentLoader  *AlignmentLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	JobLoader        *JobLoader
	SkillLoader      *SkillLoader
	TalentLoader     *TalentLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, alignmentLoader *AlignmentLoader, archetypeLoader *ArchetypeLoader,
	backgroundLoader *BackgroundLoader, jobLoader *JobLoader, skillLoader *SkillLoader, talentLoader *TalentLoader,
	traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		AlignmentLoader:  alignmentLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		JobLoader:        jobLoader,
		SkillLoader:      skillLoader,
		TalentLoader:     talentLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
	}
}
