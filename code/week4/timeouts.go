package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		channel1 <- "First Result"
	}()

	select {
	case res := <-channel1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("First Timeout")
	}

	channel2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "Second Result"
	}()
	select {
	case res := <-channel2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("Second Timeour")
	}

}
