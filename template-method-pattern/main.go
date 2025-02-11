package main

import (
	"fmt"

	"github.com/antblood/design-patterns/template-method-pattern/pkg/beverages"
)

func main() {
	tea := &beverages.Tea{}
	coffee := &beverages.Coffee{}

	fmt.Println("Making tea:")
	tea.PrepareRecipe()

	fmt.Println("\nMaking coffee:")
	coffee.PrepareRecipe()
}
