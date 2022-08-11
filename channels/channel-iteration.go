package main

import (
	"fmt"
	"time"
)

var alphabets = []string{"A", "B", "C", "D", "E"}

func main() {

	aChannel := make(chan string)

	go sendAlphabets(aChannel)

	time.Sleep(1 * time.Second)
	fmt.Println("Starting main !!")

	for a := range aChannel {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Received : ", a)
	}

	fmt.Println("Completed main !!")

}

func sendAlphabets(aChannel chan string) {
	fmt.Println("Sending Alphabets")
	for _, a := range alphabets {
		aChannel <- a
	}
	close(aChannel)
	fmt.Println("Alphabets sent")
}
