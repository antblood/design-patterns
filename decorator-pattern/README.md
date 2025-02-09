# Decorator Pattern

This project demonstrates the implementation of the Decorator Pattern in Go. The pattern is used here to create a versatile beverage system where you can start with a base beverage and add various addons to it.

## Overview

The system is designed with a `Beverage` interface, which includes methods for getting the description and the cost of a beverage. The base beverages provided are Cold Coffee and Cappuccino. You can enhance these beverages by adding addons like Milk and Espresso Shot.

## Structure

- **Beverage Interface**: This is the core interface which all beverages and decorators implement. It has two methods:
  - `Description() string`: Returns a description of the beverage.
  - `Cost() float64`: Returns the cost of the beverage.

- **Base Beverages**: These are the concrete implementations of the `Beverage` interface.
  - `ColdCoffee`: A simple cold coffee beverage.
  - `Cappuccino`: A classic cappuccino beverage.

- **Addons (Decorators)**: These add extra functionality to the base beverages.
  - `Milk`: Adds milk to the beverage.
  - `EspressoShot`: Adds an extra espresso shot to the beverage.
