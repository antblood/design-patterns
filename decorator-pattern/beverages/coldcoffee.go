package beverages

// ColdCoffee is a concrete implementation of a Beverage.
type ColdCoffee struct{}

func (c *ColdCoffee) Description() string {
	return "Cold Coffee"
}

func (c *ColdCoffee) Cost() float64 {
	return 3.00
}
