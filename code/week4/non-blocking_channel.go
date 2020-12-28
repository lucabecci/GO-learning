package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		messages <- "Hello"
	}()

	time.Sleep(2 * time.Second) //for test the first select

	select {
	case msg := <-messages:
		fmt.Println("received msg:", msg)
	default:
		fmt.Println("no msg received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("Received message:", msg)
	case sig := <-signals:
		fmt.Println("Received signal", sig)
	default:
		fmt.Println("no activity in the channels")

	}
}
