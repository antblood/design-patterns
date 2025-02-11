package linux

import (
	"github.com/antblood/design-patterns/abstract-factory-pattern/pkg/ui/abstract"
)

// LinuxFactory is a concrete factory for creating Linux components.
type LinuxFactory struct{}

func (f *LinuxFactory) CreateButton() abstract.Button {
	return &LinuxButton{}
}

func (f *LinuxFactory) CreateTextField() abstract.TextField {
	return &LinuxTextField{}
}
