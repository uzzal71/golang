package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Type the year you ware born: ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("You typed: %q", input)
}