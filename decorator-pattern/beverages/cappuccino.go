package beverages

// Cappuccino is a concrete implementation of a Beverage.
type Cappuccino struct{}

func (c *Cappuccino) Description() string {
	return "Cappuccino"
}

func (c *Cappuccino) Cost() float64 {
	return 2.00
}
