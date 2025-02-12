package network

import "fmt"

// ListeningState represents a state where the connection is listening
type ListeningState struct {
	connection *TCPConnection
}

func (l *ListeningState) Connect() {
	fmt.Println("Cannot establish a connection while listening.")
}

func (l *ListeningState) Disconnect() {
	fmt.Println("Stopping listening...")
	l.connection.setState(&ClosedState{connection: l.connection})
}

func (l *ListeningState) Listen() {
	fmt.Println("Already listening.")
}
