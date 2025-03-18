package generator

import (
	"github.com/cory-johannsen/gomud/internal/domain/htn"
)

type DomainGenerator struct {
	domains htn.Domains
}

func (s *DomainGenerator) GetDomain(name string) (*htn.Domain, error) {
	if domain, ok := s.domains[name]; ok {
		return domain, nil
	}
	return nil, nil
}

func (s *DomainGenerator) AddDomain(name string, domain *htn.Domain) {
	s.domains[name] = domain
}

func (s *DomainGenerator) DeleteDomain(name string) {
	delete(s.domains, name)
}

var _ htn.DomainResolver = &DomainGenerator{}

func NewDomainGenerator() *DomainGenerator {
	return &DomainGenerator{
		domains: make(htn.Domains),
	}
}
