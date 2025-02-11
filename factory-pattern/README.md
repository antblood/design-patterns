# Factory Pattern in Go

This repository contains a simple implementation of the Factory Pattern using Go. The Factory Pattern is a creational design pattern that provides a way to create objects without specifying the exact class of the object that will be created. This is particularly useful in scenarios where the process of object creation involves complex logic or when the exact type of the object to be created is determined at runtime.

## Overview

In this example, we have an `Animal` interface that declares a `Speak` method. We then have three concrete types: `Dog`, `Cat`, and `Bird`, each implementing the `Animal` interface. The `AnimalFactory` function serves as the factory that creates and returns objects of these types based on input.

## Why Use the Factory Pattern?

1. **Decoupling**: The Factory Pattern decouples the client code from the concrete classes it needs to instantiate, promoting loose coupling and enhancing code maintainability.

2. **Single Responsibility Principle**: The Factory Pattern adheres to the Single Responsibility Principle by centralizing the object creation logic in one place.

3. **Flexibility and Scalability**: It allows the addition of new types of products without changing existing client code, making the system more scalable and flexible.

4. **Encapsulation**: It encapsulates the instantiation logic, which might be complex, inside the factory, reducing code duplication.

## Implementation

Here is a brief look at how the Factory Pattern is implemented in Go:

### Animal Interface

```go
type Animal interface {
    Speak() string
}
```

### Concrete Types

```go
type Dog struct{}

func (d Dog) Speak() string {
    return "Woof!"
}
```
