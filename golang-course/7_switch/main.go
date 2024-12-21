package main

import (
	"fmt"
	"time"
)

func main() {
	// simple switch
	i := 3

	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	// multiple switch condition
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's weekend")
	default:
		fmt.Println("it's working day")
	}

	// type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		case bool:
			fmt.Println("it's a bool")
		default:
			fmt.Println("Other", t)
		}
	}

	whoAmI("golang")
	whoAmI(20)
	whoAmI(true)
	whoAmI(10.2)
}