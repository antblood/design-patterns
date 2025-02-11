# Adapter Pattern

The Adapter Design Pattern is a structural design pattern that allows objects with incompatible interfaces to work together. It acts as a bridge between two incompatible interfaces, enabling integration without modifying existing code.

## Why the Adapter Pattern is Good

- **Decoupling**: The Adapter Pattern decouples the client code from the specific implementations of the new services, promoting flexibility and allowing the system to evolve independently.
- **Reusability**: It allows existing code to interface with new classes without modification, increasing code reuse.
- **Single Responsibility Principle**: The Adapter Pattern adheres to the Single Responsibility Principle by separating the interface conversion logic from the business logic.
- **Ease of Integration**: It simplifies the process of integrating new systems or components into existing applications, reducing the need for extensive refactoring.

## When to Use the Adapter Pattern

- **Integrating with Legacy Code**: When you have existing code that expects a specific interface, and you need to integrate a new module or service that does not match this interface.
- **Third-party Libraries**: When you want to use a third-party library or service with an incompatible interface with your application.
- **Interface Compatibility**: When ensuring compatibility between differing interfaces without modifying existing code is necessary.

## Payment Processing Application

This application demonstrates the use of the Adapter Design Pattern in Go. The purpose of the application is to showcase how we can integrate a new payment service with an incompatible interface into an existing system that relies on a specific `PaymentGateway` interface.

### Overview

The application consists of the following components:

1. **PaymentGateway Interface**: This is the interface that our existing application uses to process payments. Any new payment service needs to implement this interface to be compatible with the system.

2. **SimplePaymentService**: A direct implementation of the `PaymentGateway` interface, representing a payment service that is already compatible with the system.

3. **NewPaymentService**: A new, third-party payment service with an incompatible interface that we want to integrate into our system.

4. **NewPaymentServiceAdapter**: An adapter that enables the `NewPaymentService` to be used as a `PaymentGateway` by translating the calls from the `PaymentGateway` interface to the methods provided by `NewPaymentService`.

5. **Main Function**: Demonstrates the usage of both the direct implementation and the adapter to process payments.

### Usage

```bash
go run main.go
```
