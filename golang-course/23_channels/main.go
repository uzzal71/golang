package main

import "fmt"

func processNum(numChan chan int) {
	fmt.Println("processing number", <-numChan)
}

func main() {
	/*
		messageChan := make(chan string)

		messageChan <- "ping"

		msg := <-messageChan
		fmt.Println(msg)
	*/

	numChan := make(chan int)
	go processNum(numChan)
	numChan <- 5
}
