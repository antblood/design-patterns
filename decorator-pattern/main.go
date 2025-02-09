package main

import (
	"fmt"

	"github.com/antblood/design-patterns/decorator-pattern/beverages"
)

func getColdCoffeeWithEspresso() beverages.Beverage {
	var beverage beverages.Beverage = &beverages.ColdCoffee{}
	beverage = &beverages.Espresso{
		AddOn: beverages.AddOn{
			Beverage: beverage,
		},
	}
	return beverage
}

func getCappuccinoWithMilk() beverages.Beverage {
	var beverage beverages.Beverage = &beverages.Cappuccino{}
	beverage = &beverages.Milk{
		AddOn: beverages.AddOn{
			Beverage: beverage,
		},
	}
	return beverage
}

func main() {
	coldCoffeeWithEspresso := getColdCoffeeWithEspresso()
	fmt.Printf("%s: $%.2f\n", coldCoffeeWithEspresso.Description(), coldCoffeeWithEspresso.Cost())

	cappuccinoWithMilk := getCappuccinoWithMilk()
	fmt.Printf("%s: $%.2f\n", cappuccinoWithMilk.Description(), cappuccinoWithMilk.Cost())
}
