package newpaymentservice

// NewPaymentServicePaymentAdapter is the adapter that makes NewPaymentService compatible with PaymentGateway.
type NewPaymentServicePaymentAdapter struct {
	newService *NewPaymentService
}

func (adapter *NewPaymentServicePaymentAdapter) Pay(amount float64) {
	// The adapter translates the Pay method call to the ExecutePayment method of NewPaymentService.
	adapter.newService.ExecutePayment(amount, "Default Source")
}
