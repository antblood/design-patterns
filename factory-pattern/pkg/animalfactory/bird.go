package animalfactory

// Bird is another concrete product that implements the Animal interface.
type Bird struct{}

func (b Bird) Speak() string {
	return "Tweet!"
}
