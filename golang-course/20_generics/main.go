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

func printSlice[T comparable, N string](items []T, name N) {
	for _, item := range items {
		fmt.Println(item)
	}
}
*/

func printSlice[T comparable](items []T) { // use any or interface{} or int | string or comparable
	for _, item := range items {
		fmt.Println(item)
	}
}

// LIFO
// use any or interface{} or int | string or comparable
type stack[T comparable] struct {
	elements []T
}

func main() {
	/*
		nums := []int{1, 2, 3}
		printSlice(nums)
		names := []string{"golan", "typescript"}
		printStringSlice(names)

		nums := []int{1, 2, 3}
		printSlice(nums)
		names := []string{"golan", "typescript"}
		printSlice(names)
		bols := []bool{true, false, true}
		printSlice(bols)
	*/

	myStack := stack[string]{
		elements: []string{"golang", "Typescript"},
	}
	fmt.Println(myStack)

}
