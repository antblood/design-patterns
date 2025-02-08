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
  - `UpdateState()`: Method that changes the state every 3 seconds and calls `Broadcast()`.

### Observer

- **Responsibilities:**
  - Register itself with the subject to receive updates.
  - Request the latest state from the subject when notified.
  
- **Methods:**
  - `Notify()`: Called by the subject to notify the observer of a state change.
  - `LatestState()`: Requests the latest state from the subject.

## How It Works

1. **Initialization**: The subject is initialized and starts changing its state every 3 seconds.
2. **Observer Registration**: Observers register themselves with the subject.
3. **State Change**: The subject changes its state and calls `notify_observers()`.
4. **Notification**: Each observer's `update()` method is triggered.
5. **State Request**: Observers call `get_latest_state()` to get the updated state from the subject.

## Getting Started

To run this project, follow the steps below:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/antblood/design-patterns.git
   cd design-patterns/observer-pattern
   ```

2. **Run the Subject**:

   ```bash
   go run cmd/subject/main.go
   ```

3. **Run the Observers**:

   ```bash
   go run cmd/observer/main.go
