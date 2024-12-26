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

// gorounting synchronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func emailSender(emailChan chan string, done chan bool) {
	defer func() { done <- true }()

	for email := range emailChan {
		fmt.Println("sending email to", email)
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

	/*
		done := make(chan bool)
		go task(done)
		<-done
	*/

	emailChan := make(chan string, 10)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 10; i++ {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending...")
	close(emailChan)
	<-done

	/*
		emailChan <- "uzzal@gmail.com"
		emailChan <- "sujon@gmail.com"

		fmt.Println(<-emailChan)
		fmt.Println(<-emailChan)

		<-done
	*/
}
