package main

import "fmt"

func main(){
	if 2+2 == 5{
		fmt.Println("The first conditional is true")
	} else {
		fmt.Println("The first conditional is false")
	}

	fmt.Println("~~~~||||~~~~")

	const operation int = 22/2

	if operation == 12{
		fmt.Println("The result of the operation is", 12)
	} else if operation == 11 {
		fmt.Println("The result of the operation is", 11)
	} else {
		fmt.Println("No result for this operation")
	}

}