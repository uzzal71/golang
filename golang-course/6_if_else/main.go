package main

import "fmt"

func main() {
	age := 16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else {
		fmt.Println("Person is not a adult")
	}

	price := 20

	if price >= 18 {
		fmt.Println("Price is a high")
	} else if price >= 12 {
		fmt.Println("Price is low")
	} else {
		fmt.Println("Price is a big low")
	}

	role := "admin"
	hasPermissions := true

	if role == "admin" && hasPermissions {
		fmt.Println("yes")
	}

	if age := 15; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("Person is teenager")
	}
}