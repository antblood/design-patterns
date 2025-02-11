package animalfactory

import "fmt"

type AnimalType string

const (
	DogType  AnimalType = "dog"
	CatType  AnimalType = "cat"
	BirdType AnimalType = "bird"
)

// AnimalFactory is the factory function that returns an Animal based on a given type.
func AnimalFactory(animalType AnimalType) (Animal, error) {
	switch animalType {
	case DogType:
		return Dog{}, nil
	case CatType:
		return Cat{}, nil
	case BirdType:
		return Bird{}, nil
	default:
		return nil, fmt.Errorf("unknown animal type: %s", animalType)
	}
}
