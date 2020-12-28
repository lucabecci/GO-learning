package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)

		channel1 <- "First Message..."
	}()

	go func() {
		time.Sleep(3 * time.Second)

		channel2 <- "Second Message..."

	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("first received:", msg1)
		case msg2 := <-channel2:
			fmt.Println("second received:", msg2)
		}
	}

}
