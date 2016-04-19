package main

import "fmt"

type VisaGateway struct {
	Name string
	Url  string
}

func NewVisaGateway() *VisaGateway {
	return &VisaGateway{
		Name: "Visa",
		Url:  "visa.com...",
	}
}

func (v *VisaGateway) Charge() bool {
	fmt.Println("I am charging Visa -->")
	return true
}

type PaymentGateway interface {
	Charge() bool
}

func ChargeCustomer(g PaymentGateway) bool {
	return g.Charge()
}

func main() {
	gateway := NewVisaGateway()
	ChargeCustomer(gateway)
}
