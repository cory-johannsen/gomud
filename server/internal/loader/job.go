package loader

import (
	"github.com/cory-johannsen/gomud/internal/config"
	"github.com/cory-johannsen/gomud/internal/domain"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
)

type JobLoader struct {
	config          *config.Config
	jobs            domain.Jobs
	archetypeLoader *ArchetypeLoader
	traitLoader     *TraitLoader
}

func NewJobLoader(cfg *config.Config, archetypeLoader *ArchetypeLoader, traitLoader *TraitLoader) *JobLoader {
	return &JobLoader{
		config:          cfg,
		jobs:            make(domain.Jobs, 0),
		archetypeLoader: archetypeLoader,
		traitLoader:     traitLoader,
	}
}

func (l *JobLoader) LoadJobs() (domain.Jobs, error) {
	if len(l.jobs) > 0 {
		return l.jobs, nil
	}
	items, err := os.ReadDir(l.config.AssetPath + "/jobs")
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.IsDir() {
			continue
		}
		if item.Name() == "tmpl.yaml" {
			continue
		}
		log.Printf("loading job %s", item.Name())
		spec := &domain.JobSpec{}
		data, err := os.ReadFile(l.config.AssetPath + "/jobs/" + item.Name())
		if err != nil {
			return nil, err
		}
		err = yaml.Unmarshal(data, spec)
		if err != nil {
			return nil, err
		}

		archetype, err := l.archetypeLoader.GetArchetype(spec.Archetype)
		if err != nil {
			return nil, err
		}
		if archetype == nil {
			log.Printf("could not find archetype %s for job %s", spec.Archetype, spec.Name)
		}
		traits := make(domain.Traits, 0)
		for _, traitName := range spec.Traits {
			trait, err := l.traitLoader.GetTrait(traitName)
			if err != nil {
				return nil, err
			}
			if trait == nil {
				log.Printf("could not find trait %s for job %s", traitName, spec.Name)
				continue
			}
			traits = append(traits, trait)
		}
		job := &domain.Job{
			Name:           spec.Name,
			Description:    spec.Description,
			Archetype:      archetype,
			Tier:           spec.Tier,
			ExperienceCost: spec.ExperienceCost,
			Traits:         traits,
		}
		l.jobs = append(l.jobs, job)
	}
	return l.jobs, nil
}

func (l *JobLoader) GetJob(name string) (*domain.Job, error) {
	jobs, err := l.LoadJobs()
	if err != nil {
		return nil, err
	}
	for _, job := range jobs {
		if job.Name == name {
			return job, nil
		}
	}
	return nil, nil
}

func (l *JobLoader) RandomJob(archetype *domain.Archetype) (*domain.Job, error) {
	jobs, err := l.LoadJobs()
	if err != nil {
		return nil, err
	}
	archetypeJobs := make(domain.Jobs, 0)
	for _, job := range jobs {
		if job.Archetype.Name == archetype.Name {
			archetypeJobs = append(archetypeJobs, job)
		}
	}

	return archetypeJobs[rand.Intn(len(archetypeJobs))], nil
}
