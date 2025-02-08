package cart

type Item struct {
	Name     string
	Price    float64
	Quantity int
}

type CartTypeHandler interface {
	CalculateTotal(items []Item) float64
}
