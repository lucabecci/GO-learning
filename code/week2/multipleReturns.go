package main

import "fmt"

func user(
	name string, 
	lastname string, 
	age int) (string, string, int) {
    return name, lastname, age
}

func main() {

	userName, _, _ := user("Luca", "Becci", 20) 
	_, lastname, age := user("Luca", "Becci", 20) 
	fmt.Println(userName)
	fmt.Println(lastname, age)
}