package main

import (
	"fmt"
	"time"
)

// customer struct
type customer struct {
    name string
    phone string
}

// conposition
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time
    customer
}


func main() {
    // newCustomer := customer{
    //     name: "Uzzal roy",
    //     phone: "+8801314578964",
    // }

	newOrder := order{
        id: "1",
        amount: 25.4,
        status: "received",
        createdAt: time.Now(),
        customer: customer{
            name: "Uzzal roy",
            phone: "+8801314578964",
        },
    }

    newOrder.customer.name = "Tapas Roy"

    fmt.Println(newOrder)

}