# Iterator Pattern

## What is the Iterator Pattern?
The Iterator pattern allows you to traverse a collection of elements without exposing the underlying details of the collection. It encapsulates the iteration logic within an object, providing a clean interface to move through the elements of a collection.

## Introduction
This repository contains an example implementation of the Iterator pattern in Go. The Iterator pattern is a behavioral design pattern that provides a way to access elements of a collection sequentially without exposing its underlying representation.

## When to Use the Iterator Pattern
The Iterator pattern is useful in the following scenarios:
- When you need to traverse a collection without exposing its internal structure.
- When you want to provide multiple ways to traverse a collection.
- When you need a uniform interface for iterating over different types of collections.

## Advantages of the Iterator Pattern
- **Encapsulation**: It hides the implementation details of the collection, providing a clean and simple interface for iteration.
- **Flexibility**: You can implement different iteration strategies by simply modifying the iterator's logic.
- **Consistency**: It provides a consistent way to traverse different collections, making it easier to work with various data structures.

## Example Explanation
This repository contains examples of iterators for different data structures in Go: a slice, a set, and a map.

1. **Slice Iterator**: 
   - A simple iterator for a slice of integers. It maintains an index to track the current position and provides methods to check for more elements and retrieve the next element.

2. **Set Iterator**:
   - This example demonstrates an iterator for a set of integers. The set is implemented using a map to ensure uniqueness. The iterator extracts keys from the map and iterates over them.

3. **Map Iterator**:
   - An iterator for a map of integers to strings. It collects keys and values into separate slices and iterates over them, allowing access to both keys and values during iteration.

Each example demonstrates how to implement an iterator for a different type of data structure, highlighting the flexibility and utility of the Iterator pattern in Go.
