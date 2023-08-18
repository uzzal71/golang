package main

import "fmt"

func mainArray() {
	arr1 := [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}
	cars := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	prices := [...]int{10,15,20}


	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(cars)
	// change element of an array
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// Array initialzation
	arr11 := [5]int{} // not initialized
	arr12 := [5]int{1,2} // partially initialized
	arr13 := [5]int{1,2,3,4,5} //full initialized

	fmt.Println(arr11)
	fmt.Println(arr12)
	fmt.Println(arr13)

	// Initialize Only Specific Elements
	arr21 := [5]int{1:10, 2:40}

	fmt.Println(arr21)

	// Find the Length of an Array
	fmt.Println(len(cars))
	fmt.Println(len(arr1))
}