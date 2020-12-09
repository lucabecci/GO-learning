package main

import "fmt"

func intSequencial() func() int{
	i := 0
	return func() int {
		fmt.Println("value now:", i)
		i++ 
		return i
	}
}

func main() {
	nextInt := intSequencial()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println("Final value:", nextInt())
}