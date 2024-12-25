package main

import "fmt"

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

func main() {
	nums := []int{1, 2, 3}
	names := []string{"golan", "typescript"}
	printSlice(nums)
	printStringSlice(names)
}
