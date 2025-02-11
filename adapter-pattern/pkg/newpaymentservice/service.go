package newpaymentservice

import "fmt"

// NewPaymentService is the incompatible interface that we want to integrate.
type NewPaymentService struct{}

func (nps *NewPaymentService) ExecutePayment(amount float64, source string) {
	fmt.Printf(`Processed payment of $%.2f through NewPaymentService. Source: %s`, amount, source)
	fmt.Println()
}
