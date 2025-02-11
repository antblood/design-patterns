# Paint Application with Command Pattern

This project demonstrates a simple paint application utilizing the Command Pattern in Go. The application allows you to add shapes to a drawing, and supports undo and redo operations to maintain the application state.

## Overview

The application simulates a basic drawing tool where you can add various shapes to a canvas. It employs the Command Pattern to encapsulate requests as objects, thereby allowing for parameterization of methods with different requests, queuing of requests, and support for undoable operations.

### Components

1. **Command Interface**: 
   - Consists of methods `Execute`, `Undo`, and `Redo`.
   - Each command knows how to execute itself and undo its operation.

2. **Concrete Commands**:
   - `AddShapeCommand`: Implements adding a shape to the drawing and knows how to undo this action.

3. **Receiver**:
   - `Drawing`: Maintains the current state of the canvas. It can add and remove shapes.

4. **Invoker**:
   - `CommandManager`: Handles the execution of commands and manages the history for undo and redo operations.

### How It Works

- **Execution**: You can execute commands to modify the drawing. Each command is an operation like adding a shape.
- **Undo/Redo**: The `CommandManager` maintains a history of executed commands. You can undo the last operation or redo an undone operation.

### Why Use the Command Pattern?

The Command Pattern is particularly useful in scenarios where you need to:
- Decouple the object that invokes the operation from the object that knows how to execute it.
- Implement undo and redo features for commands.
- Queue operations and execute them at different times (e.g., batch processing).
- Log changes for audit or debugging purposes.

In this paint application, the Command Pattern allows us to easily add new operations (commands) and manage the state changes with undo and redo functionality, making the application flexible and easy to extend.

## Usage

1. **Add Shapes**: Use `AddShapeCommand` to add shapes to the drawing.
2. **Undo**: Call the `Undo` method in `CommandManager` to revert the last operation.
3. **Redo**: Call the `Redo` method in `CommandManager` to reapply the last undone operation.

## Running the Application

To run the application, make sure you have Go installed on your machine. Clone this repository and execute the main program:

```bash
go run main.go
```