package cart_test

import (
	"testing"

	"github.com/antblood/strategy-pattern/cart"
	"github.com/stretchr/testify/require"
)

func getItems() []cart.Item {
	return []cart.Item{
		{Name: "Item 1", Price: 100, Quantity: 1},
		{Name: "Item 2", Price: 200, Quantity: 1},
	}
}

func TestStoreCart(t *testing.T) {
	must := require.New(t)
	sc := cart.StoreCart{
		StoreDiscount: 10,
	}
	total := sc.CalculateTotal(getItems())
	must.Equal(total, 290.0)
}

func TestOnlineCart(t *testing.T) {
	must := require.New(t)
	oc := cart.OnlineCart{}
	total := oc.CalculateTotal(getItems())
	must.Equal(total, 300.0)
}
