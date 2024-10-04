package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	BackgroundLoader *BackgroundLoader
	TeamLoader       *TeamLoader
	TraitLoader      *TraitLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, backgroundLoader *BackgroundLoader, traitLoader *TraitLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		BackgroundLoader: backgroundLoader,
		TeamLoader:       teamLoader,
		TraitLoader:      traitLoader,
	}
}
