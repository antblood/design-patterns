# Null Object Pattern

## Overview

The Null Object Pattern is a design pattern used to provide an object as a surrogate for the absence of an object of a given type. Instead of using `nil` to represent the absence of an object, a special "null" object is used to provide default (often do-nothing) behavior. This can help reduce the need for explicit `nil` checks and provide a more uniform interface.

## When to Use

The Null Object Pattern is particularly useful in the following scenarios:

- When you want to avoid `nil` checks scattered throughout your code.
- When you want to provide a default behavior that does nothing, but still conforms to an expected interface.
- When you want to simplify client code by providing a uniform interface, regardless of whether an object is "real" or "null".

## Advantages

- **Simplifies Code**: Reduces the number of conditional checks for `nil` objects, making the code cleaner and easier to read.
- **Consistent Interface**: Provides a consistent interface, making it easier for clients to use objects without worrying about whether they are actual implementations or null objects.
- **Decouples Client and Implementation**: The client code can remain unchanged regardless of whether a real object or a null object is used.

## Example Explanation

In this example, we demonstrate the Null Object Pattern using a `Logger` interface in Go. We define two implementations of this interface:

1. **ConsoleLogger**: A concrete implementation of the `Logger` interface that logs messages to the console. It provides actual functionality for the `Info` and `Error` methods.

2. **NullLogger**: A null object implementation of the `Logger` interface. It provides do-nothing implementations for the `Info` and `Error` methods.

### How It Works

- **Logger Interface**: The `Logger` interface defines two methods: `Info` and `Error`.

- **ConsoleLogger**: Implements the `Logger` interface. It logs informational and error messages to the console.

- **NullLogger**: Also implements the `Logger` interface, but its `Info` and `Error` methods do nothing. This fulfills the null object role.

### Usage

In the main program, you can choose between using a `ConsoleLogger` or a `NullLogger`. The client code simply calls `Info` and `Error` on the `Logger` interface without worrying about whether the logger is real or null. If the `NullLogger` is used, these methods will do nothing, effectively ignoring the log requests.

By using the Null Object Pattern, the code remains clean and free from `nil` checks, and it provides a consistent interface for logging, regardless of the actual implementation behind it.

## Conclusion

The Null Object Pattern is a powerful tool in your design pattern toolkit. It helps to simplify your code and make it more robust by eliminating `nil` checks and providing a consistent interface for various implementations, including those that do nothing.
