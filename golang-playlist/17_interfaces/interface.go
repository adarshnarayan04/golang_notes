package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

// Open close principle
type payment struct {
	
	//gateway razorpay , have to change everytime we want to change
	//gateway stripe

	// so used interface
	gateway paymenter
}


func (p payment) makePayment(amount float32) {
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razorpay", amount)
}

// type stripe struct{}

// func (s stripe) pay(amount float32) {
// 	fmt.Println("making payment using stripe", amount)
// }

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose")
}

type paypal struct{}

//paypal implements paymenter interface ( has pay and refund methods )
func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {

}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	// fakeGw := fakepayment{}
	paypalGw := paypal{}
	newPayment := payment{
		gateway: paypalGw, // now by changing this line we can change payment gateway
	}
	newPayment.makePayment(100)
}
