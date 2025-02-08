# Payment Processing System

This is a simple payment processing system implemented in Go, demonstrating the use of the Strategy Pattern. The system supports multiple payment methods: Credit Card, PayPal, and Bitcoin. The Strategy Pattern allows easy switching between these payment methods at runtime.

## Overview

In this project, the Strategy Pattern is used to encapsulate different payment algorithms into separate classes, each implementing a common interface. This design allows for flexible and extensible payment processing with minimal changes to the client code.

## Structure

- `PaymentModeHandler` Interface: Defines the method `ProcessPayment` that all payment strategies must implement.

- Concrete Strategies:
  - `CreditCardMode`: Implements payment processing for credit cards.
  - `CashMode`: Implements payment processing for cash.

- `PaymentContext`: Maintains a reference to a `PaymentStrategy` object and delegates the payment processing to it.

## How It Works

1. **Strategy Interface**: The `PaymentStrategy` interface defines a common method `ProcessPayment` for processing a payment.

2. **Concrete Strategies**: Each payment method (Credit Card, PayPal, Bitcoin) is implemented as a separate class that adheres to the `PaymentStrategy` interface.

3. **Context**: The `PaymentContext` class holds a reference to a `PaymentStrategy` and uses it to process payments. It allows changing the payment strategy at runtime.

## Usage

1. **Clone the repository**:

   ```bash
   git clone https://github.com/antblood/design-patterns.git
   cd design-patterns
   ```

2. **Run the program**:
   ```bash
   go run main.go
   ```
