# Proxy Pattern Example in Go

This repository contains a simple implementation of the Proxy Design Pattern in Go. The Proxy Pattern is used to provide a surrogate or placeholder for another object to control access to it. This can be useful in situations such as lazy loading, access control, logging, etc.


## Advantages of the Proxy Pattern

1. **Lazy Initialization**: Resources are loaded only when they are needed, which can improve performance and resource usage.
2. **Access Control**: Proxies can control access to the real object, which can be useful for security purposes.
3. **Logging and Auditing**: Proxies can be used to add logging or other auditing tasks around method calls.
4. **Remote Proxy**: Can represent an object in a different address space, useful in distributed systems.

## Use Case

In this example, we implement a proxy for an image loader. The objective is to delay the loading of an image until it is actually needed. This is beneficial in scenarios where loading resources is costly and we want to optimize resource usage.

## Implementation

- We define an interface `Image` with a method `Display()`.
- `image` is a concrete class that implements the `Image` interface. It represents the actual image and contains logic to load and display the image.
- `ExternalImage` is another concrete class implementing the `Image` interface. It acts as a proxy to `image`. It controls access to the `image` by loading it only when `Display()` is called for the first time. Also, it has a method `deleteFromDisk()` that isn't accessible from the outside of the package.

## Code Structure

- **`Image` interface**: Defines the contract for `Display()`.
- **`image` struct**: Implements the `Image` interface and handles image loading, displaying and deleting.
- **`ExternalImage` struct**: Also implements the `Image` interface. It maintains a reference to a `image` and controls its instantiation.
