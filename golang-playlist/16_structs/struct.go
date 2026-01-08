package main

import (
	"fmt"
	"time"
)
// if you don't set any field, default value is zero value
// int => 0, float => 0, string "", bool => false


// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer // embedded struct (https://gobyexample.com/struct-embedding)
	//customer customer// named field
}

// Embedded (anonymous) field:

// type order struct { customer }
// Fields/methods of customer are promoted.
// Access: newOrder.name and newOrder.phone work, as well as newOrder.customer.name.

// Named field:

// type order struct { customer customer }
// No promotion.
// Access: only newOrder.customer.name (newOrder.name does not compile).


// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here...
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// // receiver type
// this are methods of order struct ( to acess it order.methodName())
// can write any name ( but first letter should be same as struct name first letter by convention)
//so (o *order) 

//as we want to change the original struct, we use pointer receiver if not then (o order)
// func (o *order) changeStatus(status string) {
// 	o.status = status // dereferencing happens automatically ( so can use o.status instead of (*o).status)
// }

// no need to change value in this method of struct , so not use of pointer reciever here
// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	newOrder.customer.name = "robin"
	fmt.Println(newOrder)

	//struct without a name

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)


	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	//using the methods of struct

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)
}
