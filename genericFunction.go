package main

import (
	"fmt"
)

type Num interface {
	int | int8 | int32 | int64 | float32 | float64
}

type UserId int

func Add[T ~int | float64](a T, b T) T {
	return a + b;
}

func mainG() {
	a := UserId(1)
	b := UserId(3)
	result := Add(a, b)
	fmt.Printf("Result: %+v\n", result)
}