package network

// State interface
type State interface {
	Connect()
	Disconnect()
	Listen()
}
