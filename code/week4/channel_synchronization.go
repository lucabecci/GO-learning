package main

import (
	"fmt"
	"time"
)

func systemWorker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done...")

	done <- true
}

func main() {
	done := make(chan bool, 1)

	go systemWorker(done)

	<-done
}
