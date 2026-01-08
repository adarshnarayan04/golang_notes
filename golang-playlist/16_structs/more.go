package main
import (
	"fmt"
	"time"
)
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
}

func newOrder(id string, amount float32, status string) order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return myOrder
}

func newOrderpointer(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func main(){

	myOrder := newOrder("1", 30.50, "received")
	fmt.Println(myOrder)
	fmt.Println(myOrder.amount)

	myOrderPtr := newOrderpointer("2", 45.00, "shipped")
	fmt.Println(myOrderPtr) // it print with & because its a pointer
	fmt.Println(*myOrderPtr)   // value, no &
	fmt.Printf("%p\n", myOrderPtr) // address
	fmt.Println(myOrderPtr.amount)

}
