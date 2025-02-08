package cart

type OnlineCart struct{}

func (c *OnlineCart) CalculateTotal(items []Item) float64 {
	total := 0.0
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}
