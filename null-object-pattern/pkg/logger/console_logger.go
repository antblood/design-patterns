package logger

import "fmt"

// ConsoleLogger is a concrete implementation of Logger that logs to the console.
type ConsoleLogger struct{}

// Info logs an informational message to the console.
func (c *ConsoleLogger) Info(message string) {
	fmt.Println("INFO:", message)
}

// Error logs an error message to the console.
func (c *ConsoleLogger) Error(message string) {
	fmt.Println("ERROR:", message)
}
