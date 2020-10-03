package main

import (
	"fmt"
)

func pinger(ch chan string) {

	ch <- "PING"

}

func ponger(ch chan string) {

	ch <- "PONG"

}

func main() {

	fmt.Println("Hello World")

	var ch chan string = make(chan string, 3)

	ch <- "Word1"
	ch <- "Word2"
	ch <- "Word3"

	close(ch)

	// below for loop is good when channel is closed ....
	//but when channel is closed it continues to listen for the incomming message

	//for msg := range ch {

	//	fmt.Println(msg)
	//}
	fmt.Println("-----------------------")

	for {

		if msg1, ok := <-ch; ok {

			fmt.Println(msg1)

		} else {
			break
		}
	}
	fmt.Println("-----------------------")

	// creating a routine and channel to exchange the data....
	var ch2 chan string = make(chan string, 3)
	go pinger(ch2)
	go ponger(ch2)
	close(ch2)
	for msg := range ch2 {

		fmt.Println(msg)
	}

}
