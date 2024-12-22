package main

import (
	"fmt"
	"time"
)

// order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here
	customerOrder := order{
		id: id,
		amount: amount,
		status: status,
	}

	return &customerOrder;
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}

func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	/*
	customerOrder := order{
		id: "123456",
		amount: 25.6,
		status: "received",
	}

	customerOrder.createdAt = time.Now()
	customerOrder.status = "paid"

	fmt.Println(customerOrder.status)
	fmt.Println("Order struct:",customerOrder)

	newOrder := order{
		id: "123457",
		amount: 12.36,
		status: "delivered",
		createdAt: time.Now(),
	}
	newOrder.changeStatus("confirmed")

	fmt.Println("New order amount:", newOrder.getAmount())
	fmt.Println("new order:", newOrder)
	*/

	myOrder := newOrder("123458", 45.69, "received")
	fmt.Println(myOrder.amount)

	language := struct {
		name string
		isGood bool
	} {"golan", true}

	fmt.Println(language)
}