# Bridge Pattern

## Introduction to the Bridge Pattern

The Bridge pattern is a structural design pattern that aims to decouple an abstraction from its implementation, allowing both to vary independently. It is particularly useful when you want to avoid a permanent binding between an abstraction and its implementation, making it easier to extend the system with new abstraction and implementation classes.
The term "Bridge" in the Bridge pattern refers to the way the pattern functions as a bridge between two hierarchies: the abstraction and the implementation. The key idea is that the pattern bridges the gap between these two separate parts of your code structure, allowing them to evolve independently.

## Why is it Called the "Bridge" Pattern?
- **Decoupling**: The Bridge pattern decouples the abstraction (what you are working with) from its implementation (how it is done), much like a bridge connects two separate land masses. This decoupling allows changes to be made on either side without affecting the other side, providing flexibility.
- **Connecting Abstraction and Implementation**: The pattern connects an abstraction to its implementation using a bridge interface. This bridge allows the abstraction to delegate the implementation details to another class rather than handling them internally.
- **Independent Evolution**: Much like a bridge allows traffic to flow independently of the terrain it crosses, the Bridge pattern allows both the abstraction and the implementation to change independently. New abstractions can be introduced without changing existing implementations, and new implementations can be added without altering existing abstractions.

### When to Use the Bridge Pattern

- When you want to separate an abstraction from its implementation so that both can be changed independently.
- When you have a proliferation of classes caused by a combination of different abstractions and implementations.
- When you want to share an implementation among multiple objects, and this implementation should be easily replaceable.

### Advantages of the Bridge Pattern

- **Decoupling Interface and Implementation**: The bridge pattern separates the abstraction from its implementation, allowing both to evolve independently.
- **Improved Flexibility and Scalability**: Adding new abstractions and implementations becomes easier without affecting existing code.
- **Enhanced Maintainability**: Changes in implementations have minimal impact on abstractions and vice versa, simplifying maintenance.
- **Reduced Complexity**: It reduces the number of subclasses and interfaces by separating concerns.

## Example Explanation

This repository contains an example of the Bridge pattern implemented in Go. The example involves different types of documents that can be formatted using different formatters.

### Components

1. **Abstraction (`Document`)**: An interface representing the concept of a document with methods to set a formatter and print content.

2. **Implementor (`Formatter`)**: An interface that allows for different formatting strategies.

3. **Refined Abstractions (`Report`, `Book`)**: Specific types of documents that implement the `Document` interface.

4. **Concrete Implementors (`PlainTextFormatter`, `HTMLFormatter`)**: Classes that implement different formatting strategies.

5. **Bridge**: The connection is established using the `SetFormatter` method, which allows a `Document` to use any `Formatter`.

### How It Works

- We define a `Document` interface with methods to set a formatter and print content.
- A `Formatter` interface is defined to allow different formatting strategies.
- Concrete document types (`Report` and `Book`) implement the `Document` interface.
- Concrete formatters (`PlainTextFormatter` and `HTMLFormatter`) implement the `Formatter` interface.
- The bridge is created by setting a formatter to a document using the `SetFormatter` method.
