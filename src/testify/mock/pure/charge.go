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

func (v *VisaGateway) Charge() {
	fmt.Println("I am charging Visa -->")
}

type PaymentGateway interface {
	Charge()
}

func ChargeCustomer(g PaymentGateway) {
	g.Charge()
}

func main() {
	gateway := NewVisaGateway()
	ChargeCustomer(gateway)
}
