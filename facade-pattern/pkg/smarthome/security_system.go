package smarthome

import "fmt"

// SecuritySystem manages security alarms and cameras.
type SecuritySystem struct{}

func (s *SecuritySystem) Arm() {
	fmt.Println("Arming the security system.")
}

func (s *SecuritySystem) Disarm() {
	fmt.Println("Disarming the security system.")
}
