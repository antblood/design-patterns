# Template Method

## Overview

The Template Method Pattern is a behavioral design pattern that defines the skeleton of an algorithm in a method, deferring some steps to subclasses. It allows subclasses to redefine certain steps of an algorithm without changing the algorithm's structure.

## When to Use the Template Method Pattern

- When you have algorithms that share a common structure but have some steps that require custom implementations.
- When you want to enforce a sequence of steps in an algorithm but allow subclasses to provide specific implementations for some steps.
- When you want to avoid code duplication by implementing shared behavior in a common superclass.

## Advantages of the Template Method Pattern

- **Code Reusability**: Shared behavior is implemented in a base class, reducing code duplication.
- **Flexibility**: Subclasses can provide custom behavior for certain steps of the algorithm.
- **Maintainability**: Changes to the algorithm's structure only need to be made in one place, the base class.

## Example Explanation

In this example, we demonstrate the Template Method Pattern by modeling the process of preparing a beverage. We have defined a template for preparing a beverage and implemented two concrete classes—`Tea` and `Coffee`.

### Components

1. **Beverage Interface**: Defines the `PrepareRecipe` method, which is the template method outlining the steps for preparing a beverage.

2. **TemplateBeverage Struct**: Provides default implementations for some common steps like boiling water and pouring into a cup. This struct is used by both concrete beverage types.

3. **Tea and Coffee Structs**: These are concrete implementations of the `Beverage` interface. They override specific steps, such as brewing and adding condiments, to provide custom behavior for making tea and coffee.

### Workflow

- The `PrepareRecipe` method is the template method that defines the sequence of steps:
  1. Boil water.
  2. Brew the beverage.
  3. Pour into a cup.
  4. Add condiments.

- `BoilWater` and `PourInCup` are shared methods provided by `TemplateBeverage`.

- `Brew` and `AddCondiments` are overridden by `Tea` and `Coffee` to provide specific behavior.

### Purpose of PrepareRecipe() in the Interface

- Defines the Template: The `PrepareRecipe()` method serves as the template method that outlines the algorithm's structure. It dictates the sequence of steps that should be followed to prepare any beverage.

- Ensures Consistency: By having `PrepareRecipe()` in the interface, you ensure that any concrete implementation of `Beverage` respects the overall process flow. This guarantees that every subclass will implement the method, preserving the algorithm's skeleton.

### Running the Example

The `main` function demonstrates how we can use the `PrepareRecipe` method to make tea and coffee. Each beverage follows the same sequence of steps, but with specific implementations for brewing and adding condiments.

By using the Template Method Pattern, we ensure that the overall process of making a beverage remains consistent while allowing flexibility in the details of each specific beverage type.

## Conclusion

The Template Method Pattern is a powerful tool for managing common behavior and allowing specific customization in a structured way. It is particularly useful when the algorithm's sequence should be fixed, but certain steps require flexibility.
