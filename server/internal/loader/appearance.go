package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

type AppearanceLoader struct {
	config       *config.Config
	tatLocations domain.TattooLocations
	tats         domain.SeasonalTattoos
	drawbacks    domain.Drawbacks
	marks        domain.DistinguishingMarks
	effectLoader *EffectLoader
}

func NewAppearanceLoader(cfg *config.Config, effectLoader *EffectLoader) *AppearanceLoader {
	return &AppearanceLoader{
		config:       cfg,
		effectLoader: effectLoader,
	}
}

func (l *AppearanceLoader) LoadTattooLocations() (domain.TattooLocations, error) {
	if l.tatLocations != nil {
		return l.tatLocations, nil
	}
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/tattoo_locations.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &l.tatLocations)
	if err != nil {
		return nil, err
	}
	return l.tatLocations, nil
}

func (l *AppearanceLoader) LoadTattoos() (domain.SeasonalTattoos, error) {
	if l.tats != nil {
		return l.tats, nil
	}
	tatsBySeason := make(map[domain.Season][]string)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/tattoos.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &tatsBySeason)
	if err != nil {
		return nil, err
	}
	tats := make(domain.SeasonalTattoos)
	for season, seasonTats := range tatsBySeason {
		tats[season] = make(domain.Tattoos, 0)
		for _, tatDescription := range seasonTats {
			tat := domain.Tattoo{
				Description: tatDescription,
				Season:      season,
			}
			tats[season] = append(tats[season], tat)
		}
	}
	l.tats = tats
	return tats, nil
}

func (l *AppearanceLoader) LoadDrawbacks() (domain.Drawbacks, error) {
	if l.drawbacks != nil {
		return l.drawbacks, nil
	}
	specs := make(domain.DrawbackSpecs, 0)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/drawbacks.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &specs)
	if err != nil {
		return nil, err
	}

	drawbacks := make(domain.Drawbacks, 0)
	for _, d := range specs {
		effect, err := l.effectLoader.GetEffect(d.Name)
		if err != nil {
			log.Errorf("error loading effect %s: %v", d.Name, err)
			continue
		}
		drawback := &domain.Drawback{
			Name:        d.Name,
			Description: d.Description,
			Effect:      effect,
		}
		drawbacks = append(drawbacks, drawback)
	}

	l.drawbacks = drawbacks
	return drawbacks, nil
}

func (l *AppearanceLoader) GetDrawback(name string) (*domain.Drawback, error) {
	if l.drawbacks == nil {
		_, err := l.LoadDrawbacks()
		if err != nil {
			return nil, err
		}
	}
	for _, d := range l.drawbacks {
		if d.Name == name {
			return d, nil
		}
	}
	return nil, nil
}

func (l *AppearanceLoader) LoadDistinguishingMarks() (domain.DistinguishingMarks, error) {
	if l.marks != nil {
		return l.marks, nil
	}
	marks := make(domain.DistinguishingMarks, 0)
	data, err := os.ReadFile(l.config.AssetPath + "/appearance/distinguishing_marks.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &marks)
	if err != nil {
		return nil, err
	}
	l.marks = marks
	return marks, nil
}
