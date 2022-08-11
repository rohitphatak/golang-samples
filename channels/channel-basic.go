package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start Main")

	// Make a channel
	ch := make(chan string, 1)

	go sendHello(ch)

	time.Sleep(2 * time.Second)
	receiver := <-ch
	fmt.Println("Received Value : ", receiver)
}

func sendHello(ch chan string) {
	fmt.Println("From sendHello")
	ch <- "Sending Hi!!!"
	fmt.Println("Hello complete")
}
