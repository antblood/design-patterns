package main

import (
	"fmt"

	"github.com/antblood/design-patterns/null-object-pattern/pkg/logger"
)

func main() {
	// You can switch between ConsoleLogger and NullLogger
	fmt.Println("Using ConsoleLogger:")
	var loggerVar logger.Logger = &logger.ConsoleLogger{}
	loggerVar.Info("This is an info message")
	loggerVar.Error("This is an error message")

	// If NullLogger is used, no output will be produced
	fmt.Println("Using NullLogger:")
	loggerVar = &logger.NullLogger{}
	loggerVar.Info("This is an info message")
	loggerVar.Error("This is an error message")
}
