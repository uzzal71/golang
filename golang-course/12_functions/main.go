package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

// both a & b int type shared
func sub(a, b int) int {
	return a - b
}

// multiple result statements
func getLanguages() (string, string, string) {
	return "golan", "javascript", "typescript"
}

func processIt(fn func(a int) int) {
	fn(1)
}

func paymentProcess() func (a int) int {
	return func(a int) int {
		return a
	}
}

func main() {
	result := add(3, 5)
	fmt.Println("Result:", result)

	lang1, lang2, _ := getLanguages()
	fmt.Println(lang1, lang2)

	fn := func(a int) int {
		return a
	}
	processIt(fn)

	payment := paymentProcess()
	payment(10)
}