package shop

import (
	"github.com/antblood/strategy-pattern/cart"
	"github.com/antblood/strategy-pattern/payments"
)

type Shopper struct {
	Cart    cart.CartTypeHandler
	Payment payments.PaymentModeHandler
}

func (s *Shopper) Checkout(items []cart.Item) (string, error) {
	total := s.Cart.CalculateTotal(items)
	message, err := s.Payment.ProcessPayment(total)
	if err != nil {
		return "", err
	}
	return message, nil
}

func NewShopper(cart cart.CartTypeHandler, payment payments.PaymentModeHandler) *Shopper {
	return &Shopper{
		Cart:    cart,
		Payment: payment,
	}
}

func NewStoreCreditCardShopper() *Shopper {
	return NewShopper(&cart.StoreCart{}, &payments.CreditCardMode{
		CardNumber: "1234567890",
		ExpiryDate: "12/24",
		CVV:        "123",
	})
}

func NewOnlineCreditCardShopper() *Shopper {
	return NewShopper(&cart.OnlineCart{}, &payments.CreditCardMode{
		CardNumber: "1234567890",
		ExpiryDate: "12/24",
		CVV:        "123",
	})
}

func NewStoreCashShopper() *Shopper {
	return NewShopper(&cart.StoreCart{
		StoreDiscount: 10,
	}, &payments.CashMode{
		UserPhoneNumber: "1234567890",
	})
}
