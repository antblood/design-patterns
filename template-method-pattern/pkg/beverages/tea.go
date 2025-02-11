package beverages

import "fmt"

// Tea is a concrete implementation of Beverage.
type Tea struct {
	TemplateBeverage // Embedding TemplateBeverage to use default methods
}

func (t *Tea) PrepareRecipe() {
	t.BoilWater()
	t.Brew()
	t.PourInCup()
	t.AddCondiments()
}

func (t *Tea) Brew() {
	fmt.Println("Steeping the tea")
}

func (t *Tea) AddCondiments() {
	fmt.Println("Adding lemon")
}
