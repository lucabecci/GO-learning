package main

import "fmt"

func main() {
	str := make([]string, 2)
	fmt.Println("arr1", str)	

	str[0] = "a"
	str[1] = "b"
	fmt.Println("arr1", str)	

	str = append(str, "c", "d")
	fmt.Println("arr1 changed:", str)

	fmt.Println("~~~~||||~~~~")

	//new arr made with the first arr
	str2 := make([]string, len(str))
	copy(str2,str)

	fmt.Println("arr2", str2)

	fmt.Println("~~~~||||~~~~")

	extractArr1 := str[0:2]
	fmt.Println("extract arr1", extractArr1)

	fmt.Println("~~~~||||~~~~")

	inlineArr := []string{"a", "e", "i", "o", "u"}
	fmt.Println("inline arr", inlineArr)

	fmt.Println("~~~~||||~~~~")

	arr2D := make([][]int, 3)
	for i:=0; i < len(arr2D); i++{
		innerLen := i + 1
		arr2D[i] = make([]int, innerLen)
		for b:=0; b < innerLen; b++{
			arr2D[i][b] = i + b
		}
	}

	fmt.Println("2DARR", arr2D)
}