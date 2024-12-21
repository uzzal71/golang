package main

import "fmt"

const city string = "Dhaka"
const country string = "Bangladesh"

const (
	port = 5000
	host = "http://localhost"
)

func main() {
	const name string = "golang"
	fmt.Println(name)

	const age int32 = 30;
	fmt.Println(age)

	fmt.Println(city, country)

	fmt.Println(port, host)
}