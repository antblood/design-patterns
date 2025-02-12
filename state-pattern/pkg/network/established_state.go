package network

import "fmt"

// EstablishedState represents a state where the connection is established
type EstablishedState struct {
	connection *TCPConnection
}

func (e *EstablishedState) Connect() {
	fmt.Println("Connection is already established.")
}

func (e *EstablishedState) Disconnect() {
	fmt.Println("Disconnecting the connection...")
	e.connection.setState(&ClosedState{connection: e.connection})
}

func (e *EstablishedState) Listen() {
	fmt.Println("Cannot listen while the connection is established.")
}
