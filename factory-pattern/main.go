package main

import (
	"fmt"

	"github.com/antblood/design-patterns/factory-pattern/pkg/animalfactory"
)

func main() {
	// Create a Dog using the factory
	dog, err := animalfactory.AnimalFactory(animalfactory.DogType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Dog:", dog.Speak())

	// Create a Cat using the factory
	cat, err := animalfactory.AnimalFactory(animalfactory.CatType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Cat:", cat.Speak())

	// Create a Bird using the factory
	bird, err := animalfactory.AnimalFactory(animalfactory.BirdType)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Bird:", bird.Speak())

	// Try to create an unknown animal type
	unknown, err := animalfactory.AnimalFactory(animalfactory.AnimalType("fish"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Fish:", unknown.Speak())
}
