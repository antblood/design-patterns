package main

import (
	"fmt"

	"github.com/antblood/design-patterns/abstract-factory-pattern/pkg/ui/abstract"
	"github.com/antblood/design-patterns/abstract-factory-pattern/pkg/ui/linux"
	"github.com/antblood/design-patterns/abstract-factory-pattern/pkg/ui/mac"
)

// Client code
func main() {
	var factory abstract.GUIFactory

	fmt.Println("Mac GUI")
	factory = &mac.MacFactory{}
	fmt.Println(factory.CreateButton().Render())
	fmt.Println(factory.CreateTextField().Render())

	fmt.Println("Linux GUI")
	factory = &linux.LinuxFactory{}
	fmt.Println(factory.CreateButton().Render())
	fmt.Println(factory.CreateTextField().Render())
}
