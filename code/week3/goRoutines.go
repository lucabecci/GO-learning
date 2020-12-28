package main

import (
	"fmt"
	"time"
)

func normalFunction(depends string) {
	for i := 0; i < 3; i++ {
		fmt.Println(depends, ":", i)
	}
}

func goRoutineFunction(depends string) {
	for i := 0; i < 3; i++ {
		fmt.Println(depends, ":", i)
	}
}

func main() {
	normalFunction("synchronously")

	go goRoutineFunction("asynchronously")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
