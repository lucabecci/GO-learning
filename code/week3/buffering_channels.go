package main

import "fmt"

func main() {
	messages := make(chan string, 2)

	messages <- "Buffered channel..."
	messages <- "Recived Channel..."

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
