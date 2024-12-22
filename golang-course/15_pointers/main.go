package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("In change number:", *num)
}

func main() {
	num := 1
	changeNum(&num)

	fmt.Println("Memory address:", &num)
	fmt.Println("After change number in main:", num)
}