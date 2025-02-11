package smarthome

import "fmt"

// HeatingSystem controls heating and air conditioning.
type HeatingSystem struct{}

func (h *HeatingSystem) SetTemperature(temp int) {
	fmt.Printf("Setting temperature to %d degrees.\n", temp)
}
