package beverages

// AddOn is a base struct for beverage decorators.
type AddOn struct {
	Beverage Beverage
}

// Milk is a concrete decorator that adds milk to a beverage.
type Milk struct {
	AddOn
}

func (m *Milk) Description() string {
	return m.Beverage.Description() + ", Milk"
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 1.00
}

// Espresso is another concrete decorator that adds espresso to a beverage.
type Espresso struct {
	AddOn
}

func (e *Espresso) Description() string {
	return e.Beverage.Description() + ", Espresso"
}

func (e *Espresso) Cost() float64 {
	return e.Beverage.Cost() + 1.00
}
