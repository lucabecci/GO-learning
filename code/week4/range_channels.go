package main

import "fmt"

func main() {
	queue := make(chan string, 2)

	queue <- "first channel message"
	queue <- "second channel message"

	close(queue)

	for value := range queue {
		fmt.Println(value)
	}

}
