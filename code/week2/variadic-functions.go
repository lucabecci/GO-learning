package main

import "fmt"

func sum(nums ...int)  {
	fmt.Print(nums)
	total := 0

	for i, num := range nums {
		total += num
		fmt.Println("index:", i, " value:", num)
	}
	fmt.Println("total of the array:",total)
}

func main() {
	numbers := []int{22, 33, 44, 55, 66}
	sum(numbers...)
}