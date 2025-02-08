package payments

type PaymentModeHandler interface {
	ProcessPayment(amount float64) (string, error)
}
