package main

import "fmt"

func fact(n int) int {
	if n == 0{
		fmt.Println("value final:", n)
		return 1
	}
	fmt.Println("value now:", n)
	return n + fact(n - 1)
}

func main() {
	fact(7)
}