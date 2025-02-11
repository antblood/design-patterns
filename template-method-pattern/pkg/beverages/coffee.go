package beverages

import "fmt"

// Coffee is another concrete implementation of Beverage.
type Coffee struct {
	TemplateBeverage // Embedding TemplateBeverage to use default methods
}

func (c *Coffee) PrepareRecipe() {
	c.BoilWater()
	c.Brew()
	c.PourInCup()
	c.AddCondiments()
}

func (c *Coffee) Brew() {
	fmt.Println("Dripping coffee through filter")
}

func (c *Coffee) AddCondiments() {
	fmt.Println("Adding sugar and milk")
}
