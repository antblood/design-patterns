package smarthome

import "fmt"

// LightingSystem controls the lights in the house.
type LightingSystem struct{}

func (l *LightingSystem) On() {
	fmt.Println("Turning on the lights.")
}

func (l *LightingSystem) Off() {
	fmt.Println("Turning off the lights.")
}
