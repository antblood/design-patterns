package animalfactory

// Dog is a concrete product that implements the Animal interface.
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}
