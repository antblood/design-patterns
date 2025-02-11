package smarthome

import "fmt"

// SmartHomeFacade provides a simple interface to control the smart home.
type SmartHomeFacade struct {
	lighting      *LightingSystem
	security      *SecuritySystem
	heating       *HeatingSystem
	entertainment *EntertainmentSystem
}

func NewSmartHomeFacade() *SmartHomeFacade {
	return &SmartHomeFacade{
		lighting:      &LightingSystem{},
		security:      &SecuritySystem{},
		heating:       &HeatingSystem{},
		entertainment: &EntertainmentSystem{},
	}
}

func (f *SmartHomeFacade) MorningRoutine() {
	fmt.Println("Starting morning routine...")
	f.lighting.On()
	f.security.Disarm()
	f.heating.SetTemperature(22)
	f.entertainment.PlayMusic()
}

func (f *SmartHomeFacade) NightRoutine() {
	fmt.Println("Starting night routine...")
	f.entertainment.StopMusic()
	f.lighting.Off()
	f.security.Arm()
	f.heating.SetTemperature(18)
}

