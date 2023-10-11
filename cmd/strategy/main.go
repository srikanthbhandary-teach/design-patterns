package main

import (
	"design-patterns/behavioral"
	"fmt"
)

func main() {
	// Create payment strategies
	creditCardPayment := &behavioral.CreditCard{}
	payPalPayment := &behavioral.Paypal{}

	// Create a payment processor
	paymentProcessor := &behavioral.PaymentProcessor{}

	// Set the payment strategy to credit card and process payment
	paymentProcessor.SetStrategy(creditCardPayment)
	fmt.Println(paymentProcessor.Pay(100.0))

	// Change the payment strategy to PayPal and process payment
	paymentProcessor.SetStrategy(payPalPayment)
	fmt.Println(paymentProcessor.Pay(50.0))
}
