package main

// PaymentGateway is the interface that our application uses.
type PaymentGateway interface {
	Pay(amount float64)
}
