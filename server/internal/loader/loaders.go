package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	AlignmentLoader  *AlignmentLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	JobLoader        *JobLoader
	SkillLoader      *SkillLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, alignmentLoader *AlignmentLoader, archetypeLoader *ArchetypeLoader,
	backgroundLoader *BackgroundLoader, jobLoader *JobLoader, skillLoader *SkillLoader, traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		AlignmentLoader:  alignmentLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		JobLoader:        jobLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
		SkillLoader:      skillLoader,
	}
}
