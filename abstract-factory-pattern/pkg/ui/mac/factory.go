package mac

import (
	"github.com/antblood/design-patterns/abstract-factory-pattern/pkg/ui/abstract"
)

// MacFactory is a concrete factory for creating macOS components.
type MacFactory struct{}

func (f *MacFactory) CreateButton() abstract.Button {
	return &MacButton{}
}

func (f *MacFactory) CreateTextField() abstract.TextField {
	return &MacTextField{}
}
