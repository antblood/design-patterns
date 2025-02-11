package animalfactory

// Cat is another concrete product that implements the Animal interface.
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}
