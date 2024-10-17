package domain

import "fmt"

type DisorderSpec struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Effects     []string `yaml:"effects"`
}

type Disorder struct {
	Name        string
	Description string
	Effects     Effects
}

func (d *Disorder) Value() interface{} {
	return d
}

func (d *Disorder) String() string {
	return fmt.Sprintf("%s: %s", d.Name, d.Description)
}

var _ Property = &Disorder{}

type Disorders []*Disorder

func (d Disorders) Value() interface{} {
	return d
}

func (d Disorders) String() string {
	var msg string
	for _, disorder := range d {
		msg += fmt.Sprintf("%s: %s\n", disorder.Name, disorder.Description)
	}
	return msg
}

var _ Property = Disorders{}
