package network

import "fmt"

// ClosedState represents a state where the connection is closed
type ClosedState struct {
	connection *TCPConnection
}

func (c *ClosedState) Connect() {
	fmt.Println("Establishing a new connection...")
	c.connection.setState(&EstablishedState{connection: c.connection})
}

func (c *ClosedState) Disconnect() {
	fmt.Println("Connection is already closed.")
}

func (c *ClosedState) Listen() {
	fmt.Println("Listening for incoming connections...")
	c.connection.setState(&ListeningState{connection: c.connection})
}
