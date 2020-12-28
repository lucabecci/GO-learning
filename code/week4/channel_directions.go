package main

import "fmt"

func hello(hellos chan<- string, msg string) {
	hellos <- msg
}

func bye(hellos <-chan string, byes chan<- string) {
	msg := <-hellos

	byes <- msg
}

func main() {
	hellos := make(chan string, 1)
	byes := make(chan string, 1)
	hello(hellos, "Hello my friend. How are you?")
	bye(hellos, byes)

	fmt.Println(<-byes)
}
