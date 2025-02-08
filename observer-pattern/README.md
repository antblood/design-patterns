# Observer Pattern Example

This project demonstrates the Observer Design Pattern. The Observer Pattern is a behavioral design pattern that allows an object, called the subject, to maintain a list of its dependents, called observers, and notify them automatically of any state changes, usually by calling one of their methods.

## Project Overview

In this implementation:
- The **Subject** changes its state every 3 seconds.
- When the state changes, the subject broadcasts a notification to all its observers.
- Upon receiving a notification, each **Observer** can request and retrieve the latest state from the subject.

## Components

### Subject

- **Responsibilities:**
  - Maintain a list of observers.
  - Change its state every 3 seconds.
  - Notify all registered observers when its state changes.

- **Methods:**
  - `Add()`: Adds an observer to the list.
  - `Remove()`: Removes an observer from the list.
  - `Broadcast()`: Notifies all observers that the state has changed.
  - `LatestState()`: Returns the current state.
  - `UpdateState()`: Method to change the state.

### Observer

- **Responsibilities:**
  - Register itself with the subject to receive updates.
  - Request the latest state from the subject when notified.
  
- **Methods:**
  - `Notify()`: Called by the subject to notify the observer of a state change. In the same Requests the latest state from the subject.

## How It Works

1. **Subject Initialization**: The subject starts with a modifiable list of observer URLs.
2. **Observer Initialization**: Observers starts with the subject URL.
3. **State Change**: The subject changes its state and broadcasts the change to all observers.
4. **Notification**: Each observer's `Notify()` method is triggered and the observer requests the latest state from the subject within the same method.

## Getting Started

To run this project, follow the steps below:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/antblood/design-patterns.git
   cd design-patterns/observer-pattern
   go mod tidy
   ```

2. **Run the Subject**:

   ```bash
   go run cmd/subject/main.go
   ```

3. **Run the Observers**:

   ```bash
   go run cmd/observer/main.go
   ```
