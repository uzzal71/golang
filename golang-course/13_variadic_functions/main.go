package main

import "fmt"

// need to any type use ...interface{}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total;
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := sum(nums...)
	fmt.Println(result)
}