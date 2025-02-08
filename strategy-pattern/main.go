package main

import (
	"fmt"

	"github.com/antblood/strategy-pattern/cart"
	"github.com/antblood/strategy-pattern/shop"
)

func main() {
	shopper := shop.NewStoreCreditCardShopper()
	message, err := shopper.Checkout([]cart.Item{{Name: "Item 1", Price: 100, Quantity: 1}, {Name: "Item 2", Price: 200, Quantity: 1}})
	if err != nil {
		fmt.Println("error while checkout %w", err)
	}
	fmt.Println(message)
}
