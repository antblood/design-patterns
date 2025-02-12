package network

// TCPConnection holds a reference to the current state
type TCPConnection struct {
	state State
}

// NewTCPConnection initializes the TCP connection with a Closed state
func NewTCPConnection() *TCPConnection {
	connection := &TCPConnection{}
	closedState := &ClosedState{connection: connection}
	connection.setState(closedState)
	return connection
}

// setState changes the current state of the TCP connection
func (c *TCPConnection) setState(state State) {
	c.state = state
}

// Connect delegates the action to the current state
func (c *TCPConnection) Connect() {
	c.state.Connect()
}

// Disconnect delegates the action to the current state
func (c *TCPConnection) Disconnect() {
	c.state.Disconnect()
}

// Listen delegates the action to the current state
func (c *TCPConnection) Listen() {
	c.state.Listen()
}
