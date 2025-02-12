package logger

// NullLogger is a null object implementation of Logger that does nothing.
type NullLogger struct{}

// Info does nothing.
func (n *NullLogger) Info(message string) {}

// Error does nothing.
func (n *NullLogger) Error(message string) {}
