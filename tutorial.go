package main

import (
	//"bufio"
	"fmt"
	//"os"
	//"strconv"
)

func main_() {
	/*
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you ware born: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	fmt.Printf("You will be %d year old at the end of 2024", 2024 - input)
	*/
	var num1 float64 = 8
	var num2 float64 = 4
	answer := num1 / num2

	fmt.Printf("Result is %g", answer)
}