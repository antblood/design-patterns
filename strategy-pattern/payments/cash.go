package payments

import (
	"errors"
	"fmt"
)

type CashMode struct {
	UserPhoneNumber string
}

func (c *CashMode) ProcessPayment(amount float64) (string, error) {
	if amount < 10 {
		return "", errors.New("amount too low")
	}
	return fmt.Sprintf("Processing cash payment for user %s of amount %f", c.UserPhoneNumber, amount), nil
}
