package main

import "fmt"

func plus(a int, b int) int{
	return a + b
}

func user(name, lastname string) string {
	return name + " " + lastname
}

func main() {
	result1 := plus(22, 33) 
	fmt.Println(result1)

	resultUser := user("Luca", "Becci")
	fmt.Println(resultUser)
}