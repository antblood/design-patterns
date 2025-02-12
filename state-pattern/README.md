# State Pattern

## What is the State Pattern?

The State Pattern allows an object to change its behavior dynamically at runtime by altering its internal state. The pattern encapsulates state-specific behavior and delegates behavior to the current state object.

## Repo Overview

This repository contains an example implementation of the State Pattern in the Go programming language. The State Pattern is a behavioral design pattern that enables an object to change its behavior when its internal state changes. This pattern is particularly useful when an object needs to exhibit different behaviors based on its current state.

### Key Concepts

- **State Interface**: Defines the interface for encapsulating the behavior associated with a particular state of the context.
- **Concrete States**: Implementations of the State interface for each specific state.
- **Context**: Maintains an instance of a ConcreteState subclass that defines the current state.

## When to Use the State Pattern

- When an object's behavior depends on its state, and it must change its behavior at runtime based on that state.
- When operations have large, multipart conditional statements that depend on the object's state.
- When you have multiple states and want to avoid complex conditional logic.

## Advantages of the State Pattern

- **Encapsulation**: State-specific behavior is encapsulated within concrete state classes, making the system easier to understand and maintain.
- **Simplifies Control Logic**: Removes complex conditional logic from the context, simplifying code maintenance.
- **Extensibility**: Adding new states is straightforward and does not require altering existing state classes or the context.

## Example Overview

In this example, we simulate a simple TCP connection using the State Pattern. The connection can be in one of three states: `Established`, `Closed`, and `Listening`. Each state has specific behaviors for the actions `Connect`, `Disconnect`, and `Listen`.

### Components

- **State Interface**: Defines methods `Connect()`, `Disconnect()`, and `Listen()` to be implemented by all concrete states.
- **Concrete States**:
  - `EstablishedState`: Represents a state where the connection is established.
  - `ClosedState`: Represents a state where the connection is closed.
  - `ListeningState`: Represents a state where the connection is listening for incoming connections.
- **TCPConnection (Context)**: Maintains an instance of a state and delegates state-specific behavior to the current state.

### Behavior

- The `TCPConnection` starts in the `Closed` state.
- Calling `Connect()` in the `Closed` state transitions the connection to the `Established` state.
- Calling `Disconnect()` in the `Established` state transitions the connection back to the `Closed` state.
- Calling `Listen()` in the `Closed` state transitions the connection to the `Listening` state.
- Each state has its own logic for handling actions, ensuring that the connection behaves correctly based on its current state.

## How to Run the Example

1. Make sure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the directory containing the example code.
4. Run the code using the command:

   ```bash
   go run main.go
