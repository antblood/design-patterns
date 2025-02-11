package beverages

import "fmt"

// Beverage is the template interface that defines the skeleton of the algorithm.
type Beverage interface {
	PrepareRecipe() // Template method
	BoilWater()
	Brew()
	PourInCup()
	AddCondiments()
}

// TemplateBeverage provides default implementations for some steps.
type TemplateBeverage struct{}

func (t *TemplateBeverage) BoilWater() {
	fmt.Println("Boiling water")
}

func (t *TemplateBeverage) PourInCup() {
	fmt.Println("Pouring into cup")
}
