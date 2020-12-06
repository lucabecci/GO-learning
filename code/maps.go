package main

import "fmt"

func main() {
	myMap := make(map[string]int)

	myMap["ident1"] = 3
	myMap["ident2"] = 6  

	fmt.Println(myMap)

	fmt.Println("~~~~||||~~~~")

	value1 := myMap["ident1"]

	fmt.Println("value one:", value1)

	fmt.Println("~~~~||||~~~~")

	delete(myMap, "ident2")
	fmt.Println("removed", myMap)

	fmt.Println("~~~~||||~~~~")

	value, exists := myMap["ident1"]
    fmt.Println("exists:", exists, "|| value:", value)

}