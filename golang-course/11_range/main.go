package main

import "fmt"

func main() {
	products := []string{"Apple 15s", "Samsung 23ultra"}
	for i :=0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	for _, product := range products {
		fmt.Println(product)
	}

	prices := []int{21, 47, 63}
	total := 0
	for _, price := range prices {
		total = total + price
	}
	fmt.Println("Total price:", total)

	for i, price := range prices {
		fmt.Println(i, price)
	}

	st := map[string]string{"name": "uzzal", "department": "CSE", "batch": "45"}
	for key, value := range st {
		fmt.Println(key, value)
	}

	for key := range st {
		fmt.Println(key)
	}

	for i, c := range "golang" {
		fmt.Println(i, c)
	}
}