package main

import "fmt"

/*
func printSlice(items []int) {
	for _, item := range items {
		fmt.Println(item)
	}
}


func printStringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}
*/

func printSlice[T int | string | bool](items []T) { // use any or interface{} or int | string
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	/*
		nums := []int{1, 2, 3}
		printSlice(nums)
		names := []string{"golan", "typescript"}
		printStringSlice(names)
	*/

	nums := []int{1, 2, 3}
	printSlice(nums)
	names := []string{"golan", "typescript"}
	printSlice(names)
	bols := []bool{true, false, true}
	printSlice(bols)

}
