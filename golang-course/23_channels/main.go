package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
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
	// numChan <- 5
	for {
		numChan <- rand.Intn(10)
	}
}
