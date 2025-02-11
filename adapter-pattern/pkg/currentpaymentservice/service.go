package currentpaymentservice

import "fmt"

// PaymentGateway is the interface that our application uses.
type PaymentGateway interface {
	Pay(amount float64)
}

// SimplePaymentService is a payment service that directly implements PaymentGateway.
type SimplePaymentService struct{}

func (sps *SimplePaymentService) Pay(amount float64) {
	fmt.Printf(`Processed payment of $%.2f through SimplePaymentService`, amount)
	fmt.Println()
}
