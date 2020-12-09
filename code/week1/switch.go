package main

import (
	"fmt"
	"time"
)

func main() {
	var i int = 2

	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	fmt.Println("~~~~||||~~~~")

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend:", time.Now().Weekday())
	default:
		fmt.Println("It's the weekday:", time.Now().Weekday())
		
	}

	fmt.Println("~~~~||||~~~~")

	t := time.Now()

	switch  {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default: 
		fmt.Println("It's after noon")
	}
}