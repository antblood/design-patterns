package beverages

// Beverage is the base interface that all beverages and decorators implement.
type Beverage interface {
	Description() string
	Cost() float64
}
