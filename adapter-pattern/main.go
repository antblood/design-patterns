package main

import (
	"github.com/antblood/design-patterns/adapter-pattern/pkg/currentpaymentservice"
	"github.com/antblood/design-patterns/adapter-pattern/pkg/newpaymentservice"
)

func main() {
	oldService := currentpaymentservice.SimplePaymentService{}
	oldService.Pay(100)

	newServiceAdapter := newpaymentservice.NewPaymentServicePaymentAdapter{}
	newServiceAdapter.Pay(100)
}
