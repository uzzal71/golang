package main

import (
	"fmt"
	"time"
)

// Sending
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

// Receive
func sum(result chan int, num1 int, num2 int) {
	sumResult := num1 + num2
	result <- sumResult
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func main() {
	/*
		messageChan := make(chan string)

		messageChan <- "ping"

		msg := <-messageChan
		fmt.Println(msg)
	*/

	/*
		numChan := make(chan int)
		go processNum(numChan)
		for {
			numChan <- rand.Intn(10)
		}
	*/

	/*
		result := make(chan int)
		go sum(result, 4, 5)
		res := <-result
		fmt.Println(res)
	*/

	done := make(chan bool)
	go task(done)
	<-done
}
