package main

import "fmt"

func main() {
	// while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// infinite loop
	// for {
	// 	fmt.Println("1")
	// }

	// Classic for loop
	for i := 0; i <= 3; i++ {
		
		if i == 2 {
			continue
		}

		fmt.Println(i)
	}

	// range
	for j := range 3 {
		fmt.Println(j)
	}
}