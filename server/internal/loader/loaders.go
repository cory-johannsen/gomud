package loader

type Loaders struct {
	AppearanceLoader *AppearanceLoader
	TeamLoader       *TeamLoader
}

func NewLoaders(appearanceLoader *AppearanceLoader, teamLoader *TeamLoader) *Loaders {
	return &Loaders{
		AppearanceLoader: appearanceLoader,
		TeamLoader:       teamLoader,
	}
}
