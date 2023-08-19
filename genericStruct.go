package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	ID	int
	Name string
	Data T
}

func main() {
	user := User[string] {
		ID: 0,
		Name: "",
		Data: "abc",
	}

	fmt.Printf("user: %+v\n", user)
}