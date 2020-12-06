package main

import "fmt"

func main() {
	var firstArr [3]int
	firstArr[0] = 1
	firstArr[2] = 22
	fmt.Println("first:", firstArr)

	fmt.Println("arr longitude:", len(firstArr))

	fmt.Println("~~~~||||~~~~")

	secondArr := [5]int{22, 3, 22, 3, 22}
	fmt.Println("second:", secondArr)
	fmt.Println("arr longitude:", len(secondArr))

	fmt.Println("~~~~||||~~~~")

	var threeArr[3][4] int
	for i := 0; i < len(threeArr); i++{
		for j := 0; j < len(threeArr[0]); j++{
			threeArr[i][j] = i + j
		}
	} 
	fmt.Println("2D array:", threeArr)

	
}