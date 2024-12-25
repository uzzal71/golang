package main

import "fmt"

// enumerated types
type OrderStatus string

const (
	Received  OrderStatus = "received" //iota auto increments default: 0
	Confirmed             = "confirmed"
	Prepared              = "prepared"
	Delivered             = "delivered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("changing order status to", status)
}

func main() {
	changeOrderStatus(Prepared)
}
