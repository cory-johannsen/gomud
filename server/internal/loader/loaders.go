package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	ArchetypeLoader  *ArchetypeLoader
	BackgroundLoader *BackgroundLoader
	JobLoader        *JobLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, archetypeLoader *ArchetypeLoader, backgroundLoader *BackgroundLoader, jobLoader *JobLoader, traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		ArchetypeLoader:  archetypeLoader,
		BackgroundLoader: backgroundLoader,
		JobLoader:        jobLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
	}
}
