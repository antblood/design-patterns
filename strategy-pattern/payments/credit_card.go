package payments

import (
	"errors"
	"fmt"
)

type CreditCardMode struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

func (c *CreditCardMode) ProcessPayment(amount float64) (string, error) {
	if amount > 1000 {
		return "", errors.New("amount too high")
	}
	return fmt.Sprintf("Processing credit card payment using %s expiry date %s and CVV %s for amount %f", c.CardNumber, c.ExpiryDate, c.CVV, amount), nil
}
