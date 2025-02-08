package shop_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/antblood/strategy-pattern/cart"
	"github.com/antblood/strategy-pattern/shop"
	"github.com/stretchr/testify/require"
)

func getItems() []cart.Item {
	return []cart.Item{
		{Name: "Item 1", Price: 100, Quantity: 1},
		{Name: "Item 2", Price: 200, Quantity: 1},
	}
}

func TestShopper(t *testing.T) {
	must := require.New(t)

	tests := []struct {
		name            string
		shopper         func() *shop.Shopper
		items           []cart.Item
		expectedMessage string
		expectedError   error
	}{
		{
			name:            "StoreCreditCardShopper",
			shopper:         shop.NewStoreCreditCardShopper,
			items:           getItems(),
			expectedMessage: "Processing credit card payment using 1234567890 expiry date 12/24 and CVV 123 for amount 300.000000",
			expectedError:   nil,
		},
		{
			name:            "OnlineCreditCardShopper",
			shopper:         shop.NewOnlineCreditCardShopper,
			items:           getItems(),
			expectedMessage: "Processing credit card payment using 1234567890 expiry date 12/24 and CVV 123 for amount 300.000000",
			expectedError:   nil,
		},
		{
			name:            "OnlineCreditCardShopperHighValue",
			shopper:         shop.NewOnlineCreditCardShopper,
			items:           []cart.Item{{Name: "Item 1", Price: 1100, Quantity: 1}},
			expectedMessage: "",
			expectedError:   fmt.Errorf("amount too high"),
		},
		{
			name:            "StoreCashShopper",
			shopper:         shop.NewStoreCashShopper,
			items:           getItems(),
			expectedMessage: "Processing cash payment for user 1234567890 of amount 290.000000",
			expectedError:   nil,
		},
		{
			name:            "StoreCashShopperLowValue",
			shopper:         shop.NewStoreCashShopper,
			items:           []cart.Item{},
			expectedMessage: "",
			expectedError:   errors.New("amount too low"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			message, err := test.shopper().Checkout(test.items)
			must.Equal(err, test.expectedError)
			must.Equal(message, test.expectedMessage)
		})
	}
}
