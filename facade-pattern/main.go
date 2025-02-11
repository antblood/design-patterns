package main

import (
	"fmt"

	"github.com/antblood/design-patterns/facade-pattern/pkg/smarthome"
)

func main() {
	smartHome := smarthome.NewSmartHomeFacade()

	smartHome.MorningRoutine()
	fmt.Println()
	smartHome.NightRoutine()
}
