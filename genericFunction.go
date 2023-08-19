package main

import "fmt"

func Add(a int, b int) int {
	return a + b;
}

func main() {
	result := Add(4, 5)
	fmt.Printf("Result: %+v\n", result)
}