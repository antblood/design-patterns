# Composite Pattern

## What is the Composite Pattern?

The Composite Pattern is designed to simplify client interaction with complex tree structures, enabling clients to treat individual objects and compositions of objects in a uniform manner. This pattern is particularly useful for scenarios where you need to represent hierarchies, such as file systems, UI components, or document structures.


## Overview

This repository contains an example implementation of the Composite Pattern in Go, modeled as a simple file system. The Composite Pattern is a structural design pattern that enables you to compose objects into tree structures to represent part-whole hierarchies. It allows clients to treat individual objects and compositions of objects uniformly.

### Key Concepts

- **Component**: An interface that defines the operations applicable to both simple and complex objects of the composition.
- **Leaf**: Represents end objects of a composition with no children. Implements the component interface.
- **Composite**: Represents complex objects that can have children. Implements the component interface and stores child components.
- **Client**: Interacts with objects in the composition through the component interface.

## When to Use the Composite Pattern

- When you need to represent part-whole hierarchies.
- When you want clients to be able to ignore the difference between compositions of objects and individual objects.
- When dealing with tree structures where nodes and leaves need to be managed uniformly.

## Advantages of the Composite Pattern

- **Uniformity**: Treats individual objects and compositions uniformly, simplifying client code.
- **Flexibility**: Easily extendable to include new types of components, whether they are leaves or composites.
- **Simplicity**: Reduces complexity by allowing complex tree structures and simple leaf nodes to be treated uniformly.

## Example Explanation

In this example, we implement a simple file system using the Composite Pattern:

- **Component Interface**: The `Component` interface defines a `Show` method that all components (both files and directories) must implement.
  
- **Leaf (File)**: The `File` struct represents a file in the file system, which is a leaf node. It implements the `Show` method to display its name with proper indentation.

- **Composite (Directory)**: The `Directory` struct represents a directory that can contain other files or directories. It has a `Name` and a slice of `Children` that implements the `Component` interface. It also provides an `Add` method to add components to the directory and a `Show` method to display its contents recursively.

- **Client Code**: In the `main` function, we create files and directories, build a hierarchical structure by adding files to directories, and then display the entire structure using the `Show` method. This demonstrates how the Composite Pattern allows you to manage both files and directories uniformly.

This example showcases the effectiveness of the Composite Pattern in handling hierarchical data structures, providing a clean and flexible way to manage part-whole hierarchies in a file system context.
