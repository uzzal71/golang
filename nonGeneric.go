package main

import "fmt"

func MapValues(values []int, mapFunc func(int) int) []int {
	var newValues []int
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

func mainNon() {
	result := MapValues([]int{1,2,3}, func(n int) int {
		return n * 3
	})
	fmt.Printf("result %+v\n", result)
}