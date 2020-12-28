package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "Buffered channel..."
	messages <- "Recived Channel..."
	messages <- "Three message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
