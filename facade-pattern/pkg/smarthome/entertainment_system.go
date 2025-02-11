package smarthome

import "fmt"

// EntertainmentSystem handles the TV and audio systems.
type EntertainmentSystem struct{}

func (e *EntertainmentSystem) PlayMusic() {
	fmt.Println("Playing music.")
}

func (e *EntertainmentSystem) StopMusic() {
	fmt.Println("Stopping music.")
}
