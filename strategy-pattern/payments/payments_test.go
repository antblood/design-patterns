package payments_test

import (
	"fmt"
	"testing"

	"github.com/antblood/strategy-pattern/payments"
	"github.com/stretchr/testify/require"
)

func TestPaymentMode_CreditCard(t *testing.T) {
	must := require.New(t)
	pm := payments.CreditCardMode{
		CardNumber: "1234567890",
		ExpiryDate: "12/25",
		CVV:        "123",
	}

	tests := []struct {
		name            string
		amount          float64
		expectedMessage string
		expectedError   error
	}{
		{name: "100", amount: 100, expectedMessage: "Processing credit card payment using 1234567890 expiry date 12/25 and CVV 123 for amount 100.000000", expectedError: nil},
		{name: "10", amount: 10, expectedMessage: "Processing credit card payment using 1234567890 expiry date 12/25 and CVV 123 for amount 10.000000", expectedError: nil},
		{name: "1001", amount: 1001, expectedMessage: "", expectedError: fmt.Errorf("amount too high")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := pm.ProcessPayment(test.amount)
			must.Equal(err, test.expectedError)
			must.Equal(message, test.expectedMessage)
		})
	}
}

func TestPaymentMode_Cash(t *testing.T) {
	must := require.New(t)
	pm := payments.CashMode{
		UserPhoneNumber: "1234567890",
	}

	tests := []struct {
		name            string
		amount          float64
		expectedMessage string
		expectedError   error
	}{
		{name: "100", amount: 100, expectedMessage: "Processing cash payment for user 1234567890 of amount 100.000000", expectedError: nil},
		{name: "10", amount: 10, expectedMessage: "Processing cash payment for user 1234567890 of amount 10.000000", expectedError: nil},
		{name: "9", amount: 9, expectedMessage: "", expectedError: fmt.Errorf("amount too low")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := pm.ProcessPayment(test.amount)
			must.Equal(err, test.expectedError)
			must.Equal(message, test.expectedMessage)
		})
	}
}
