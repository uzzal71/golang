package main

import "fmt"

type payment struct {
	gateway stripe
}

// Open close principle
func (p payment) makePayment(amount float32) {
    // razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct {}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

type stripe struct {}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

func main() {
	newPayment := payment{}
	newPayment.makePayment(100)
}