package main

import (
	"fmt"
	"time"
)

// Define the referenced struct
type Customer struct {
	Name  string
	Phone string
}

// Struct with a pointer to another struct
type OrderWithPointer struct {
	ID        string
	Amount    float32
	Status    string
	CreatedAt time.Time
	Customer  *Customer // pointer to Customer struct
}

// Struct with a named field (not pointer)
type OrderNamed struct {
	ID        string
	Amount    float32
	Status    string
	CreatedAt time.Time
	Customer  Customer // value, not pointer
}

// Struct with embedded field (anonymous)
type OrderEmbedded struct {
	ID        string
	Amount    float32
	Status    string
	CreatedAt time.Time
	Customer  // embedded, fields promoted
}

func main() {
	// Using pointer field
	cust := &Customer{Name: "Alice", Phone: "1112223333"}
	orderPtr := OrderWithPointer{
		ID:        "1",
		Amount:    100,
		Status:    "pending",
		CreatedAt: time.Now(),
		Customer:  cust,
	}
	fmt.Println("Pointer field:", orderPtr.Customer.Name) // Access via pointer

	// Change the original cust struct
	cust.Name = "Alice Updated"
	cust.Phone = "9998887777"

	// orderPtr.Customer reflects the changes, since it's the same pointer
	fmt.Println("After changing cust:")
	fmt.Println("cust:", cust)
	fmt.Println("orderPtr.Customer:", orderPtr.Customer)

	// Using named field (value)
	orderNamed := OrderNamed{
		ID:        "2",
		Amount:    200,
		Status:    "confirmed",
		CreatedAt: time.Now(),
		Customer:  Customer{Name: "Bob", Phone: "4445556666"},
	}
	fmt.Println("Named field:", orderNamed.Customer.Name) // Access via named field

	// Using embedded field
	orderEmbed := OrderEmbedded{
		ID:        "3",
		Amount:    300,
		Status:    "shipped",
		CreatedAt: time.Now(),
		Customer:  Customer{Name: "Carol", Phone: "7778889999"},
	}
	fmt.Println("Embedded field:", orderEmbed.Name, orderEmbed.Name) // Direct access due to promotion
	// also can access like orderEmbed.Customer.Name

	// Difference:
	// - Pointer field: Can be nil, allows sharing/mutation, must check for nil before access.
	// - Named field: No promotion, access via orderNamed.Customer.Name.
	// - Embedded field: Fields are promoted, access via orderEmbed.Name.
}
