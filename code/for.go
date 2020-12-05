package main

import "fmt"

func main(){
	var i int = 0;

	for i <= 2 { 
		fmt.Println("first loop:", i)
		i++
	}

	fmt.Println("~~~~||||~~~~")

	for x :=0; x<=2; x++{
		fmt.Println("second loop", x)
	}

	fmt.Println("~~~~||||~~~~")

	for {
		fmt.Println("This is a normal for but this have a break in the first loop")
		break
	}

	fmt.Println("~~~~||||~~~~")

	for b := 0; b<=6; b++{
		if(b%2 == 0){
			fmt.Println("Continue because in this loop b is even", b)
			continue
		}
		fmt.Println("odd:", b)

	} 

}