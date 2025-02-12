package main

import (
	"github.com/antblood/design-patterns/state-pattern/pkg/network"
)

func main() {
	connection := network.NewTCPConnection()

	connection.Connect()    // Establishes the connection
	connection.Listen()     // Fails because it is already established
	connection.Disconnect() // Closes the connection
	connection.Listen()     // Starts listening
	connection.Connect()    // Fails because it is listening
	connection.Disconnect() // Stops listening and closes the connection
}
